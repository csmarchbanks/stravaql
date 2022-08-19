package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"golang.org/x/oauth2"
)

type tokenSource struct {
	src oauth2.TokenSource
}

func newTokenSource(conf *oauth2.Config, token *oauth2.Token) oauth2.TokenSource {
	src := conf.TokenSource(context.Background(), token)
	return oauth2.ReuseTokenSource(token, &tokenSource{src: src})
}

func (s *tokenSource) Token() (*oauth2.Token, error) {
	token, err := s.src.Token()
	if err != nil {
		return nil, err
	}
	credDir := ".stravaql"
	credFile := path.Join(credDir, "token.json")

	// Cache the token for future use.
	encoded, err := json.Marshal(token)
	if err != nil {
		return nil, fmt.Errorf("marshaling token: %w", err)
	}
	err = os.MkdirAll(credDir, 0700)
	if err != nil {
		return nil, fmt.Errorf("making directory: %w", err)
	}
	err = os.WriteFile(credFile, encoded, 0600)
	if err != nil {
		return nil, fmt.Errorf("caching token: %w", err)
	}

	return token, nil
}

func oauthToken() oauth2.TokenSource {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("STRAVAQL_CLIENT_ID"),
		ClientSecret: os.Getenv("STRAVAQL_CLIENT_SECRET"),
		Scopes:       []string{"activity:read_all"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.strava.com/oauth/authorize",
			TokenURL: "https://www.strava.com/oauth/token",
		},
		RedirectURL: "http://localhost:8952",
	}

	credDir := ".stravaql"
	credFile := path.Join(credDir, "token.json")
	bytes, err := ioutil.ReadFile(credFile)
	if err == nil {
		token := &oauth2.Token{}
		err = json.Unmarshal(bytes, token)
		if err == nil {
			return newTokenSource(conf, token)
		}

	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	state := fmt.Sprintf("%d", time.Now().UnixNano())
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Println(url)

	server := http.Server{
		Addr: ":8952",
	}

	ch := make(chan *oauth2.Token)
	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "invalid state", http.StatusBadRequest)
			return
		}
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "invalid code", http.StatusBadRequest)
			return
		}

		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not exchange code: %+v", err), http.StatusInternalServerError)
			return
		}

		encoded, err := json.Marshal(token)
		if err == nil {
			err = os.MkdirAll(credDir, 0700)
			if err != nil {
				fmt.Println("error making directory", err)
			}
			err = os.WriteFile(credFile, encoded, 0600)
			if err != nil {
				fmt.Println("error writing file", err)
			}
		}

		ch <- token
		close(ch)
	})

	defer server.Shutdown(context.Background())
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error with oauth server: %+v\n", err)
		}
	}()
	return newTokenSource(conf, <-ch)
}
