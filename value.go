package gostructureless

import (
	"fmt"
)

type Value interface {
	GetPath() string

	Bool() (bool, error)
	BoolOr(def bool) bool

	I64() (int64, error)
	I64Or(def int64) int64

	F64() (float64, error)
	F64Or(def float64) float64

	String() (string, error)
	StringOr(def string) string

	At(i int) Value
	StringArray() ([]string, error)
	StringArrayOr(def []string) []string
	Arr() ([]Value, error)

	Key(key string) Value
	Map() (map[string]interface{}, error)
	MapOr(def map[string]interface{}) map[string]interface{}
	Obj() (map[string]Value, error)

	IsNil() bool
}

func Root(v interface{}) Value {
	return NewValue("<root>", v)
}

func NewValue(path string, v interface{}) Value {
	if v == nil {
		return newNil(path)
	} else {
		switch x := v.(type) {
		case bool:
			return newBool(path, x)
		case int8:
			return newInt(path, int64(x))
		case int16:
			return newInt(path, int64(x))
		case int32:
			return newInt(path, int64(x))
		case int64:
			return newInt(path, int64(x))
		case float32:
			return newFloat(path, float64(x))
		case float64:
			return newFloat(path, float64(x))
		case string:
			return newString(path, x)
		case []interface{}:
			return newArray(path, x)
		case map[string]interface{}:
			return newMapS(path, x)
		case map[interface{}]interface{}:
			return newMapI(path, x)
		default:
			if d1, ok := v.(map[interface{}]interface{}); ok == true {
				return newMapI(path, d1)
			} else if d2, ok := v.(map[string]interface{}); ok == true {
				return newMapS(path, d2)
			}
		}

		return newTypeErrorNode(path, fmt.Sprintf("%T", v))
	}
}
