package annotate

import (
	"fmt"
	"reflect"
	"strconv"
)

type structTagDefaulterImpl struct {
	fields map[int]reflect.Value
}

func (s *structTagDefaulterImpl) ApplyDefaults(value any) {
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		panic(fmt.Errorf("expected pointer to struct, got %v", reflect.TypeOf(value)))
	}
	v := reflect.ValueOf(value).Elem()
	for i, defaultValue := range s.fields {
		if field := v.Field(i); field.CanSet() {
			field.Set(defaultValue)
		}
	}
}

func newStructTagDefaulter(t reflect.Type) *structTagDefaulterImpl {
	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("expected struct type, got %v", t))
	}

	fields := map[int]reflect.Value{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("default"); ok && tag != "-" {
			var defaultValue reflect.Value
			switch field.Type.Kind() {
			case reflect.String:
				defaultValue = reflect.ValueOf(tag)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if val, err := strconv.ParseInt(tag, 10, 64); err == nil {
					defaultValue = reflect.ValueOf(int(val))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if val, err := strconv.ParseUint(tag, 10, 64); err == nil {
					defaultValue = reflect.ValueOf(uint(val))
				}
			case reflect.Float32, reflect.Float64:
				if val, err := strconv.ParseFloat(tag, 64); err == nil {
					defaultValue = reflect.ValueOf(val)
				}
			case reflect.Bool:
				if val, err := strconv.ParseBool(tag); err == nil {
					defaultValue = reflect.ValueOf(val)
				}
			}

			fields[i] = defaultValue.Convert(field.Type)
		}
	}

	return &structTagDefaulterImpl{fields}
}

var (
	structTagDefaulter = NewTypeCache[*structTagDefaulterImpl](newStructTagDefaulter)
)

// ApplyDefaults applies default values from it's struct tags to the given value.
func ApplyDefaults(value any) {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Struct {
		structTagDefaulter.Get(t).ApplyDefaults(value)
	}
}
