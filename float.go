package gostructureless

import "fmt"

type Float struct {
	Path string
	t    string
	x    float64
}

func newFloat(path string, f float64) *Float {
	return &Float{
		Path: path,
		t:    "float",
		x:    f,
	}
}
func (f *Float) GetPath() string {
	return f.Path
}
func (f *Float) Bool() (bool, error) {
	return false, newTypeError(f.Path, f.t)
}
func (f *Float) BoolOr(def bool) bool {
	return def
}
func (f *Float) I64() (int64, error) {
	return int64(f.x), nil
}
func (f *Float) I64Or(def int64) int64 {
	x, _ := f.I64()
	return x
}
func (f *Float) F64() (float64, error) {
	return f.x, nil
}
func (f *Float) F64Or(def float64) float64 {
	x, _ := f.F64()
	return x
}
func (f *Float) String() (string, error) {
	return fmt.Sprintf("%f", f.x), nil
}
func (f *Float) StringOr(def string) string {
	s, _ := f.String()
	return s
}
func (f *Float) At(i int) Value {
	return newTypeErrorNode(f.Path, f.t)
}
func (f *Float) StringArray() ([]string, error) {
	return nil, newTypeError(f.Path, f.t)
}
func (f *Float) StringArrayOr(def []string) []string {
	return def
}
func (f *Float) Arr() ([]Value, error) {
	return nil, newTypeError(f.Path, f.t)
}
func (f *Float) Key(key string) Value {
	return newTypeErrorNode(f.Path, f.t)
}
func (f *Float) Map() (map[string]interface{}, error) {
	return nil, newTypeError(f.Path, f.t)
}
func (f *Float) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (f *Float) Obj() (map[string]Value, error) {
	return nil, newTypeError(f.Path, f.t)
}
func (f *Float) IsNil() bool {
	return false
}
