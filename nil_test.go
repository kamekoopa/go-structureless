package gostructureless

import "testing"

func Test_NilではBoolは扱えない(t *testing.T) {

	x := root(`null`)

	_, err := x.Bool()

	assertError(t, err)
}

func Test_NilではBoolOrはデフォルト値(t *testing.T) {

	x := root(`null`)

	v := x.BoolOr(true)

	assertTrue(t, v)
}

func Test_NilではInt64は扱えない(t *testing.T) {

	x := root(`null`)

	_, err := x.I64()

	assertError(t, err)
}

func Test_NilではI64Orはデフォルト値(t *testing.T) {

	x := root(`null`)

	i := x.I64Or(0)

	assertInt64(t, 0, i)
}

func Test_NilではFloatは扱えない(t *testing.T) {

	x := root(`null`)

	_, err := x.F64()

	assertError(t, err)
}

func Test_NilではF64Orデフォルト値(t *testing.T) {

	x := root(`null`)

	f := x.F64Or(0)

	assertFloat(t, 0, f)
}

func Test_NilではStringは扱えない(t *testing.T) {

	x := root(`null`)

	_, err := x.String()

	assertError(t, err)
}

func Test_NilではStringOrのデフォルト値が利用される(t *testing.T) {

	x := root(`null`)

	s := x.StringOr("nothing")

	assertString(t, "nothing", s)
}

func Test_NilではAtで値は取り出せない(t *testing.T) {
	x := root(`null`)

	v := x.At(0)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_NilではStringArrayで値は取り出せない(t *testing.T) {
	x := root(`null`)

	_, err := x.StringArray()

	assertError(t, err)
}

func Test_NilではStringArrayOrでデフォルト値(t *testing.T) {
	x := root(`null`)

	arr := x.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(arr))
	assertString(t, "a", arr[0])
}

func Test_NilではArrで値は取り出せない(t *testing.T) {
	x := root(`null`)

	_, err := x.Arr()

	assertError(t, err)
}

func Test_NilではKeyでは値を取り出せない(t *testing.T) {

	x := root(`null`)

	v := x.Key("key")

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_NilではMapは扱えない(t *testing.T) {

	x := root(`null`)

	_, err := x.Map()

	assertError(t, err)
}

func Test_NilではMapOrはデフォルト値(t *testing.T) {

	x := root(`null`)

	v := x.MapOr(map[string]interface{}{
		"key1": "value1",
	})

	val, ok := v["key1"].(string)
	assertTrue(t, ok)
	assertString(t, "value1", val)
}

func Test_NilではObjでは値を取り出せない(t *testing.T) {

	x := root(`null`)

	_, err := x.Obj()

	assertError(t, err)
}

func Test_NilはNil(t *testing.T) {

	x := root(`null`)

	assertTrue(t, x.IsNil())
}

func Test_NilをRootノードとして生成した時のPathが正しい(t *testing.T) {

	x := root(`null`)

	path := x.GetPath()

	assertString(t, "<root>", path)
}
