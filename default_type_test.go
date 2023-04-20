package annotate

import (
	"reflect"
	"testing"
)

type concreteDXType struct {
	value string
}

func (c concreteDXType) Default() *concreteDXType {
	c.value = "default"
	return &c
}

func TestDX_Value(t *testing.T) {
	type testCase[T Defaulter[T]] struct {
		name       string
		d          DX[T]
		wantFields string
	}
	tests := []testCase[concreteDXType]{
		{"should default when ref is nil", DX[concreteDXType]{Ref: nil}, "default"},
		{"should use actual value", DX[concreteDXType]{Ref: &concreteDXType{"value"}}, "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Value(); !reflect.DeepEqual(got.value, tt.wantFields) {
				t.Errorf("Value() = %v, wantFields %v", got.value, tt.wantFields)
			}
		})
	}
}

type concreteDType struct {
	Value string `default:"default"`
}

func TestD_Value(t *testing.T) {
	type testCase[T any] struct {
		name       string
		d          D[T]
		wantFields string
	}
	tests := []testCase[concreteDType]{
		{"should default when ref is nil", D[concreteDType]{Ref: nil}, "default"},
		{"should use actual value", D[concreteDType]{Ref: &concreteDType{"value"}}, "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Value(); !reflect.DeepEqual(got.Value, tt.wantFields) {
				t.Errorf("Value() = %v, wantFields %v", got.Value, tt.wantFields)
			}
		})
	}
}
