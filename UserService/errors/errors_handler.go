package errors

import (
	"UserService/helpers"
	"net/http"
)

func Handle(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}

	switch e := err.(type) {
	case *ErrValidation:
		helpers.JSONResponse(w, 400, map[string]interface{}{
			"error": e.Error(),
		})
	case *ErrBadRequest:
		helpers.JSONResponse(w, 400, map[string]interface{}{
			"error": e.Error(),
		})
	case *ErrNotFound:
		helpers.JSONResponse(w, 404, map[string]interface{}{
			"error": e.Error(),
		})
	default:
		helpers.JSONResponse(w, 500, map[string]interface{}{
			"error": "Internal Error",
		})
	}

	return true
}
