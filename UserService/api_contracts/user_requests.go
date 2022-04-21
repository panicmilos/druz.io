package api_contracts

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUserRequest struct {
	Username string
	Password string
}

func (request *CreateUserRequest) Validate() error {
	return validation.ValidateStruct(request,
		validation.Field(&request.Username, validation.Required.Error("Username must be provided"), validation.Length(5, 35).Error("Username must have at least 5 characters")),
	)
}
