package xerrors

import "fmt"

type XError interface {
	Error

	// GetMessage returns general error message.
	GetMessage() string
	// GetDescription returns detailed, probably sensitive, error description
	GetDescription() string
	// GetExtra returns public extra info
	GetExtra() map[string]interface{}
	// GetInternalExtra returns private extra info
	GetInternalExtra() map[string]interface{}
}

type XErr struct {
	*Err

	// Extra contains any public extra info that should be sent in the response
	Extra map[string]interface{} `json:"extra,omitempty"`
	// InternalExtra contains private extra info that could be helpful for internal usage
	// and shouldn't shown to external users.
	InternalExtra map[string]interface{} `json:"-"`
}

// NewXErr returns new XErr
func NewXErr(msg, descr string, extra, intExtra map[string]interface{}) *XErr {
	return &XErr{
		Err:           NewErr(msg, descr),
		Extra:         extra,
		InternalExtra: intExtra,
	}
}

// Error unifying XError with Go error interface.
func (err *XErr) Error() string {
	if err == nil {
		return "nil error"
	}

	return fmt.Sprintf("%s: %s; %v", err.Message, err.Description, err.Extra)
}

func (err *XErr) Sanitize() {
	if err == nil {
		return
	}

	err.Description = ""
}

func (err *XErr) GetMessage() string {
	if err == nil {
		return ""
	}

	return err.Message
}

func (err *XErr) GetDescription() string {
	if err == nil {
		return ""
	}

	return err.Description
}

func (err *XErr) GetExtra() map[string]interface{} {
	if err == nil {
		return nil
	}

	return err.Extra
}

func (err *XErr) GetInternalExtra() map[string]interface{} {
	if err == nil {
		return nil
	}

	return err.InternalExtra
}
