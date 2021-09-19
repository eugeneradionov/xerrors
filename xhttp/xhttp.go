package xhttp

import (
	"net/http"

	"github.com/eugeneradionov/xerrors"
)

// NewError creates new HTTP XErr with following structure:
// message: msg, extra: {"http_code": code}, internal_extra: {"error": err}.
func NewError(err error, msg string, code int, opts ...xerrors.XErrOpt) *xerrors.XErr {
	if err == nil {
		return nil
	}

	opts = append([]xerrors.XErrOpt{
		xerrors.WithExtra(map[string]interface{}{"http_code": code}),
		xerrors.WithInternalExtra(map[string]interface{}{"error": err}),
	}, opts...)

	return xerrors.New(msg, opts...)
}

// NewBadRequestError creates new HTTP BadRequest(400) error.
func NewBadRequestError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Bad Request", http.StatusBadRequest, opts...)
}

// NewUnauthorizedError creates new HTTP Unauthorized(401) error.
func NewUnauthorizedError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Unauthorized", http.StatusUnauthorized, opts...)
}

// NewForbiddenError creates new HTTP Forbidden(403) error.
func NewForbiddenError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Forbidden", http.StatusForbidden, opts...)
}

// NewNotFoundError creates new HTTP NotFound(404) error.
func NewNotFoundError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Not Found", http.StatusNotFound, opts...)
}

// NewUnprocessableEntityError creates new HTTP UnprocessableEntity(422) error.
func NewUnprocessableEntityError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Unprocessable Entity", http.StatusUnprocessableEntity, opts...)
}

// NewInternalServerError creates new HTTP InternalServerError(500) error.
func NewInternalServerError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Internal Server Error", http.StatusInternalServerError, opts...)
}
