// Code generated by goa v3.21.1, DO NOT EDIT.
//
// profiles HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/design/api
// -o ./services/profiles/

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	profiles "github.com/ynoacamino/infra-sustainable-classrooms/services/profiles/gen/profiles"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateStudentProfilePayload builds the payload for the profiles
// CreateStudentProfile endpoint from CLI flags.
func BuildCreateStudentProfilePayload(profilesCreateStudentProfileBody string, profilesCreateStudentProfileSessionToken string) (*profiles.CreateStudentProfilePayload, error) {
	var err error
	var body CreateStudentProfileRequestBody
	{
		err = json.Unmarshal([]byte(profilesCreateStudentProfileBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"avatar_url\": \"Nemo corporis ratione quo.\",\n      \"bio\": \"Qui ut aut et facilis excepturi.\",\n      \"email\": \"lera@farrelljast.name\",\n      \"first_name\": \"Dolorem hic sapiente ea.\",\n      \"grade_level\": \"Harum cum velit laboriosam tenetur maxime maxime.\",\n      \"last_name\": \"Consequatur id facilis eaque incidunt.\",\n      \"major\": \"Voluptatem ab.\",\n      \"phone\": \"Totam laboriosam veniam.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if err != nil {
			return nil, err
		}
	}
	var sessionToken string
	{
		sessionToken = profilesCreateStudentProfileSessionToken
	}
	v := &profiles.CreateStudentProfilePayload{
		FirstName:  body.FirstName,
		LastName:   body.LastName,
		Email:      body.Email,
		Phone:      body.Phone,
		AvatarURL:  body.AvatarURL,
		Bio:        body.Bio,
		GradeLevel: body.GradeLevel,
		Major:      body.Major,
	}
	v.SessionToken = sessionToken

	return v, nil
}

// BuildCreateTeacherProfilePayload builds the payload for the profiles
// CreateTeacherProfile endpoint from CLI flags.
func BuildCreateTeacherProfilePayload(profilesCreateTeacherProfileBody string, profilesCreateTeacherProfileSessionToken string) (*profiles.CreateTeacherProfilePayload, error) {
	var err error
	var body CreateTeacherProfileRequestBody
	{
		err = json.Unmarshal([]byte(profilesCreateTeacherProfileBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"avatar_url\": \"Omnis alias dolor voluptatibus occaecati nemo.\",\n      \"bio\": \"Qui consectetur perferendis tempore porro ut velit.\",\n      \"email\": \"rosalinda@osinskigoodwin.com\",\n      \"first_name\": \"Quaerat officia ut velit odit architecto.\",\n      \"last_name\": \"Voluptatibus voluptas sequi voluptatibus.\",\n      \"phone\": \"Voluptatem adipisci in.\",\n      \"position\": \"Quaerat amet et consectetur.\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))
		if err != nil {
			return nil, err
		}
	}
	var sessionToken string
	{
		sessionToken = profilesCreateTeacherProfileSessionToken
	}
	v := &profiles.CreateTeacherProfilePayload{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
		AvatarURL: body.AvatarURL,
		Bio:       body.Bio,
		Position:  body.Position,
	}
	v.SessionToken = sessionToken

	return v, nil
}

// BuildGetCompleteProfilePayload builds the payload for the profiles
// GetCompleteProfile endpoint from CLI flags.
func BuildGetCompleteProfilePayload(profilesGetCompleteProfileSessionToken string) (*profiles.GetCompleteProfilePayload, error) {
	var sessionToken string
	{
		sessionToken = profilesGetCompleteProfileSessionToken
	}
	v := &profiles.GetCompleteProfilePayload{}
	v.SessionToken = sessionToken

	return v, nil
}

// BuildGetPublicProfileByIDPayload builds the payload for the profiles
// GetPublicProfileById endpoint from CLI flags.
func BuildGetPublicProfileByIDPayload(profilesGetPublicProfileByIDUserID string) (*profiles.GetPublicProfileByIDPayload, error) {
	var err error
	var userID int64
	{
		userID, err = strconv.ParseInt(profilesGetPublicProfileByIDUserID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for userID, must be INT64")
		}
	}
	v := &profiles.GetPublicProfileByIDPayload{}
	v.UserID = userID

	return v, nil
}

// BuildUpdateProfilePayload builds the payload for the profiles UpdateProfile
// endpoint from CLI flags.
func BuildUpdateProfilePayload(profilesUpdateProfileBody string, profilesUpdateProfileSessionToken string) (*profiles.UpdateProfilePayload, error) {
	var err error
	var body UpdateProfileRequestBody
	{
		err = json.Unmarshal([]byte(profilesUpdateProfileBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"avatar_url\": \"Qui alias provident vel minima saepe delectus.\",\n      \"bio\": \"In ut id aut fuga blanditiis.\",\n      \"email\": \"aglae.hudson@mosciski.name\",\n      \"first_name\": \"Repudiandae rem quia.\",\n      \"last_name\": \"Consequatur aperiam ad suscipit aut distinctio et.\",\n      \"major\": \"Error rerum.\",\n      \"phone\": \"Magni itaque non quae assumenda.\",\n      \"position\": \"Est quaerat commodi.\"\n   }'")
		}
		if body.Email != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
		}
		if err != nil {
			return nil, err
		}
	}
	var sessionToken string
	{
		sessionToken = profilesUpdateProfileSessionToken
	}
	v := &profiles.UpdateProfilePayload{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
		AvatarURL: body.AvatarURL,
		Bio:       body.Bio,
		Major:     body.Major,
		Position:  body.Position,
	}
	v.SessionToken = sessionToken

	return v, nil
}
