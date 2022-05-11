package helpers

import (
	"github.com/panicmilos/druz.io/ChatService/dto"
)

func ContainsChat(s []dto.Chat, e dto.Chat) bool {
	for _, a := range s {
		if a.Chat == e.Chat {
			return true
		}
	}
	return false
}
