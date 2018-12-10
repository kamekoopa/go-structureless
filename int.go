package gostructureless

import "fmt"

type Int struct {
	Path string
	t    string
	x    int64
}

func newInt(path string, b int64) *Int {
	return &Int{
		Path: path,
		t:    "int",
		x:    b,
	}
}
func (i *Int) GetPath() string {
	return i.Path
}
func (i *Int) Bool() (bool, error) {
	return false, newTypeError(i.Path, i.t)
}
func (i *Int) BoolOr(def bool) bool {
	if b, e := i.Bool(); e != nil {
		return def
	} else {
		return b
	}
}
func (i *Int) I64() (int64, error) {
	return i.x, nil
}
func (i *Int) I64Or(def int64) int64 {
	if i, e := i.I64(); e != nil {
		return def
	} else {
		return i
	}
}
func (i *Int) F64() (float64, error) {
	return float64(i.x), nil
}
func (i *Int) F64Or(def float64) float64 {
	if f, e := i.F64(); e != nil {
		return def
	} else {
		return f
	}
}
func (i *Int) String() (string, error) {
	return fmt.Sprintf("%d", i.x), nil
}
func (i *Int) StringOr(def string) string {
	if s, e := i.String(); e != nil {
		return def
	} else {
		return s
	}
}
func (in *Int) At(i int) Value {
	return in
}
func (i *Int) StringArray() ([]string, error) {
	return nil, newTypeError(i.Path, i.t)
}
func (i *Int) StringArrayOr(def []string) []string {
	if a, e := i.StringArray(); e != nil {
		return def
	} else {
		return a
	}
}
func (i *Int) Arr() ([]Value, error) {
	return nil, newTypeError(i.Path, i.t)
}
func (i *Int) Key(key string) Value {
	return i
}
func (i *Int) Map() (map[string]interface{}, error) {
	return nil, newTypeError(i.Path, i.t)
}
func (i *Int) MapOr(def map[string]interface{}) map[string]interface{} {
	if m, e := i.Map(); e != nil {
		return def
	} else {
		return m
	}
}
func (i *Int) Obj() (map[string]Value, error) {
	return nil, newTypeError(i.Path, i.t)
}
func (i *Int) IsNil() bool {
	return false
}
