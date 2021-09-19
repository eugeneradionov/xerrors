// nolint:dupl,goerr113,funlen
package xhttp

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/eugeneradionov/xerrors"
)

func TestNewBadRequestError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Bad Request",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusBadRequest},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusBadRequest},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewBadRequestError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBadRequestError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		msg  string
		code int
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("some error"),
				msg:  "db connection failed",
				code: http.StatusInternalServerError,
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "db connection failed",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusInternalServerError},
				InternalExtra: map[string]interface{}{"error": errors.New("some error")},
			},
		},
		{
			name: "error with options",
			args: args{
				err:  errors.New("some error"),
				msg:  "db connection failed",
				code: http.StatusInternalServerError,
				opts: []xerrors.XErrOpt{
					xerrors.WithDescription("description"),
				},
			},
			want: &xerrors.XErr{
				Message:       "db connection failed",
				Description:   "description",
				Extra:         map[string]interface{}{"http_code": http.StatusInternalServerError},
				InternalExtra: map[string]interface{}{"error": errors.New("some error")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewError(tt.args.err, tt.args.msg, tt.args.code, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewForbiddenError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Forbidden",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusForbidden},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusForbidden},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewForbiddenError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForbiddenError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalServerError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Internal Server Error",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusInternalServerError},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusInternalServerError},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewInternalServerError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNotFoundError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Not Found",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusNotFound},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusNotFound},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewNotFoundError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Unauthorized",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusUnauthorized},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusUnauthorized},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewUnauthorizedError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnauthorizedError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnprocessableEntityError(t *testing.T) {
	t.Parallel()

	type args struct {
		err  error
		opts []xerrors.XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *xerrors.XErr
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "error without options",
			args: args{
				err:  errors.New("db connection failed"),
				opts: nil,
			},
			want: &xerrors.XErr{
				Message:       "Unprocessable Entity",
				Description:   "",
				Extra:         map[string]interface{}{"http_code": http.StatusUnprocessableEntity},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
		{
			name: "error with options",
			args: args{
				err: errors.New("db connection failed"),
				opts: []xerrors.XErrOpt{
					xerrors.WithMessage("rewrite message"),
					xerrors.WithDescription("db connection failed"),
				},
			},
			want: &xerrors.XErr{
				Message:       "rewrite message",
				Description:   "db connection failed",
				Extra:         map[string]interface{}{"http_code": http.StatusUnprocessableEntity},
				InternalExtra: map[string]interface{}{"error": errors.New("db connection failed")},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewUnprocessableEntityError(tt.args.err, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnprocessableEntityError() = %v, want %v", got, tt.want)
			}
		})
	}
}
