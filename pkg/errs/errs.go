package errs

import "net/http"

type Err struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e Err) Error() string {
	return e.Message
}

func (e Err) SetStatus(status string) Err {
	e.Status = status
	return e
}

func Wrap(code int, err error) Err {
	return Err{
		Code:    code,
		Status:  http.StatusText(code),
		Message: err.Error(),
	}
}
