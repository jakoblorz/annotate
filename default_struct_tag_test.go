package annotate

import (
	"reflect"
	"testing"
)

func compareReflectionMaps(map1 map[int]reflect.Value, map2 map[int]reflect.Value) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key := range map1 {
		if !reflect.DeepEqual(map1[key].Interface(), map2[key].Interface()) {
			return false
		}
	}
	return true
}

func Test_newStructTagDefaulter(t *testing.T) {
	type args struct {
		t reflect.Type
	}
	type testCase struct {
		name       string
		args       args
		wantFields map[int]reflect.Value
	}

	rootTests := []testCase{
		{"should return empty map when type is struct with no fields", args{t: reflect.TypeOf(struct{}{})}, map[int]reflect.Value{}},
		{"should return empty map when type is struct with one field with no default tag", args{t: reflect.TypeOf(struct{ value string }{})}, map[int]reflect.Value{}},
		{"should return empty map when type is struct with one field with no default tag but other tags", args{t: reflect.TypeOf(struct {
			value string `json:"value"`
		}{})}, map[int]reflect.Value{}},
		{"should return map with one field when type is struct with one field with default tag", args{t: reflect.TypeOf(struct {
			value string `default:"default"`
		}{})}, map[int]reflect.Value{0: reflect.ValueOf("default")}},
		{"should return map with one field when type is struct with one field with default tag and other tags", args{t: reflect.TypeOf(struct {
			value string `json:"value" default:"default"`
		}{})}, map[int]reflect.Value{0: reflect.ValueOf("default")}},
		{"should return map with one field when type is struct with one field with default tag and other tags", args{t: reflect.TypeOf(struct {
			value string `json:"value" default:"default"`
		}{})}, map[int]reflect.Value{0: reflect.ValueOf("default")}},
		{"should return map with two fields when type is struct with two fields with default tag and other tags", args{t: reflect.TypeOf(struct {
			value1 string `json:"value1" default:"default1"`
			value2 string `json:"value2" default:"default2"`
		}{})}, map[int]reflect.Value{0: reflect.ValueOf("default1"), 1: reflect.ValueOf("default2")}},
		{"should return map with one field when type is struct with two fields with one default tag and one other tags", args{t: reflect.TypeOf(struct {
			value1 string `json:"value1"`
			value2 string `json:"value2" default:"default2"`
		}{})}, map[int]reflect.Value{1: reflect.ValueOf("default2")}},
	}
	for _, tt := range rootTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStructTagDefaulter(tt.args.t); !compareReflectionMaps(got.fields, tt.wantFields) {
				t.Errorf("newStructTagDefaulter() = %v, want %v", got.fields, tt.wantFields)
			}
		})
	}
	typeTests := []testCase{
		{"should parse int types", args{t: reflect.TypeOf(struct {
			valueInt8  int8  `default:"1"`
			valueInt16 int16 `default:"2"`
			valueInt32 int32 `default:"3"`
			valueInt64 int64 `default:"4"`
			valueInt   int   `default:"5"`
		}{})}, map[int]reflect.Value{
			0: reflect.ValueOf(int8(1)),
			1: reflect.ValueOf(int16(2)),
			2: reflect.ValueOf(int32(3)),
			3: reflect.ValueOf(int64(4)),
			4: reflect.ValueOf(int(5)),
		}},
		{"should parse uint types", args{t: reflect.TypeOf(struct {
			valueUint8  uint8  `default:"1"`
			valueUint16 uint16 `default:"2"`
			valueUint32 uint32 `default:"3"`
			valueUint64 uint64 `default:"4"`
			valueUint   uint   `default:"5"`
		}{})}, map[int]reflect.Value{
			0: reflect.ValueOf(uint8(1)),
			1: reflect.ValueOf(uint16(2)),
			2: reflect.ValueOf(uint32(3)),
			3: reflect.ValueOf(uint64(4)),
			4: reflect.ValueOf(uint(5)),
		}},
		{"should parse float types", args{t: reflect.TypeOf(struct {
			valueFloat32 float32 `default:"1.1"`
			valueFloat64 float64 `default:"2.2"`
		}{})}, map[int]reflect.Value{
			0: reflect.ValueOf(float32(1.1)),
			1: reflect.ValueOf(float64(2.2)),
		}},
	}
	for _, tt := range typeTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStructTagDefaulter(tt.args.t); !compareReflectionMaps(got.fields, tt.wantFields) {
				t.Errorf("newStructTagDefaulter() = %v, want %v", got.fields, tt.wantFields)
			}
		})
	}
}

func Test_structTagDefaulterImpl_ApplyReplacements(t *testing.T) {

	type basicStruct struct {
		StringValue  string  `default:"default"`
		BoolValue    bool    `default:"true"`
		IntValue     int     `default:"1"`
		Int8Value    int8    `default:"2"`
		Int16Value   int16   `default:"3"`
		Int32Value   int32   `default:"4"`
		Int64Value   int64   `default:"5"`
		UintValue    uint    `default:"6"`
		Uint8Value   uint8   `default:"7"`
		Uint16Value  uint16  `default:"8"`
		Uint32Value  uint32  `default:"9"`
		Uint64Value  uint64  `default:"10"`
		Float32Value float32 `default:"11.1"`
		Float64Value float64 `default:"12.2"`
	}
	defaultValues := basicStruct{
		StringValue:  "default",
		BoolValue:    true,
		IntValue:     1,
		Int8Value:    2,
		Int16Value:   3,
		Int32Value:   4,
		Int64Value:   5,
		UintValue:    6,
		Uint8Value:   7,
		Uint16Value:  8,
		Uint32Value:  9,
		Uint64Value:  10,
		Float32Value: 11.1,
		Float64Value: 12.2,
	}

	t.Run("should set default values for all fields", func(t *testing.T) {
		s := structTagDefaulter.Get(removeIndirect(reflect.TypeOf(basicStruct{})))

		got := basicStruct{}
		s.ApplyReplacements(&got)
		if !reflect.DeepEqual(got, defaultValues) {
			t.Errorf("ApplyReplacements() = %v, want %v", got, defaultValues)
		}
	})
}
