package api_contracts

import (
	"UserService/errors"
	"UserService/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
)

type PasswordRecoveryRequest struct {
	Email string
}

func (request *PasswordRecoveryRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Email,
			*helpers.ValidateEmail()...,
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
