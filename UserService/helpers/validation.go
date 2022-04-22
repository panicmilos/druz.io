package helpers

import (
	"errors"
	"reflect"
	"regexp"
	"time"

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

func ValidateEmail() *[]validation.Rule {
	return &[]validation.Rule{
		validation.Match(regexp.MustCompile(`^([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`)).Error("Email must be valid"),
	}
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

func ValidateBirthday() *[]validation.Rule {
	return &[]validation.Rule{
		validation.Required.Error("Birthdate must be provided"),
		validation.By(func(value interface{}) error {
			birthday := value.(time.Time).AddDate(13, 0, 0)
			if birthday.After(time.Now()) {
				return errors.New("You have to be at least 13 years old")
			}
			return nil
		}),
	}
}

func ValidateGender() *[]validation.Rule {
	return &[]validation.Rule{
		validation.Min(0).Error("Gender must be >= 0"),
		validation.Max(2).Error("Gender must be <= 2"),
	}
}
