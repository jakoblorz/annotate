package annotate

import (
	"reflect"
	"sync"
)

func removeIndirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

type TypeCache[T any] interface{}

type typeCache[T any] struct {
	sync.RWMutex
	seen map[reflect.Type]T

	computeValue func(reflect.Type) T
}

func (cache *typeCache[T]) Get(t reflect.Type) T {
	t = removeIndirect(t)

	cache.RLock()
	computedValue, ok := cache.seen[t]
	cache.RUnlock()

	if ok {
		return computedValue
	}

	computedValue = cache.computeValue(t)
	cache.Lock()
	cache.seen[t] = computedValue
	cache.Unlock()

	return computedValue
}

func NewTypeCache[T any](computeValue func(reflect.Type) T) TypeCache[T] {
	return &typeCache[T]{
		seen:         make(map[reflect.Type]T),
		computeValue: computeValue,
	}
}
