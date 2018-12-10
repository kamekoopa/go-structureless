package gostructureless

import "testing"

func Test_MapではBoolは扱えない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.Bool()

	assertError(t, err)
}

func Test_MapではBoolOrはデフォルト値(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	v := x.BoolOr(true)

	assertTrue(t, v)
}

func Test_MapではInt64は扱えない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.I64()

	assertError(t, err)
}

func Test_MapではI64Orはデフォルト値(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	i := x.I64Or(0)

	assertInt64(t, 0, i)
}

func Test_MapではFloatは扱えない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.F64()

	assertError(t, err)
}

func Test_MapではF64Orデフォルト値(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	f := x.F64Or(0)

	assertFloat(t, 0, f)
}

func Test_MapではStringは扱えない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.String()

	assertError(t, err)
}

func Test_MapではStringOrのデフォルト値が利用される(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	s := x.StringOr("nothing")

	assertString(t, "nothing", s)
}

func Test_MapではAtで値は取り出せない(t *testing.T) {
	x := root(`{"key1": "value1", "key2": "value2"}`)

	v := x.At(0)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_MapではStringArrayで値は取り出せない(t *testing.T) {
	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.StringArray()

	assertError(t, err)
}

func Test_MapではStringArrayOrでデフォルト値(t *testing.T) {
	x := root(`{"key1": "value1", "key2": "value2"}`)

	arr := x.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(arr))
	assertString(t, "a", arr[0])
}

func Test_MapではArrで値は取り出せない(t *testing.T) {
	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, err := x.Arr()

	assertError(t, err)
}

func Test_MapではKeyで値を取り出せる(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	v, err := x.Key("key1").String()

	assertNotError(t, err)
	assertString(t, "value1", v)
}

func Test_Mapでは存在しないキーでは値を取り出せない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	_, ok := x.Key("key3").(*ErrorNode)

	assertTrue(t, ok)
}

func Test_MapではMapを扱える(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	v, err := x.Map()

	assertNotError(t, err)
	assertInt(t, 2, len(v))

	v1, ok1 := v["key1"]
	vs1, _ := v1.(string)
	assertTrue(t, ok1)
	assertString(t, "value1", vs1)

	v2, ok2 := v["key2"]
	vs2, _ := v2.(string)
	assertTrue(t, ok2)
	assertString(t, "value2", vs2)
}

func Test_MapではMapOrはデフォルト値は利用されない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	v := x.MapOr(map[string]interface{}{
		"aaa": "bbb",
	})

	_, ok := v["aaa"].(string)
	assertFalse(t, ok)
}

func Test_MapではObjで値を取り出せる(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	v, err := x.Obj()

	assertNotError(t, err)
	assertInt(t, 2, len(v))
	assertString(t, "value1", v["key1"].StringOr(""))
	assertString(t, "value2", v["key2"].StringOr(""))
}

func Test_MapはNilではない(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	assertFalse(t, x.IsNil())
}

func Test_MapをRootノードとして生成した時のPathが正しい(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	path := x.GetPath()

	assertString(t, "<root>", path)
}

func Test_Mapの要素のPathが正しい(t *testing.T) {

	x := root(`{"key1": "value1", "key2": "value2"}`)

	path := x.Key("key1").GetPath()

	assertString(t, "<root>.key1", path)
}
