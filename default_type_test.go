package annotate

import (
	"reflect"
	"testing"
)

type concreteType struct {
	value string
}

func (c concreteType) Default() *concreteType {
	c.value = "default"
	return &c
}

func TestDefault_Value(t *testing.T) {
	type testCase[T defaulter[T]] struct {
		name string
		d    Default[T]
		want string
	}
	tests := []testCase[concreteType]{
		{"should default when ref is nil", Default[concreteType]{Ref: nil}, "default"},
		{"should use actual value", Default[concreteType]{Ref: &concreteType{"value"}}, "value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Value(); !reflect.DeepEqual(got.value, tt.want) {
				t.Errorf("Value() = %v, want %v", got.value, tt.want)
			}
		})
	}
}
