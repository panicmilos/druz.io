package api_contracts

import (
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"

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
