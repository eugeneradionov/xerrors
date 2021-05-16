package xerrors

import (
	"reflect"
	"testing"
)

func TestErrs_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		err []Error
	}
	tests := []struct {
		name string
		errs *Errs
		args args
		want *Errs
	}{
		{
			name: "nil errs",
			errs: nil,
			args: args{
				err: []Error{NewErr("test msg", "test descr")},
			},
			want: nil,
		},
		{
			name: "0 errs",
			errs: NewErrs(),
			args: args{
				err: []Error{NewErr("test msg", "test descr")},
			},
			want: &Errs{
				Errors: []Error{NewErr("test msg", "test descr")},
			},
		},
		{
			name: "non zero errs",
			errs: &Errs{
				Errors: []Error{
					NewErr("test msg 1", "test descr"),
				},
			},
			args: args{
				err: []Error{NewErr("test msg 2", "test descr")},
			},
			want: &Errs{
				Errors: []Error{
					NewErr("test msg 1", "test descr"),
					NewErr("test msg 2", "test descr"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.errs.Add(tt.args.err...)

			if !reflect.DeepEqual(tt.errs, tt.want) {
				t.Errorf("Add() want %v, got %v", tt.want, tt.errs)
			}
		})
	}
}

func TestErrs_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		errs *Errs
		want string
	}{
		{
			name: "nil errrs",
			errs: nil,
			want: "",
		},
		{
			name: "0 errrs",
			errs: NewErrs(),
			want: "",
		},
		{
			name: "non zero errs",
			errs: &Errs{
				Errors: []Error{
					NewErr("test msg 1", "test descr 1"),
					NewErr("test msg 2", "test descr 2"),
				},
			},
			want: "test msg 1:test descr 1; test msg 2:test descr 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.errs.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrs_GetErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		errs *Errs
		want []Error
	}{
		{
			name: "nil errs",
			errs: nil,
			want: nil,
		},
		{
			name: "0 errs",
			errs: NewErrs(),
			want: make([]Error, 0, 1),
		},
		{
			name: "not empty errs",
			errs: &Errs{
				Errors: []Error{
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
				},
			},
			want: []Error{
				NewErr("test msg", "test descr"),
				NewErr("test msg", "test descr"),
				NewErr("test msg", "test descr"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.errs.GetErrors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrs_Len(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		errs *Errs
		want int
	}{
		{
			name: "nil errs",
			errs: nil,
			want: 0,
		},
		{
			name: "0 errs",
			errs: NewErrs(),
			want: 0,
		},
		{
			name: "not empty errs",
			errs: &Errs{
				Errors: []Error{
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.errs.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrs_Sanitize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		errs *Errs
	}{
		{
			name: "nil errs",
			errs: nil,
		},
		{
			name: "not nil errs",
			errs: &Errs{
				Errors: []Error{
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
					NewErr("test msg", "test descr"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.errs.Sanitize()

			for _, err := range tt.errs.GetErrors() {
				got := err.Error()

				if got != "test msg:" {
					t.Errorf("Sanitize() want empty description, got error: %v", got)
				}
			}
		})
	}
}

func TestNewErrs(t *testing.T) { // nolint:dupl
	t.Parallel()

	tests := []struct {
		name string
		want *Errs
	}{
		{
			name: "new errs",
			want: &Errs{Errors: make([]Error, 0, 1)},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewErrs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewErrsWithLen(t *testing.T) {
	t.Parallel()

	type args struct {
		l int
		c int
	}
	tests := []struct {
		name string
		args args
		want *Errs
	}{
		{
			name: "0 len, 0 cap",
			args: args{
				l: 0,
				c: 0,
			},
			want: &Errs{Errors: make([]Error, 0)},
		},
		{
			name: "0 len, 10 cap",
			args: args{
				l: 0,
				c: 10,
			},
			want: &Errs{Errors: make([]Error, 0, 10)},
		},
		{
			name: "5 len, 10 cap",
			args: args{
				l: 5,
				c: 10,
			},
			want: &Errs{Errors: make([]Error, 5, 10)},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewErrsWithLen(tt.args.l, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrsWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
