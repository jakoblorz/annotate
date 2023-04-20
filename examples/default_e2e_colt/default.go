package main

import (
	"encoding/json"
	"github.com/jakoblorz/annotate"
	"go.mongodb.org/mongo-driver/bson"
)

type Default[T any] struct{ annotate.D[T] }

func (d *Default[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Value())
}

func (d *Default[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, d.Value())
}

func (d *Default[T]) MarshalBSON() ([]byte, error) {
	return bson.Marshal(d.Value())
}

func (d *Default[T]) UnmarshalBSON(data []byte) error {
	return bson.Unmarshal(data, d.Value())
}

type DefaultX[T annotate.Defaulter[T]] struct{ annotate.DX[T] }

func (d *DefaultX[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Value())
}

func (d *DefaultX[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, d.Value())
}

func (d *DefaultX[T]) MarshalBSON() ([]byte, error) {
	return bson.Marshal(d.Value())
}

func (d *DefaultX[T]) UnmarshalBSON(data []byte) error {
	return bson.Unmarshal(data, d.Value())
}
