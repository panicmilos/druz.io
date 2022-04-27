package api_contracts

import (
	"UserService/errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ReportUserRequest struct {
	Reason string
}

func (request *ReportUserRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Reason,
			validation.Required.Error("Reason must be provided"),
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
