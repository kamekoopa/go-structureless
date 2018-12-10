package gostructureless

import "fmt"

type Bool struct {
	Path string
	t    string
	x    bool
}

func newBool(path string, b bool) *Bool {
	return &Bool{
		Path: path,
		t:    "bool",
		x:    b,
	}
}
func (b *Bool) GetPath() string {
	return b.Path
}

func (b *Bool) Bool() (bool, error) {
	return b.x, nil
}
func (b *Bool) BoolOr(def bool) bool {
	return b.x
}
func (b *Bool) I64() (int64, error) {
	return 0, newTypeError(b.Path, b.t)
}
func (b *Bool) I64Or(def int64) int64 {
	return def
}
func (b *Bool) F64() (float64, error) {
	return 0.0, newTypeError(b.Path, b.t)
}
func (b *Bool) F64Or(def float64) float64 {
	return def
}
func (b *Bool) String() (string, error) {
	return fmt.Sprintf("%t", b.x), nil
}
func (b *Bool) StringOr(def string) string {
	s, _ := b.String()
	return s
}
func (b *Bool) At(i int) Value {
	return newTypeErrorNode(b.Path, b.t)
}
func (b *Bool) StringArray() ([]string, error) {
	return nil, newTypeError(b.Path, b.t)
}
func (b *Bool) StringArrayOr(def []string) []string {
	return def
}
func (b *Bool) Arr() ([]Value, error) {
	return nil, newTypeError(b.Path, b.t)
}
func (b *Bool) Key(key string) Value {
	return newTypeErrorNode(b.Path, b.t)
}
func (b *Bool) Map() (map[string]interface{}, error) {
	return nil, newTypeError(b.Path, b.t)
}
func (b *Bool) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (b *Bool) Obj() (map[string]Value, error) {
	return nil, newTypeError(b.Path, b.t)
}
func (b *Bool) IsNil() bool {
	return false
}
