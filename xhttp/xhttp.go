package xhttp

import (
	"net/http"

	"github.com/eugeneradionov/xerrors"
)

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

func NewBadRequestError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Bad Request", http.StatusBadRequest, opts...)
}

func NewUnauthorizedError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Unauthorized", http.StatusUnauthorized, opts...)
}

func NewForbiddenError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Forbidden", http.StatusForbidden, opts...)
}

func NewNotFoundError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Not Found", http.StatusNotFound, opts...)
}

func NewUnprocessableEntityError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Unprocessable Entity", http.StatusUnprocessableEntity, opts...)
}

func NewInternalServerError(err error, opts ...xerrors.XErrOpt) *xerrors.XErr {
	return NewError(err, "Internal Server Error", http.StatusInternalServerError, opts...)
}
