package annotate

type defaulter[T any] interface {
	Default() *T
}

type Default[T defaulter[T]] struct {
	Ref *T
}

func (d *Default[T]) Value() *T {
	if d.Ref == nil {
		d.Ref = (*new(T)).Default()
	}
	return d.Ref
}
