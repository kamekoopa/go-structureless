package gostructureless

import "fmt"

func newTypeError(path, _type string) error {
	return fmt.Errorf("path: %s is not type `%s`", path, _type)
}

func newNotFoundError(path string) error {
	return fmt.Errorf("path: %s does not found", path)
}

type ErrorNode struct {
	Path string
	err  error
}

func newTypeErrorNode(path, _type string) *ErrorNode {
	return &ErrorNode{
		Path: path,
		err:  newTypeError(path, _type),
	}
}
func newNotFoundErrorNode(path string) *ErrorNode {
	return &ErrorNode{
		Path: path,
		err:  newNotFoundError(path),
	}
}
func (e *ErrorNode) GetPath() string {
	return e.Path
}
func (e *ErrorNode) Bool() (bool, error) {
	return false, e.err
}
func (e *ErrorNode) BoolOr(def bool) bool {
	return def
}
func (e *ErrorNode) I64() (int64, error) {
	return 0, e.err
}
func (e *ErrorNode) I64Or(def int64) int64 {
	return def
}
func (e *ErrorNode) F64() (float64, error) {
	return 0.0, e.err
}
func (e *ErrorNode) F64Or(def float64) float64 {
	return def
}
func (e *ErrorNode) String() (string, error) {
	return "", e.err
}
func (e *ErrorNode) StringOr(def string) string {
	return def
}
func (e *ErrorNode) At(i int) Value {
	return e
}
func (e *ErrorNode) StringArray() ([]string, error) {
	return nil, e.err
}
func (e *ErrorNode) StringArrayOr(def []string) []string {
	return def
}
func (e *ErrorNode) Arr() ([]Value, error) {
	return nil, e.err
}
func (e *ErrorNode) Key(key string) Value {
	return e
}
func (e *ErrorNode) Map() (map[string]interface{}, error) {
	return nil, e.err
}
func (e *ErrorNode) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (e *ErrorNode) Obj() (map[string]Value, error) {
	return nil, e.err
}
func (e *ErrorNode) IsNil() bool {
	return false
}
