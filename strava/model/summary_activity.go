// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SummaryActivity summary activity
//
// swagger:model summaryActivity
type SummaryActivity struct {
	MetaActivity

	// The number of achievements gained during this activity
	AchievementCount int64 `json:"achievement_count,omitempty"`

	// athlete
	Athlete *MetaAthlete `json:"athlete,omitempty"`

	// The number of athletes for taking part in a group activity
	// Minimum: 1
	AthleteCount int64 `json:"athlete_count,omitempty"`

	// The activity's average speed, in meters per second
	AverageSpeed float32 `json:"average_speed,omitempty"`

	// Average power output in watts during this activity. Rides only
	AverageWatts float32 `json:"average_watts,omitempty"`

	// The number of comments for this activity
	CommentCount int64 `json:"comment_count,omitempty"`

	// Whether this activity is a commute
	Commute bool `json:"commute,omitempty"`

	// Whether the watts are from a power meter, false if estimated
	DeviceWatts bool `json:"device_watts,omitempty"`

	// The activity's distance, in meters
	Distance float32 `json:"distance,omitempty"`

	// The activity's elapsed time, in seconds
	ElapsedTime int64 `json:"elapsed_time,omitempty"`

	// The activity's highest elevation, in meters
	ElevHigh float32 `json:"elev_high,omitempty"`

	// The activity's lowest elevation, in meters
	ElevLow float32 `json:"elev_low,omitempty"`

	// end latlng
	EndLatlng LatLng `json:"end_latlng,omitempty"`

	// The identifier provided at upload time
	ExternalID string `json:"external_id,omitempty"`

	// Whether this activity is flagged
	Flagged bool `json:"flagged,omitempty"`

	// The id of the gear for the activity
	GearID string `json:"gear_id,omitempty"`

	// Whether the logged-in athlete has kudoed this activity
	HasKudoed bool `json:"has_kudoed,omitempty"`

	// Whether the activity is muted
	HideFromHome bool `json:"hide_from_home,omitempty"`

	// The total work done in kilojoules during this activity. Rides only
	Kilojoules float32 `json:"kilojoules,omitempty"`

	// The number of kudos given for this activity
	KudosCount int64 `json:"kudos_count,omitempty"`

	// Whether this activity was created manually
	Manual bool `json:"manual,omitempty"`

	// map
	Map *PolylineMap `json:"map,omitempty"`

	// The activity's max speed, in meters per second
	MaxSpeed float32 `json:"max_speed,omitempty"`

	// Rides with power meter data only
	MaxWatts int64 `json:"max_watts,omitempty"`

	// The activity's moving time, in seconds
	MovingTime int64 `json:"moving_time,omitempty"`

	// The name of the activity
	Name string `json:"name,omitempty"`

	// The number of Instagram photos for this activity
	PhotoCount int64 `json:"photo_count,omitempty"`

	// Whether this activity is private
	Private bool `json:"private,omitempty"`

	// sport type
	SportType SportType `json:"sport_type,omitempty"`

	// The time at which the activity was started.
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// The time at which the activity was started in the local timezone.
	// Format: date-time
	StartDateLocal strfmt.DateTime `json:"start_date_local,omitempty"`

	// start latlng
	StartLatlng LatLng `json:"start_latlng,omitempty"`

	// The timezone of the activity
	Timezone string `json:"timezone,omitempty"`

	// The activity's total elevation gain.
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

	// The number of Instagram and Strava photos for this activity
	TotalPhotoCount int64 `json:"total_photo_count,omitempty"`

	// Whether this activity was recorded on a training machine
	Trainer bool `json:"trainer,omitempty"`

	// type
	Type ActivityType `json:"type,omitempty"`

	// The identifier of the upload that resulted in this activity
	UploadID int64 `json:"upload_id,omitempty"`

	// The unique identifier of the upload in string format
	UploadIDStr string `json:"upload_id_str,omitempty"`

	// Similar to Normalized Power. Rides with power meter data only
	WeightedAverageWatts int64 `json:"weighted_average_watts,omitempty"`

	// The activity's workout type
	WorkoutType int64 `json:"workout_type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SummaryActivity) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MetaActivity
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MetaActivity = aO0

	// AO1
	var dataAO1 struct {
		AchievementCount int64 `json:"achievement_count,omitempty"`

		Athlete *MetaAthlete `json:"athlete,omitempty"`

		AthleteCount int64 `json:"athlete_count,omitempty"`

		AverageSpeed float32 `json:"average_speed,omitempty"`

		AverageWatts float32 `json:"average_watts,omitempty"`

		CommentCount int64 `json:"comment_count,omitempty"`

		Commute bool `json:"commute,omitempty"`

		DeviceWatts bool `json:"device_watts,omitempty"`

		Distance float32 `json:"distance,omitempty"`

		ElapsedTime int64 `json:"elapsed_time,omitempty"`

		ElevHigh float32 `json:"elev_high,omitempty"`

		ElevLow float32 `json:"elev_low,omitempty"`

		EndLatlng LatLng `json:"end_latlng,omitempty"`

		ExternalID string `json:"external_id,omitempty"`

		Flagged bool `json:"flagged,omitempty"`

		GearID string `json:"gear_id,omitempty"`

		HasKudoed bool `json:"has_kudoed,omitempty"`

		HideFromHome bool `json:"hide_from_home,omitempty"`

		Kilojoules float32 `json:"kilojoules,omitempty"`

		KudosCount int64 `json:"kudos_count,omitempty"`

		Manual bool `json:"manual,omitempty"`

		Map *PolylineMap `json:"map,omitempty"`

		MaxSpeed float32 `json:"max_speed,omitempty"`

		MaxWatts int64 `json:"max_watts,omitempty"`

		MovingTime int64 `json:"moving_time,omitempty"`

		Name string `json:"name,omitempty"`

		PhotoCount int64 `json:"photo_count,omitempty"`

		Private bool `json:"private,omitempty"`

		SportType SportType `json:"sport_type,omitempty"`

		StartDate strfmt.DateTime `json:"start_date,omitempty"`

		StartDateLocal strfmt.DateTime `json:"start_date_local,omitempty"`

		StartLatlng LatLng `json:"start_latlng,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

		TotalPhotoCount int64 `json:"total_photo_count,omitempty"`

		Trainer bool `json:"trainer,omitempty"`

		Type ActivityType `json:"type,omitempty"`

		UploadID int64 `json:"upload_id,omitempty"`

		UploadIDStr string `json:"upload_id_str,omitempty"`

		WeightedAverageWatts int64 `json:"weighted_average_watts,omitempty"`

		WorkoutType int64 `json:"workout_type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AchievementCount = dataAO1.AchievementCount

	m.Athlete = dataAO1.Athlete

	m.AthleteCount = dataAO1.AthleteCount

	m.AverageSpeed = dataAO1.AverageSpeed

	m.AverageWatts = dataAO1.AverageWatts

	m.CommentCount = dataAO1.CommentCount

	m.Commute = dataAO1.Commute

	m.DeviceWatts = dataAO1.DeviceWatts

	m.Distance = dataAO1.Distance

	m.ElapsedTime = dataAO1.ElapsedTime

	m.ElevHigh = dataAO1.ElevHigh

	m.ElevLow = dataAO1.ElevLow

	m.EndLatlng = dataAO1.EndLatlng

	m.ExternalID = dataAO1.ExternalID

	m.Flagged = dataAO1.Flagged

	m.GearID = dataAO1.GearID

	m.HasKudoed = dataAO1.HasKudoed

	m.HideFromHome = dataAO1.HideFromHome

	m.Kilojoules = dataAO1.Kilojoules

	m.KudosCount = dataAO1.KudosCount

	m.Manual = dataAO1.Manual

	m.Map = dataAO1.Map

	m.MaxSpeed = dataAO1.MaxSpeed

	m.MaxWatts = dataAO1.MaxWatts

	m.MovingTime = dataAO1.MovingTime

	m.Name = dataAO1.Name

	m.PhotoCount = dataAO1.PhotoCount

	m.Private = dataAO1.Private

	m.SportType = dataAO1.SportType

	m.StartDate = dataAO1.StartDate

	m.StartDateLocal = dataAO1.StartDateLocal

	m.StartLatlng = dataAO1.StartLatlng

	m.Timezone = dataAO1.Timezone

	m.TotalElevationGain = dataAO1.TotalElevationGain

	m.TotalPhotoCount = dataAO1.TotalPhotoCount

	m.Trainer = dataAO1.Trainer

	m.Type = dataAO1.Type

	m.UploadID = dataAO1.UploadID

	m.UploadIDStr = dataAO1.UploadIDStr

	m.WeightedAverageWatts = dataAO1.WeightedAverageWatts

	m.WorkoutType = dataAO1.WorkoutType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SummaryActivity) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MetaActivity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AchievementCount int64 `json:"achievement_count,omitempty"`

		Athlete *MetaAthlete `json:"athlete,omitempty"`

		AthleteCount int64 `json:"athlete_count,omitempty"`

		AverageSpeed float32 `json:"average_speed,omitempty"`

		AverageWatts float32 `json:"average_watts,omitempty"`

		CommentCount int64 `json:"comment_count,omitempty"`

		Commute bool `json:"commute,omitempty"`

		DeviceWatts bool `json:"device_watts,omitempty"`

		Distance float32 `json:"distance,omitempty"`

		ElapsedTime int64 `json:"elapsed_time,omitempty"`

		ElevHigh float32 `json:"elev_high,omitempty"`

		ElevLow float32 `json:"elev_low,omitempty"`

		EndLatlng LatLng `json:"end_latlng,omitempty"`

		ExternalID string `json:"external_id,omitempty"`

		Flagged bool `json:"flagged,omitempty"`

		GearID string `json:"gear_id,omitempty"`

		HasKudoed bool `json:"has_kudoed,omitempty"`

		HideFromHome bool `json:"hide_from_home,omitempty"`

		Kilojoules float32 `json:"kilojoules,omitempty"`

		KudosCount int64 `json:"kudos_count,omitempty"`

		Manual bool `json:"manual,omitempty"`

		Map *PolylineMap `json:"map,omitempty"`

		MaxSpeed float32 `json:"max_speed,omitempty"`

		MaxWatts int64 `json:"max_watts,omitempty"`

		MovingTime int64 `json:"moving_time,omitempty"`

		Name string `json:"name,omitempty"`

		PhotoCount int64 `json:"photo_count,omitempty"`

		Private bool `json:"private,omitempty"`

		SportType SportType `json:"sport_type,omitempty"`

		StartDate strfmt.DateTime `json:"start_date,omitempty"`

		StartDateLocal strfmt.DateTime `json:"start_date_local,omitempty"`

		StartLatlng LatLng `json:"start_latlng,omitempty"`

		Timezone string `json:"timezone,omitempty"`

		TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`

		TotalPhotoCount int64 `json:"total_photo_count,omitempty"`

		Trainer bool `json:"trainer,omitempty"`

		Type ActivityType `json:"type,omitempty"`

		UploadID int64 `json:"upload_id,omitempty"`

		UploadIDStr string `json:"upload_id_str,omitempty"`

		WeightedAverageWatts int64 `json:"weighted_average_watts,omitempty"`

		WorkoutType int64 `json:"workout_type,omitempty"`
	}

	dataAO1.AchievementCount = m.AchievementCount

	dataAO1.Athlete = m.Athlete

	dataAO1.AthleteCount = m.AthleteCount

	dataAO1.AverageSpeed = m.AverageSpeed

	dataAO1.AverageWatts = m.AverageWatts

	dataAO1.CommentCount = m.CommentCount

	dataAO1.Commute = m.Commute

	dataAO1.DeviceWatts = m.DeviceWatts

	dataAO1.Distance = m.Distance

	dataAO1.ElapsedTime = m.ElapsedTime

	dataAO1.ElevHigh = m.ElevHigh

	dataAO1.ElevLow = m.ElevLow

	dataAO1.EndLatlng = m.EndLatlng

	dataAO1.ExternalID = m.ExternalID

	dataAO1.Flagged = m.Flagged

	dataAO1.GearID = m.GearID

	dataAO1.HasKudoed = m.HasKudoed

	dataAO1.HideFromHome = m.HideFromHome

	dataAO1.Kilojoules = m.Kilojoules

	dataAO1.KudosCount = m.KudosCount

	dataAO1.Manual = m.Manual

	dataAO1.Map = m.Map

	dataAO1.MaxSpeed = m.MaxSpeed

	dataAO1.MaxWatts = m.MaxWatts

	dataAO1.MovingTime = m.MovingTime

	dataAO1.Name = m.Name

	dataAO1.PhotoCount = m.PhotoCount

	dataAO1.Private = m.Private

	dataAO1.SportType = m.SportType

	dataAO1.StartDate = m.StartDate

	dataAO1.StartDateLocal = m.StartDateLocal

	dataAO1.StartLatlng = m.StartLatlng

	dataAO1.Timezone = m.Timezone

	dataAO1.TotalElevationGain = m.TotalElevationGain

	dataAO1.TotalPhotoCount = m.TotalPhotoCount

	dataAO1.Trainer = m.Trainer

	dataAO1.Type = m.Type

	dataAO1.UploadID = m.UploadID

	dataAO1.UploadIDStr = m.UploadIDStr

	dataAO1.WeightedAverageWatts = m.WeightedAverageWatts

	dataAO1.WorkoutType = m.WorkoutType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this summary activity
