package helpers

import (
	"reflect"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

func NestedValidation(target interface{}, fieldRules ...*validation.FieldRules) *validation.FieldRules {
	return validation.Field(target, validation.By(func(value interface{}) error {
		valueV := reflect.Indirect(reflect.ValueOf(value))
		if valueV.CanAddr() {
			addr := valueV.Addr().Interface()
			return validation.ValidateStruct(addr, fieldRules...)
		}
		return validation.ValidateStruct(target, fieldRules...)
	}))
}

func ValidatePassword() *[]validation.Rule {
	return &[]validation.Rule{
		validation.Length(14, 30).Error("Password should have atleast 14 characters"),
		validation.Match(regexp.MustCompile(`[A-Z]`)).Error("Password should contain capital letters"),
		validation.Match(regexp.MustCompile(`[a-z]`)).Error("Password should contain lower letters"),
		validation.Match(regexp.MustCompile(`[0-9]`)).Error("Password should contain numbers"),
		validation.Match(regexp.MustCompile(`[^a-zA-Z0-9]`)).Error("Password should contain special character"),
	}
}
