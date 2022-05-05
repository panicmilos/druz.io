package api_contracts

import (
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (request *LoginRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email,
			*helpers.ValidateEmail()...,
		),
		validation.Field(&request.Password,
			*helpers.ValidatePassword()...,
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
