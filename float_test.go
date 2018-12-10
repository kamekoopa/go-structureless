package gostructureless

import "testing"

func Test_Float64ではBoolは扱えない(t *testing.T) {

	x := root(`1`)

	_, err := x.Bool()

	assertError(t, err)
}

func Test_Float64ではBoolOrはデフォルト値(t *testing.T) {

	x := root(`1`)

	v := x.BoolOr(true)

	assertTrue(t, v)
}

func Test_Float64ではInt64を扱える(t *testing.T) {

	x := root(`1`)

	v, err := x.I64()

	assertNotError(t, err)
	assertInt64(t, 1, v)
}

func Test_Float64ではI64Orのデフォルト値は使われない(t *testing.T) {

	b := root(`1`)

	v := b.I64Or(0)

	assertInt64(t, 1, v)
}

func Test_Float64ではFloatとして扱える(t *testing.T) {

	x := root(`1`)

	v, err := x.F64()

	assertNotError(t, err)
	assertFloat64(t, 1.0, v)
}

func Test_Float64ではF64Orデフォルト値は使われない(t *testing.T) {

	x := root(`1`)

	v := x.F64Or(100.0)

	assertFloat(t, 1.0, v)
}

func Test_Float64では1は1(t *testing.T) {

	x := root(`1`)

	s, err := x.String()

	assertNotError(t, err)
	assertString(t, "1.000000", s)
}

func Test_Float64ではStringOrのデフォルト値は利用されない(t *testing.T) {

	x := root(`1`)

	s := x.StringOr("2")

	assertString(t, "1.000000", s)
}

func Test_Float64ではAtで値は取り出せない(t *testing.T) {
	x := root(`1`)

	v := x.At(0)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_Float64ではStringArrayで値は取り出せない(t *testing.T) {
	x := root(`1`)

	_, err := x.StringArray()

	assertError(t, err)
}

func Test_Float64ではStringArrayOrでデフォルト値(t *testing.T) {
	x := root(`1`)

	arr := x.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(arr))
	assertString(t, "a", arr[0])
}

func Test_Float64ではArrで値は取り出せない(t *testing.T) {
	x := root(`1`)

	_, err := x.Arr()

	assertError(t, err)
}

func Test_Float64ではKeyでは値を取り出せない(t *testing.T) {

	x := root(`1`)

	v := x.Key("key")

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_Float64ではMapは扱えない(t *testing.T) {

	x := root(`1`)

	_, err := x.Map()

	assertError(t, err)
}

func Test_Float64ではMapOrはデフォルト値(t *testing.T) {

	x := root(`1`)

	v := x.MapOr(map[string]interface{}{
		"key1": "value1",
	})

	val, ok := v["key1"].(string)
	assertTrue(t, ok)
	assertString(t, "value1", val)
}

func Test_Float64ではObjでは値を取り出せない(t *testing.T) {

	x := root(`1`)

	_, err := x.Obj()

	assertError(t, err)
}

func Test_Float64はNilではない(t *testing.T) {

	x := root(`1`)

	assertFalse(t, x.IsNil())
}

func Test_Float64をRootノードとして生成した時のPathが正しい(t *testing.T) {

	x := root(`1`)

	path := x.GetPath()

	assertString(t, "<root>", path)
}
