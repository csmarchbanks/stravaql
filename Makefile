UID := $(shell id -u)
GID := $(shell id -g)

strava-old:
	docker run --rm -v "${PWD}:/local" --user $(UID):$(GID) \
		swaggerapi/swagger-codegen-cli generate \
		-i https://developers.strava.com/swagger/swagger.json \
		-l go \
		--additional-properties packageName=strava \
		-o /local/strava

strava:
	docker run --rm -it --user $(UID):$(GID) \
		-v "${PWD}:/local" \
		-w "/local" \
		quay.io/goswagger/swagger generate client \
		--skip-validation \
		-c strava \
		-m strava/model \
		-f https://developers.strava.com/swagger/swagger.json