func (m *SummaryActivity) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetaActivity
	if err := m.MetaActivity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAthlete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAthleteCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndLatlng(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSportType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartLatlng(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryActivity) validateAthlete(formats strfmt.Registry) error {

	if swag.IsZero(m.Athlete) { // not required
		return nil
	}

	if m.Athlete != nil {
		if err := m.Athlete.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

func (m *SummaryActivity) validateAthleteCount(formats strfmt.Registry) error {

	if swag.IsZero(m.AthleteCount) { // not required
		return nil
	}

	if err := validate.MinimumInt("athlete_count", "body", m.AthleteCount, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *SummaryActivity) validateEndLatlng(formats strfmt.Registry) error {

	if swag.IsZero(m.EndLatlng) { // not required
		return nil
	}

	if err := m.EndLatlng.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("end_latlng")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("end_latlng")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) validateMap(formats strfmt.Registry) error {

	if swag.IsZero(m.Map) { // not required
		return nil
	}

	if m.Map != nil {
		if err := m.Map.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *SummaryActivity) validateSportType(formats strfmt.Registry) error {

	if swag.IsZero(m.SportType) { // not required
		return nil
	}

	if err := m.SportType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sport_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sport_type")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SummaryActivity) validateStartDateLocal(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDateLocal) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date_local", "body", "date-time", m.StartDateLocal.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SummaryActivity) validateStartLatlng(formats strfmt.Registry) error {

	if swag.IsZero(m.StartLatlng) { // not required
		return nil
	}

	if err := m.StartLatlng.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start_latlng")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("start_latlng")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this summary activity based on the context it is used
func (m *SummaryActivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetaActivity
	if err := m.MetaActivity.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAthlete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndLatlng(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSportType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartLatlng(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryActivity) contextValidateAthlete(ctx context.Context, formats strfmt.Registry) error {

	if m.Athlete != nil {
		if err := m.Athlete.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("athlete")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("athlete")
			}
			return err
		}
	}

	return nil
}

func (m *SummaryActivity) contextValidateEndLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.EndLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("end_latlng")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("end_latlng")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) contextValidateMap(ctx context.Context, formats strfmt.Registry) error {

	if m.Map != nil {
		if err := m.Map.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("map")
			}
			return err
		}
	}

	return nil
}

func (m *SummaryActivity) contextValidateSportType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SportType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sport_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sport_type")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) contextValidateStartLatlng(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StartLatlng.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("start_latlng")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("start_latlng")
		}
		return err
	}

	return nil
}

func (m *SummaryActivity) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryActivity) UnmarshalBinary(b []byte) error {
	var res SummaryActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
