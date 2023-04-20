package annotate

import "reflect"

// Defaulter is an interface that can be implemented by types that need to
// provide a default value. It is used by the DX[T] type.
type Defaulter[T any] interface {
	Default() *T
}

// DX is a wrapper type that can be used to wrap a pointer to a type that
// implements the Defaulter[T] interface. It will automatically call the
// Default() method on the wrapped value if it is nil. It will also apply the
// default values from struct tags if the wrapped value is a struct.
type DX[T Defaulter[T]] struct {
	Ref *T
}

// Value returns the wrapped value. If the wrapped value is nil, it will be
// initialized with a new value and the Default() method will be called on it.
// If the wrapped value is a struct, the default values from struct tags will
// be applied to it.
func (d *DX[T]) Value() *T {
	if d.Ref == nil {
		d.Ref = new(T)
		if t := removeIndirect(reflect.TypeOf(d.Ref)); t.Kind() == reflect.Struct {
			structTagDefaulter.Get(t).ApplyDefaults(d.Ref)
		}
		d.Ref = (*d.Ref).Default()
	}
	return d.Ref
}

// D is a wrapper type that can be used to wrap a pointer to a type. It will
// automatically initialize the wrapped value if it is nil. If the wrapped
// value is a struct, the default values from struct tags will be applied to it.
type D[T any] struct {
	Ref *T
}

// Value returns the wrapped value. If the wrapped value is nil, it will be
// initialized with a new value. If the wrapped value is a struct, the default
// values from struct tags will be applied to it.
func (d *D[T]) Value() *T {
	if d.Ref == nil {
		d.Ref = new(T)
		if t := removeIndirect(reflect.TypeOf(d.Ref)); t.Kind() == reflect.Struct {
			structTagDefaulter.Get(t).ApplyDefaults(d.Ref)
		}
	}
	return d.Ref
}
