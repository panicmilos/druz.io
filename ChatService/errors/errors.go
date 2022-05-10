package errors

type ErrValidation struct {
	msg string
}

func NewErrValidation(msg string) *ErrValidation {
	return &ErrValidation{msg: msg}
}

func (err *ErrValidation) Error() string {
	return err.msg
}

type ErrNotFound struct {
	msg string
}

func NewErrNotFound(msg string) *ErrNotFound {
	return &ErrNotFound{msg: msg}
}

func (err *ErrNotFound) Error() string {
	return err.msg
}

type ErrBadRequest struct {
	msg string
}

func NewErrBadRequest(msg string) *ErrBadRequest {
	return &ErrBadRequest{msg: msg}
}

func (err *ErrBadRequest) Error() string {
	return err.msg
}

type ErrUnauthorized struct {
	msg string
}

func NewErrUnauthorized(msg string) *ErrUnauthorized {
	return &ErrUnauthorized{msg: msg}
}

func (err *ErrUnauthorized) Error() string {
	return err.msg
}

type ErrForbidden struct {
	msg string
}

func NewErrForbidden(msg string) *ErrForbidden {
	return &ErrForbidden{msg: msg}
}

func (err *ErrForbidden) Error() string {
	return err.msg
}
