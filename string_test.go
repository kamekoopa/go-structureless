package gostructureless

import "testing"

func Test_StringではBoolはtrueはtrue(t *testing.T) {

	x := root(`"true"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertTrue(t, v)
}

func Test_StringではBoolはTrueはtrue(t *testing.T) {

	x := root(`"True"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertTrue(t, v)
}

func Test_StringではBoolはTRUEはtrue(t *testing.T) {

	x := root(`"TRUE"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertTrue(t, v)
}

func Test_StringではBoolは1はtrue(t *testing.T) {

	x := root(`"1"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertTrue(t, v)
}

func Test_StringではBoolはfalseはfalse(t *testing.T) {

	x := root(`"false"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertFalse(t, v)
}

func Test_StringではBoolはFalseはfalse(t *testing.T) {

	x := root(`"False"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertFalse(t, v)
}

func Test_StringではBoolはFALSEはfalse(t *testing.T) {

	x := root(`"FALSE"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertFalse(t, v)
}

func Test_StringではBoolは0はfalse(t *testing.T) {

	x := root(`"0"`)

	v, err := x.Bool()

	assertNotError(t, err)
	assertFalse(t, v)
}

func Test_StringではBoolOrはパースできればそれ(t *testing.T) {

	x := root(`"true"`)

	v := x.BoolOr(false)

	assertTrue(t, v)
}

func Test_StringではBoolOrはパースできなければデフォルト値(t *testing.T) {

	x := root(`"nya-n"`)

	v := x.BoolOr(false)

	assertFalse(t, v)
}

func Test_Stringでは数値文字列ならInt64を扱える(t *testing.T) {

	x := root(`"1"`)

	v, err := x.I64()

	assertNotError(t, err)
	assertInt64(t, 1, v)
}

func Test_Stringでは数値文字列でないならInt64を扱えない(t *testing.T) {

	x := root(`"a"`)

	_, err := x.I64()

	assertError(t, err)
}

func Test_Stringでは数値文字列ならI64Orのデフォルト値は使われない(t *testing.T) {

	b := root(`"1"`)

	v := b.I64Or(0)

	assertInt64(t, 1, v)
}

func Test_Stringでは数値文字列でないならI64Orのデフォルト値が使われる(t *testing.T) {

	b := root(`"a"`)

	v := b.I64Or(0)

	assertInt64(t, 0, v)
}

func Test_Stringでは数値文字列ならFloat64を扱える(t *testing.T) {

	x := root(`"1.0"`)

	v, err := x.F64()

	assertNotError(t, err)
	assertFloat64(t, 1.0, v)
}

func Test_Stringでは数値文字列でないならFloat64を扱えない(t *testing.T) {

	x := root(`"a"`)

	_, err := x.F64()

	assertError(t, err)
}

func Test_Stringでは数値文字列ならF64Orのデフォルト値は使われない(t *testing.T) {

	b := root(`"1.0"`)

	v := b.F64Or(0)

	assertFloat64(t, 1.0, v)
}

func Test_Stringでは数値文字列でないならF64Orのデフォルト値が使われる(t *testing.T) {

	b := root(`"a"`)

	v := b.F64Or(0)

	assertFloat64(t, 0, v)
}

func Test_Stringではaaaはaaa(t *testing.T) {

	x := root(`"aaa"`)

	s, err := x.String()

	assertNotError(t, err)
	assertString(t, "aaa", s)
}

func Test_StringではStringOrのデフォルト値は利用されない(t *testing.T) {

	x := root(`"aaa"`)

	s := x.StringOr("bbb")

	assertString(t, "aaa", s)
}

func Test_StringではAtで値は取り出せない(t *testing.T) {
	x := root(`"aaa"`)

	v := x.At(0)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_StringではStringArrayで値は取り出せない(t *testing.T) {
	x := root(`"aaa"`)

	_, err := x.StringArray()

	assertError(t, err)
}

func Test_StringではStringArrayOrでデフォルト値(t *testing.T) {
	x := root(`"aaa"`)

	arr := x.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(arr))
	assertString(t, "a", arr[0])
}

func Test_StringではArrで値は取り出せない(t *testing.T) {
	x := root(`"aaa"`)

	_, err := x.Arr()

	assertError(t, err)
}

func Test_StringではKeyでは値を取り出せない(t *testing.T) {

	x := root(`"aaa"`)

	v := x.Key("key")

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_StringではMapは扱えない(t *testing.T) {

	x := root(`"aaa"`)

	_, err := x.Map()

	assertError(t, err)
}

func Test_StringではMapOrはデフォルト値(t *testing.T) {

	x := root(`"aaa"`)

	v := x.MapOr(map[string]interface{}{
		"key1": "value1",
	})

	val, ok := v["key1"].(string)
	assertTrue(t, ok)
	assertString(t, "value1", val)
}

func Test_StringではObjでは値を取り出せない(t *testing.T) {

	x := root(`"aaa"`)

	_, err := x.Obj()

	assertError(t, err)
}

func Test_StringはNilではない(t *testing.T) {

	x := root(`"aaa"`)

	assertFalse(t, x.IsNil())
}

func Test_StringをRootノードとして生成した時のPathが正しい(t *testing.T) {

	x := root(`"aaa"`)

	path := x.GetPath()

	assertString(t, "<root>", path)
}
