package gostructureless

type Nil struct {
	Path string
	t    string
}

func newNil(path string) *Nil {
	return &Nil{
		Path: path,
		t:    "nil",
	}
}
func (n *Nil) GetPath() string {
	return n.Path
}
func (n *Nil) Bool() (bool, error) {
	return false, newTypeError(n.Path, n.t)
}
func (n *Nil) BoolOr(def bool) bool {
	return def
}
func (n *Nil) I64() (int64, error) {
	return 0, newTypeError(n.Path, n.t)
}
func (n *Nil) I64Or(def int64) int64 {
	return def
}
func (n *Nil) F64() (float64, error) {
	return 0.0, newTypeError(n.Path, n.t)
}
func (n *Nil) F64Or(def float64) float64 {
	return def
}
func (n *Nil) String() (string, error) {
	return "", newTypeError(n.Path, n.t)
}
func (n *Nil) StringOr(def string) string {
	return def
}
func (n *Nil) At(i int) Value {
	return newTypeErrorNode(n.Path, n.t)
}
func (n *Nil) StringArray() ([]string, error) {
	return nil, newTypeError(n.Path, n.t)
}
func (n *Nil) StringArrayOr(def []string) []string {
	return def
}
func (n *Nil) Arr() ([]Value, error) {
	return nil, newTypeError(n.Path, n.t)
}
func (n *Nil) Key(key string) Value {
	return newTypeErrorNode(n.Path, n.t)
}
func (n *Nil) Map() (map[string]interface{}, error) {
	return nil, newTypeError(n.Path, n.t)
}
func (n *Nil) MapOr(def map[string]interface{}) map[string]interface{} {
	return def
}
func (n *Nil) Obj() (map[string]Value, error) {
	return nil, newTypeError(n.Path, n.t)
}

func (n *Nil) IsNil() bool {
	return true
}
