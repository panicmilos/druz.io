package api_contracts

import (
	"time"

	"UserService/errors"
	"UserService/helpers"
	"UserService/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateAccountRequest struct {
	Email    string
	Password string

	Profile CreateProfileRequest
}

type CreateProfileRequest struct {
	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    models.Gender
}

func (request *CreateAccountRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email,
			*helpers.ValidateEmail()...,
		),
		validation.Field(&request.Password,
			*helpers.ValidatePassword()...,
		),
		helpers.NestedValidation(&request.Profile,
			validation.Field(&request.Profile.FirstName,
				validation.Required.Error("First name must be provided"),
			),
			validation.Field(&request.Profile.LastName,
				validation.Required.Error("Last name must be provided"),
			),
			validation.Field(&request.Profile.Birthday,
				validation.Required.Error("Birthdate must be provided"),
				validation.By(func(value interface{}) error {
					birthday := value.(time.Time).AddDate(13, 0, 0)
					if birthday.After(time.Now()) {
						return errors.NewErrValidation("You have to be at least 13 years old")
					}
					return nil
				}),
			),
			validation.Field(&request.Profile.Gender,
				validation.Min(0).Error("Gender must be >= 0"),
				validation.Max(2).Error("Gender must be <= 2"),
			),
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
