package annotate

import "reflect"

type defaulter[T any] interface {
	Default() *T
}

type Default[T defaulter[T]] struct {
	Ref *T
}

func (d *Default[T]) Value() *T {
	if d.Ref == nil {
		d.Ref = new(T)
		if t := reflect.TypeOf(d.Ref); t.Kind() == reflect.Struct {
			structTagDefaulter.Get(t).ApplyReplacements(d.Ref)
		}
		d.Ref = (*d.Ref).Default()
	}
	return d.Ref
}
