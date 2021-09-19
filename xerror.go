package xerrors

import "fmt"

// Error - base interface that contains minimum amount of required functions.
type Error interface {
	error

	// Sanitize removes sensitive information from Error.
	Sanitize()
	// GetMessage returns general error message.
	GetMessage() string
	// GetDescription returns detailed, probably sensitive, error description.
	GetDescription() string
}

// XError - extended interface with extra and internal extra.
type XError interface {
	Error

	// GetExtra returns public extra info.
	GetExtra() map[string]interface{}
	// GetInternalExtra returns private extra info.
	GetInternalExtra() map[string]interface{}
}

// XErr represents extended error.
type XErr struct {
	// Message contains general error message.
	Message string `json:"message,omitempty"`
	// Description contains detailed error description.
	Description string `json:"description,omitempty"`

	// Extra contains any public extra info that can be sent in the response.
	Extra map[string]interface{} `json:"extra,omitempty"`
	// InternalExtra contains private extra info that could be helpful for internal usage
	// and shouldn't be sent to external users.
	InternalExtra map[string]interface{} `json:"-"`
}

// XErrOpt represents option for XErr constructor New.
type XErrOpt func(err *XErr)

func WithMessage(msg string) XErrOpt                 { return func(err *XErr) { err.Message = msg } }
func WithDescription(descr string) XErrOpt           { return func(err *XErr) { err.Description = descr } }
func WithExtra(extra map[string]interface{}) XErrOpt { return func(err *XErr) { err.Extra = extra } }
func WithInternalExtra(extra map[string]interface{}) XErrOpt {
	return func(err *XErr) { err.InternalExtra = extra }
}

// New - constructor for XErr with options, returns new *XErr.
func New(msg string, opts ...XErrOpt) *XErr {
	err := &XErr{
		Message: msg,
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

// NewXErr returns new XErr.
func NewXErr(msg, descr string, extra, intExtra map[string]interface{}) *XErr {
	return New(msg, WithDescription(descr), WithExtra(extra), WithInternalExtra(intExtra))
}

// Error unifying XErr with Go error interface.
func (err *XErr) Error() string {
	if err == nil {
		return ""
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
