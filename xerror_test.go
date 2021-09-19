// nolint:dupl,funlen
package xerrors

import (
	"reflect"
	"testing"
)

func TestXErr_GetDescription(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
		want string
	}{
		{
			name: "nil XErr",
			xErr: nil,
			want: "",
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:       "",
				Description:   "test description",
				Extra:         nil,
				InternalExtra: nil,
			},
			want: "test description",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.xErr.GetDescription(); got != tt.want {
				t.Errorf("GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErr_GetExtra(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
		want map[string]interface{}
	}{
		{
			name: "nil XErr",
			xErr: nil,
			want: nil,
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:     "test message",
				Description: "test description",
				Extra: map[string]interface{}{
					"hello": "world",
				},
				InternalExtra: nil,
			},
			want: map[string]interface{}{
				"hello": "world",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.xErr.GetExtra(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExtra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErr_GetInternalExtra(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
		want map[string]interface{}
	}{
		{
			name: "nil XErr",
			xErr: nil,
			want: nil,
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:     "test message",
				Description: "test description",
				Extra:       nil,
				InternalExtra: map[string]interface{}{
					"hello": "world",
				},
			},
			want: map[string]interface{}{
				"hello": "world",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.xErr.GetInternalExtra(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInternalExtra() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErr_GetMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
		want string
	}{
		{
			name: "nil XErr",
			xErr: nil,
			want: "",
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:       "test message",
				Description:   "test description",
				Extra:         nil,
				InternalExtra: nil,
			},
			want: "test message",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.xErr.GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErr_Sanitize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
	}{
		{
			name: "nil XErr",
			xErr: (*XErr)(nil),
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:       "test message",
				Description:   "test description",
				Extra:         nil,
				InternalExtra: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.xErr.Sanitize()
			if tt.xErr != nil && tt.xErr.Description != "" {
				t.Errorf("Sanitize() want empty description, got: %v", tt.xErr.Description)
			}
		})
	}
}

func TestXErr_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		xErr *XErr
		want string
	}{
		{
			name: "nil XErr",
			xErr: (*XErr)(nil),
			want: "",
		},
		{
			name: "not nil XErr",
			xErr: &XErr{
				Message:       "test message",
				Description:   "test description",
				Extra:         map[string]interface{}{"filed": "user", "user_id": 123},
				InternalExtra: map[string]interface{}{"error_info": "connect to db"},
			},
			want: "test message: test description; map[filed:user user_id:123]",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.xErr
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		msg  string
		opts []XErrOpt
	}

	tests := []struct {
		name string
		args args
		want *XErr
	}{
		{
			name: "new error without options",
			args: args{
				msg: "some error",
			},
			want: &XErr{
				Message: "some error",
			},
		},
		{
			name: "new error with options",
			args: args{
				msg: "some error",
				opts: []XErrOpt{
					WithDescription("error description"),
					WithExtra(map[string]interface{}{"filed": "user", "user_id": 123}),
					WithInternalExtra(map[string]interface{}{"error_info": "connect to db"}),
				},
			},
			want: &XErr{
				Message:       "some error",
				Description:   "error description",
				Extra:         map[string]interface{}{"filed": "user", "user_id": 123},
				InternalExtra: map[string]interface{}{"error_info": "connect to db"},
			},
		},
		{
			name: "new error with options overwrite",
			args: args{
				msg: "some error",
				opts: []XErrOpt{
					WithDescription("error description"),
					WithExtra(map[string]interface{}{"filed": "user", "user_id": 123}),
					WithInternalExtra(map[string]interface{}{"error_info": "connect to db"}),
					WithMessage("new message"),
					WithDescription("new description"),
				},
			},
			want: &XErr{
				Message:       "new message",
				Description:   "new description",
				Extra:         map[string]interface{}{"filed": "user", "user_id": 123},
				InternalExtra: map[string]interface{}{"error_info": "connect to db"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := New(tt.args.msg, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
