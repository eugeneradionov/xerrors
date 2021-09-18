package xerrors // nolint:dupl

import "strings"

type XErrors interface {
	error

	// Add adds XError to collection
	Add(...XError)
	// GetErrors returns errors collection
	GetErrors() []XError
	// Len returns length of errors collection
	Len() int
	// Sanitize sanitizes errors collection
	Sanitize()
}

func NewXErrs() *XErrs {
	return NewXErrsWithLen(0, 1)
}

func NewXErrsWithLen(l, c int) *XErrs {
	return &XErrs{Errs: make([]XError, l, c)}
}

type XErrs struct {
	Errs []XError `json:"errors"`
}

func (errs *XErrs) Error() string {
	if errs == nil {
		return ""
	}

	errors := make([]string, len(errs.Errs))

	for i := range errs.Errs {
		errors[i] = errs.Errs[i].Error()
	}

	return strings.Join(errors, ";")
}

func (errs *XErrs) Add(xerrs ...XError) {
	if errs != nil {
		errs.Errs = append(errs.Errs, xerrs...)
	}
}

func (errs *XErrs) GetErrors() []XError {
	if errs == nil {
		return nil
	}

	return errs.Errs
}

func (errs *XErrs) Len() int {
	if errs == nil {
		return 0
	}

	return len(errs.Errs)
}

func (errs *XErrs) Sanitize() {
	if errs == nil {
		return
	}

	for i := range errs.Errs {
		errs.Errs[i].Sanitize()
	}
}
