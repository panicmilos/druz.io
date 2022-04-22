package helpers

import (
	"reflect"

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
