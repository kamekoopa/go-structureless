package gostructureless

import "fmt"

type Array struct {
	Path string
	t    string
	x    []interface{}
}

func newArray(path string, a []interface{}) *Array {
	return &Array{
		Path: path,
		t:    "array",
		x:    a,
	}
}
func (a *Array) GetPath() string {
	return a.Path
}
func (a *Array) Bool() (bool, error) {
	return false, newTypeError(a.Path, a.t)
}
func (a *Array) BoolOr(def bool) bool {
	return def
}
func (a *Array) I64() (int64, error) {
	return 0, newTypeError(a.Path, a.t)
}
func (a *Array) I64Or(def int64) int64 {
	return def
}
func (a *Array) F64() (float64, error) {
	return 0.0, newTypeError(a.Path, a.t)
}
func (a *Array) F64Or(def float64) float64 {
	return def
}
func (a *Array) String() (string, error) {
	return "", newTypeError(a.Path, a.t)
}
func (a *Array) StringOr(def string) string {
	return def
}
func (a *Array) At(i int) Value {
	arr := a.x
	path := fmt.Sprintf("%s[%d]", a.Path, i)
	if i < len(arr) {
		return NewValue(path, arr[i])
	} else {
		return newNotFoundErrorNode(path)
	}
}
func (a *Array) StringArray() ([]string, error) {
	var result []string
	for i, elem := range a.x {
		if s, ok := elem.(string); ok == true {
			result = append(result, s)
		} else {
			return nil, a.elementTypeError(i, "string")
		}
	}
	return result, nil
}
func (a *Array) StringArrayOr(def []string) []string {
	if a, e := a.StringArray(); e != nil {
		return def
	} else {
		return a
	}
}
func (a *Array) Arr() ([]Value, error) {
	var arr []Value
	for i := 0; i < len(a.x); i++ {
		arr = append(arr, a.At(i))
	}
	return arr, nil
}
func (a *Array) Key(key string) Value {
	return newTypeErrorNode(a.Path, "array")
}
func (a *Array) Map() (map[string]interface{}, error) {
	return nil, newTypeError(a.Path, a.t)
}
func (a *Array) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (a *Array) Obj() (map[string]Value, error) {
	return nil, newTypeError(a.Path, a.t)
}
func (a *Array) IsNil() bool {
	return false
}
func (a *Array) elementTypeError(i int, t string) error {
	return newTypeError(fmt.Sprintf("%s[%d]", a.Path, i), t)
}
