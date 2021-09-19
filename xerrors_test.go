// nolint:dupl,funlen
package xerrors

import (
	"reflect"
	"testing"
)

func TestNewXErrs(t *testing.T) {
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

func TestXErrs_Error(t *testing.T) {
	t.Parallel()

	type args struct {
		xerrs *XErrs
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nil XErrs",
			args: args{
				xerrs: nil,
			},
			want: "",
		},
		{
			name: "no errors",
			args: args{
				xerrs: &XErrs{},
			},
			want: "",
		},
		{
			name: "1 error",
			args: args{
				xerrs: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
			},
			want: "test msg: test descr; map[]",
		},
		{
			name: "multiple errors",
			args: args{
				xerrs: &XErrs{Errs: []XError{
					NewXErr("test msg", "test descr", nil, nil),
					NewXErr("test msg 2", "test descr 2", map[string]interface{}{"filed": "user_id"}, nil),
					NewXErr("test msg 3", "test descr 3",
						map[string]interface{}{"filed": "user_id"}, map[string]interface{}{"error_info": "some error"}),
				}},
			},
			// nolint:lll
			want: "test msg: test descr; map[];test msg 2: test descr 2; map[filed:user_id];test msg 3: test descr 3; map[filed:user_id]",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.args.xerrs
			if got := err.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErrs_GetErrors(t *testing.T) {
	t.Parallel()

	type args struct {
		xerrs *XErrs
	}

	tests := []struct {
		name string
		args args
		want []XError
	}{
		{
			name: "nil XErrs",
			args: args{
				xerrs: nil,
			},
			want: nil,
		},
		{
			name: "no errors",
			args: args{
				xerrs: &XErrs{},
			},
			want: nil,
		},
		{
			name: "1 error",
			args: args{
				xerrs: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
			},
			want: []XError{NewXErr("test msg", "test descr", nil, nil)},
		},
		{
			name: "multiple errors",
			args: args{
				xerrs: &XErrs{Errs: []XError{
					NewXErr("test msg", "test descr", nil, nil),
					NewXErr("test msg 2", "test descr 2", map[string]interface{}{"filed": "user_id"}, nil),
					NewXErr("test msg 3", "test descr 3",
						map[string]interface{}{"filed": "user_id"}, map[string]interface{}{"error_info": "some error"}),
				}},
			},
			want: []XError{
				NewXErr("test msg", "test descr", nil, nil),
				NewXErr("test msg 2", "test descr 2", map[string]interface{}{"filed": "user_id"}, nil),
				NewXErr("test msg 3", "test descr 3",
					map[string]interface{}{"filed": "user_id"}, map[string]interface{}{"error_info": "some error"}),
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.args.xerrs.GetErrors()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErrs_Len(t *testing.T) {
	t.Parallel()

	type args struct {
		xerrs *XErrs
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nil XErrs",
			args: args{
				xerrs: nil,
			},
			want: 0,
		},
		{
			name: "no errors",
			args: args{
				xerrs: &XErrs{},
			},
			want: 0,
		},
		{
			name: "1 error",
			args: args{
				xerrs: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
			},
			want: 1,
		},
		{
			name: "multiple errors",
			args: args{
				xerrs: &XErrs{Errs: []XError{
					NewXErr("test msg", "test descr", nil, nil),
					NewXErr("test msg 2", "test descr 2", map[string]interface{}{"filed": "user_id"}, nil),
					NewXErr("test msg 3", "test descr 3",
						map[string]interface{}{"filed": "user_id"}, map[string]interface{}{"error_info": "some error"}),
				}},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.args.xerrs.Len()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXErrs_Sanitize(t *testing.T) {
	t.Parallel()

	type args struct {
		xerrs *XErrs
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil XErrs",
			args: args{
				xerrs: nil,
			},
		},
		{
			name: "no errors",
			args: args{
				xerrs: &XErrs{},
			},
		},
		{
			name: "1 error",
			args: args{
				xerrs: &XErrs{Errs: []XError{NewXErr("test msg", "test descr", nil, nil)}},
			},
		},
		{
			name: "multiple errors",
			args: args{
				xerrs: &XErrs{Errs: []XError{
					NewXErr("test msg", "test descr", nil, nil),
					NewXErr("test msg 2", "test descr 2", map[string]interface{}{"filed": "user_id"}, nil),
					NewXErr("test msg 3", "test descr 3",
						map[string]interface{}{"filed": "user_id"}, map[string]interface{}{"error_info": "some error"}),
				}},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.args.xerrs.Sanitize()

			for i, err := range tt.args.xerrs.GetErrors() {
				if err.GetDescription() != "" {
					t.Errorf("Sanitize(): xerrs[%d].Description = %s, want = \"\"", i, err.GetDescription())
				}
			}
		})
	}
}
