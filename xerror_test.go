// nolint:dupl
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
