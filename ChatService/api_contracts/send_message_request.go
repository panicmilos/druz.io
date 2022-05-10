package api_contracts

import (
	goerrors "errors"
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/panicmilos/druz.io/ChatService/errors"
)

type SendMessageRequest struct {
	Message string
	Type    string
}

func (request *SendMessageRequest) Validate() error {
	err := validation.ValidateStruct(request,
		validation.Field(&request.Type,
			validation.In("Text", "Image").Error("Type must be Text or Image"),
		),
		validation.Field(&request.Message,
			validation.By(func(value interface{}) error {
				message := value.(string)
				if request.Type == "Text" && len(strings.TrimSpace(message)) == 0 {
					return goerrors.New("Text message must be provided")
				}

				matched, _ := regexp.MatchString(`^(http(s)?:\/\/)[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`, message)
				if request.Type == "Image" && !matched {
					return goerrors.New("Image must be valid url")
				}

				return nil
			}),
		),
	)

	if err == nil {
		return nil
	}

	return errors.NewErrValidation(err.Error())
}
