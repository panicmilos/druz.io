package models

import "github.com/panicmilos/druz.io/UserService/errors"

type Gender int64

const (
	Male Gender = iota
	Female
	Other
)

func GenderFromString(value string) (Gender, error) {
	switch value {
	case "0":
		return Male, nil
	case "1":
		return Female, nil
	case "2":
		return Other, nil
	default:
		return Other, errors.NewErrBadRequest("Bad Gender")
	}
}
