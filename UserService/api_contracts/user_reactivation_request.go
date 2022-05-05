package api_contracts

import (
	"github.com/panicmilos/druz.io/UserService/errors"
	"github.com/panicmilos/druz.io/UserService/helpers"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserReactivationRequest struct {
	Email string
}

func (request *UserReactivationRequest) Validate() error {
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
