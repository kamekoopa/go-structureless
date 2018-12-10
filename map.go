package gostructureless

import "fmt"

type Map struct {
	Path string
	t    string
	x    map[string]interface{}
}

func newMapI(path string, m map[interface{}]interface{}) *Map {

	var x = make(map[string]interface{})
	for k, v := range m {
		x[k.(string)] = v
	}

	return newMapS(path, x)
}
func newMapS(path string, m map[string]interface{}) *Map {

	return &Map{
		Path: path,
		t:    "map",
		x:    m,
	}
}
func (m *Map) GetPath() string {
	return m.Path
}
func (m *Map) Bool() (bool, error) {
	return false, newTypeError(m.Path, m.t)
}
func (m *Map) BoolOr(def bool) bool {
	return def
}
func (m *Map) I64() (int64, error) {
	return 0, newTypeError(m.Path, m.t)
}
func (m *Map) I64Or(def int64) int64 {
	return def
}
func (m *Map) F64() (float64, error) {
	return 0.0, newTypeError(m.Path, m.t)
}
func (m *Map) F64Or(def float64) float64 {
	return def
}
func (m *Map) String() (string, error) {
	return "", newTypeError(m.Path, m.t)
}
func (m *Map) StringOr(def string) string {
	return def
}
func (m *Map) At(i int) Value {
	return newTypeErrorNode(m.Path, m.t)
}
func (m *Map) StringArray() ([]string, error) {
	return nil, newTypeError(m.Path, m.t)
}
func (m *Map) StringArrayOr(def []string) []string {
	return def
}
func (m *Map) Arr() ([]Value, error) {
	return nil, newTypeError(m.Path, m.t)
}
func (m *Map) Key(key string) Value {
	path := fmt.Sprintf("%s.%s", m.Path, key)
	if v, ok := m.x[key]; ok == true {
		return NewValue(path, v)
	} else {
		return newNotFoundErrorNode(path)
	}
}
func (m *Map) Map() (map[string]interface{}, error) {
	return m.x, nil
}
func (m *Map) MapOr(def map[string]interface{}) map[string]interface{} {
	v, _ := m.Map()
	return v
}
func (m *Map) Obj() (map[string]Value, error) {

	var ret = make(map[string]Value)
	for k, v := range m.x {
		ret[k] = NewValue(fmt.Sprintf("%s.%s", m.Path, k), v)
	}

	return ret, nil
}

func (m *Map) IsNil() bool {
	return false
}
