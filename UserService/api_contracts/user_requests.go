package api_contracts

import (
	"regexp"
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
			validation.Match(regexp.MustCompile(`^([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`)).Error("Email must be valid")),
		validation.Field(&request.Password,
			validation.Length(14, 30).Error("Password should have atleast 14 characters"),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("Password should contain capital letters"),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("Password should contain lower letters"),
			validation.Match(regexp.MustCompile(`[0-9]`)).Error("Password should contain numbers"),
			validation.Match(regexp.MustCompile(`[^a-zA-Z0-9]`)).Error("Password should contain numbers"),
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
