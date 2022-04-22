package api_contracts

import (
	"UserService/errors"
	"UserService/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ChangePasswordRequest struct {
	CurrentPassword string
	NewPassword     string
}

func (request *ChangePasswordRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.CurrentPassword,
			*helpers.ValidatePassword()...,
		),
		validation.Field(&request.NewPassword,
			*helpers.ValidatePassword()...,
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
