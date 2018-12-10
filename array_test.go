package gostructureless

import (
	"testing"
)

func Test_文字列配列から各要素をValue型として取り出せる(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	actual0 := arr.At(0).StringOr("")
	assertString(t, "a", actual0)

	actual1 := arr.At(1).StringOr("")
	assertString(t, "b", actual1)

	actual2 := arr.At(2).StringOr("")
	assertString(t, "c", actual2)
}

func Test_文字列配列から生成したValueから文字列配列をとして直接取り出せる(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	actuals, err := arr.StringArray()
	if err != nil {
		t.Error(err.Error())
	}

	assertString(t, "a", actuals[0])
	assertString(t, "b", actuals[1])
	assertString(t, "c", actuals[2])
}

func Test_文字列配列から生成したValueからStringArrayOrで取り出すとそれが返る(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	actuals := arr.StringArrayOr([]string{})

	assertInt(t, 3, len(actuals))
	assertString(t, "a", actuals[0])
	assertString(t, "b", actuals[1])
	assertString(t, "c", actuals[2])
}

func Test_文字列じゃない配列から生成したValueから文字列配列を取り出そうとするとerror(t *testing.T) {

	arr := root(`[{}, {}, {}]`)

	_, err := arr.StringArray()
	assertError(t, err)
}

func Test_文字列じゃない配列から生成したValueから文字列配列からStringArrayOrで取り出すとデフォルト値(t *testing.T) {

	arr := root(`[{}, {}, {}]`)

	actuals := arr.StringArrayOr([]string{"a"})

	assertInt(t, 1, len(actuals))
	assertString(t, "a", actuals[0])
}

func Test_オブジェクト配列から生成したValueからValueの配列を生成できる(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	actuals, err := arr.Arr()

	assertNotError(t, err)
	assertInt(t, 3, len(actuals))
	assertString(t, "a", actuals[0].Key("key").StringOr(""))
	assertString(t, "b", actuals[1].Key("key").StringOr(""))
	assertString(t, "c", actuals[2].Key("key").StringOr(""))
}

func Test_配列のサイズ外にアクセスするとエラー(t *testing.T) {

	arr := root(`[{"key": "a"}]`)

	v := arr.At(1)

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_ArrayではBoolは扱えない(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	_, err := arr.Bool()

	assertError(t, err)
}

func Test_ArrayではBoolOrデフォルト値(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	b := arr.BoolOr(true)

	assertTrue(t, b)
}

func Test_ArrayではIntは扱えない(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	_, err := arr.I64()

	assertError(t, err)
}

func Test_ArrayではI64Orデフォルト値(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	i := arr.I64Or(0)

	assertInt64(t, 0, i)
}

func Test_ArrayではFloatは扱えない(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	_, err := arr.F64()

	assertError(t, err)
}

func Test_ArrayではF64Orデフォルト値(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	f := arr.F64Or(0)

	assertFloat(t, 0, f)
}

func Test_ArrayではStringは扱えない(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	_, err := arr.String()

	assertError(t, err)
}

func Test_ArrayではStringOrデフォルト値(t *testing.T) {

	arr := root(`[{"key": "a"}, {"key": "b"}, {"key": "c"}]`)

	s := arr.StringOr("def")

	assertString(t, "def", s)
}

func Test_ArrayではKeyでは値を取り出せない(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	v := arr.Key("key")

	_, ok := v.(*ErrorNode)
	assertTrue(t, ok)
}

func Test_ArrayではMapは扱えない(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	_, err := arr.Map()

	assertError(t, err)
}

func Test_ArrayではMapOrはデフォルト値(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	v := arr.MapOr(map[string]interface{}{
		"key1": "value1",
	})

	val, ok := v["key1"].(string)
	assertTrue(t, ok)
	assertString(t, "value1", val)
}

func Test_ArrayではObjでは値を取り出せない(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	_, err := arr.Obj()

	assertError(t, err)
}

func Test_ArrayはNilではない(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	b := arr.IsNil()

	assertFalse(t, b)
}

func Test_ArrayをRootノードとして生成した時のPathが正しい(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	path := arr.GetPath()

	assertString(t, "<root>", path)
}

func Test_Arrayの要素のPathが正しい(t *testing.T) {

	arr := root(`["a", "b", "c"]`)

	path := arr.At(0).GetPath()

	assertString(t, "<root>[0]", path)
}
