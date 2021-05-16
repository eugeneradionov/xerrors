package xerrors

import (
	"reflect"
	"testing"
)

func TestNewXErrs(t *testing.T) { // nolint:dupl
	t.Parallel()

	tests := []struct {
		name string
		want *XErrs
	}{
		{
			name: "new XErrs",
			want: &XErrs{Errs: make([]XError, 0, 1)},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewXErrs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXErrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewXErrsWithLen(t *testing.T) {
	t.Parallel()

	type args struct {
		l int
		c int
	}
	tests := []struct {
		name string
		args args
		want *XErrs
	}{
		{
			name: "0 len 0 cap",
			args: args{
				l: 0,
				c: 0,
			},
			want: &XErrs{
				Errs: make([]XError, 0),
			},
		},
		{
			name: "0 len 10 cap",
			args: args{
				l: 0,
				c: 10,
			},
			want: &XErrs{
				Errs: make([]XError, 0, 10),
			},
		},
		{
			name: "5 len 10 cap",
			args: args{
				l: 5,
				c: 10,
			},
			want: &XErrs{
				Errs: make([]XError, 5, 10),
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NewXErrsWithLen(tt.args.l, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewXErrsWithLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErrs_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		xerrs []XError
	}
	tests := []struct {
		name  string
		xerrs *XErrs
		args  args
		want  *XErrs
	}{
		{
			name:  "nil XErr",
			xerrs: nil,
			args: args{
				xerrs: []XError{NewXErr("test msg", "test descr", nil, nil)},
			},
			want: nil,
		},
		{
			name:  "empty XErr",
			xerrs: NewXErrs(),
			args: args{
				xerrs: []XError{NewXErr("test msg", "test descr", nil, nil)},
			},
			want: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
		},
		{
			name:  "not empty XErr",
			xerrs: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
			args: args{
				xerrs: []XError{NewXErr("test msg 2", "test descr 2", nil, nil)},
			},
			want: &XErrs{
				Errs: []XError{
					NewXErr("test msg", "test descr", nil, nil),
					NewXErr("test msg 2", "test descr 2", nil, nil),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.xerrs.Add(tt.args.xerrs...)

			if !reflect.DeepEqual(tt.want, tt.xerrs) {
				t.Errorf("Add() want: %v, got: %v", tt.want, tt.xerrs)
			}
		})
	}
}
