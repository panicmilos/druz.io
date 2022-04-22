package api_contracts

import (
	"UserService/errors"
	"UserService/helpers"
	"UserService/models"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateProfileRequest struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    models.Gender

	About       string
	PhoneNumber string
	LivePlaces  []LivePlaceRequest
	WorkPlaces  []WorkPlaceRequest
	Educations  []EducationRequest
	Intereses   []InteresRequest
}

type LivePlaceRequest struct {
	Place         string
	LivePlaceType models.LivePlaceType
}

type WorkPlaceRequest struct {
	Place string
	From  time.Time
	To    time.Time
}

type EducationRequest struct {
	Place string
	From  time.Time
	To    time.Time
}

type InteresRequest struct {
	Interes string
}

func (request *UpdateProfileRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.FirstName,
			validation.Required.Error("First name must be provided"),
		),
		validation.Field(&request.LastName,
			validation.Required.Error("First name must be provided"),
		),
		validation.Field(&request.Birthday,
			*helpers.ValidateBirthday()...,
		),
		validation.Field(&request.Gender,
			*helpers.ValidateGender()...,
		),
		validation.Field(&request.PhoneNumber,
			validation.Match(regexp.MustCompile(`^(\+\d{1,3}\s?)?((\(\d{3}\)\s?)|(\d{3})(\s|-?))(\d{3}(\s|-?))(\d{4})(\s?(([E|e]xt[:|.|]?)|x|X)(\s?\d+))?$`)).Error("Phone number must be valid"),
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
