package api_contracts

import (
	"UserService/errors"
	"UserService/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RecoverPasswordRequest struct {
	NewPassword string
}

func (request *RecoverPasswordRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.NewPassword,
			*helpers.ValidatePassword()...,
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
