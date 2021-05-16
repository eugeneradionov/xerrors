package xerrors // nolint:dupl

import "strings"

type Errors interface {
	error

	// Add adds Error to collection
	Add(...Error)
	// GetErrors returns errors collection
	GetErrors() []Error
	// Len returns length of errors collection
	Len() int
	// Sanitize sanities errors collection
	Sanitize()
}

func NewErrs() *Errs {
	return NewErrsWithLen(0, 1)
}

func NewErrsWithLen(l, c int) *Errs {
	return &Errs{Errors: make([]Error, l, c)}
}

type Errs struct {
	Errors []Error `json:"errors"`
}

func (errs *Errs) Error() string {
	if errs == nil {
		return ""
	}

	var errors = make([]string, len(errs.Errors))

	for i := range errs.Errors {
		errors[i] = errs.Errors[i].Error()
	}

	return strings.Join(errors, "; ")
}

func (errs *Errs) Add(err ...Error) {
	if errs != nil {
		errs.Errors = append(errs.Errors, err...)
	}
}

func (errs *Errs) GetErrors() []Error {
	if errs == nil {
		return nil
	}

	return errs.Errors
}

func (errs *Errs) Len() int {
	if errs == nil {
		return 0
	}

	return len(errs.Errors)
}

func (errs *Errs) Sanitize() {
	if errs == nil {
		return
	}

	for i := range errs.Errors {
		errs.Errors[i].Sanitize()
	}
}
