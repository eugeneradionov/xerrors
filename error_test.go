package xerrors

import (
	"reflect"
	"testing"
)

func TestErr_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  *Err
		want string
	}{
		{
			name: "nil error",
			err:  nil,
			want: "nil error",
		},
		{
			name: "err with message",
			err: &Err{
				Message:     "test message",
				Description: "",
			},
			want: "test message:",
		},
		{
			name: "err with message and description",
			err: &Err{
				Message:     "test message",
				Description: "test description",
			},
			want: "test message:test description",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErr_Sanitize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  *Err
	}{
		{
			name: "Err with description",
			err: &Err{
				Message:     "test message",
				Description: "test description",
			},
		},
		{
			name: "nil Err",
			err:  nil,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.err.Sanitize()

			if tt.err != nil && tt.err.Description != "" {
				t.Errorf("Sanitize() want description to be empty, go: %v", tt.err.Description)
			}
		})
	}
}

func TestNewErr(t *testing.T) {
	t.Parallel()

	type args struct {
		msg   string
		descr string
	}
	tests := []struct {
		name string
		args args
		want *Err
	}{
		{
			name: "new error",
			args: args{
				msg:   "test message",
				descr: "test description",
			},
			want: &Err{
				Message:     "test message",
				Description: "test description",
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewErr(tt.args.msg, tt.args.descr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
