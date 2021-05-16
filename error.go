package xerrors

import (
	"fmt"
)

type Error interface {
	error

	// Sanitize removes description from XError
	Sanitize()
}

type Err struct {
	// Message contains general error message.
	Message string `json:"message,omitempty"`
	// Description contains detailed, probably sensitive, error description
	// and should be sanitized before marshaling.
	Description string `json:"description,omitempty"`
}

func NewErr(msg, descr string) *Err {
	return &Err{
		Message:     msg,
		Description: descr,
	}
}

func (err *Err) Error() string {
	if err == nil {
		return "nil error"
	}

	return fmt.Sprintf("%s:%s", err.Message, err.Description)
}

func (err *Err) Sanitize() {
	if err == nil {
		return
	}

	err.Description = ""
}
