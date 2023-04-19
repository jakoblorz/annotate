package annotate

import "reflect"

type fieldValueDefaulter struct{}

func createFieldValueDefaulterForStructTag(t reflect.Type) *fieldValueDefaulter {
	return nil
}

var (
	fieldValueDefaulterCache = NewTypeCache[*fieldValueDefaulter](createFieldValueDefaulterForStructTag)
)
