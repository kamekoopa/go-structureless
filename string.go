package gostructureless

import (
	"strconv"
)

type String struct {
	Path string
	t    string
	x    string
}

func newString(path string, s string) *String {
	return &String{
		Path: path,
		t:    "string",
		x:    s,
	}
}
func (s *String) GetPath() string {
	return s.Path
}
func (s *String) Bool() (bool, error) {
	if b, e := strconv.ParseBool(s.x); e != nil {
		return false, newTypeError(s.Path, s.t)
	} else {
		return b, nil
	}
}
func (s *String) BoolOr(def bool) bool {
	if b, e := s.Bool(); e != nil {
		return def
	} else {
		return b
	}
}
func (s *String) I64() (int64, error) {
	if i, e := strconv.ParseInt(s.x, 10, 64); e != nil {
		return 0, newTypeError(s.Path, s.t)
	} else {
		return i, nil
	}
}
func (s *String) I64Or(def int64) int64 {
	if i, e := s.I64(); e != nil {
		return def
	} else {
		return i
	}
}
func (s *String) F64() (float64, error) {
	if f, err := strconv.ParseFloat(s.x, 64); err != nil {
		return 0.0, newTypeError(s.Path, s.t)
	} else {
		return f, nil
	}
}
func (s *String) F64Or(def float64) float64 {
	if f, e := s.F64(); e != nil {
		return def
	} else {
		return f
	}
}
func (s *String) String() (string, error) {
	return s.x, nil
}
func (s *String) StringOr(def string) string {
	x, _ := s.String()
	return x
}
func (s *String) At(i int) Value {
	return newTypeErrorNode(s.Path, s.t)
}
func (s *String) StringArray() ([]string, error) {
	return nil, newTypeError(s.Path, s.t)
}
func (s *String) StringArrayOr(def []string) []string {
	return def
}
func (s *String) Arr() ([]Value, error) {
	return nil, newTypeError(s.Path, s.t)
}
func (s *String) Key(key string) Value {
	return newTypeErrorNode(s.Path, s.t)
}
func (s *String) Map() (map[string]interface{}, error) {
	return nil, newTypeError(s.Path, s.t)
}
func (s *String) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (s *String) Obj() (map[string]Value, error) {
	return nil, newTypeError(s.Path, s.t)
}
func (s *String) IsNil() bool {
	return false
}
