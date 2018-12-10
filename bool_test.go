package gostructureless

import "testing"

func Test_trueから生成できる(t *testing.T) {

	b, e := root(`true`).Bool()

	assertNotError(t, e)
	assertTrue(t, b)
}

func Test_falseから生成できる(t *testing.T) {

	b, e := root(`false`).Bool()

	assertNotError(t, e)
	assertFalse(t, b)
}

func Test_BoolではBoolを扱える(t *testing.T) {

	b := root(`true`)

	v, err := b.Bool()

	assertNotError(t, err)
	assertTrue(t, v)
}

func Test_BoolではBoolOrのデフォルト値は使われない(t *testing.T) {

	b := root(`true`)

	v := b.BoolOr(false)

	assertTrue(t, v)
}

func Test_BoolではIntは扱えない(t *testing.T) {

	b := root(`true`)

	_, err := b.I64()

	assertError(t, err)
}

func Test_BoolではI64Orデフォルト値(t *testing.T) {

	b := root(`true`)

	i := b.I64Or(0)

	assertInt64(t, 0, i)
}

func Test_BoolではFloatは扱えない(t *testing.T) {

	b := root(`true`)

	_, err := b.F64()

	assertError(t, err)
}

func Test_BoolではF64Orデフォルト値(t *testing.T) {

	b := root(`true`)

	f := b.F64Or(0)

	assertFloat(t, 0, f)
}

func Test_Boolではtrueはtrue(t *testing.T) {

	b := root(`true`)

	s, err := b.String()

	assertNotError(t, err)
	assertString(t, "true", s)
}

func Test_Boolではfalseはfalse(t *testing.T) {

	b := root(`false`)

	s, err := b.String()

	assertNotError(t, err)
	assertString(t, "false", s)
}

func Test_BoolではStringOrのデフォルト値は利用されない(t *testing.T) {

	b := root(`true`)

	s := b.StringOr("false")

	assertString(t, "true", s)
}

func Test_BoolではAtで値は取り出せない(t *testing.T) {
	b := root(`true`)

	v := b.At(0)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_BoolではStringArrayで値は取り出せない(t *testing.T) {
	b := root(`true`)

	_, err := b.StringArray()

	assertError(t, err)
}

func Test_BoolではStringArrayOrでデフォルト値(t *testing.T) {
	b := root(`true`)

	arr := b.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(arr))
	assertString(t, "a", arr[0])
}

func Test_BoolではArrで値は取り出せない(t *testing.T) {
	b := root(`true`)

	_, err := b.Arr()

	assertError(t, err)
}

func Test_BoolではKeyでは値を取り出せない(t *testing.T) {

	b := root(`true`)

	v := b.Key("key")

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_BoolではMapは扱えない(t *testing.T) {

	b := root(`true`)

	_, err := b.Map()

	assertError(t, err)
}

func Test_BoolではMapOrはデフォルト値(t *testing.T) {

	b := root(`true`)

	v := b.MapOr(map[string]interface{}{
		"key1": "value1",
	})

	val, ok := v["key1"].(string)
	assertTrue(t, ok)
	assertString(t, "value1", val)
}

func Test_BoolではObjでは値を取り出せない(t *testing.T) {

	b := root(`true`)

	_, err := b.Obj()

	assertError(t, err)
}

func Test_BoolはNilではない(t *testing.T) {

	b := root(`true`)

	assertFalse(t, b.IsNil())
}

func Test_BoolをRootノードとして生成した時のPathが正しい(t *testing.T) {

	b := root(`true`)

	path := b.GetPath()

	assertString(t, "<root>", path)
}
