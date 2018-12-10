package gostructureless

import (
	"encoding/json"
	"testing"
)

func root(str string) Value {

	var data interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}

	return Root(data)
}

func assertError(t *testing.T, err error) {
	if err == nil {
		t.Error("must be error")
	}
}

func assertNotError(t *testing.T, err error) {
	if err != nil {
		t.Error(err.Error())
	}
}

func assertTrue(t *testing.T, b bool) {
	if !b {
		t.Errorf("expected %t, but actual %t", true, b)
	}
}

func assertFalse(t *testing.T, b bool) {
	if b {
		t.Errorf("expected %t, but actual %t", false, b)
	}
}

func assertString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected %s, but actual %s", expected, actual)
	}
}

func assertInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %d, but actual %d", expected, actual)
	}
}

func assertInt64(t *testing.T, expected, actual int64) {
	if expected != actual {
		t.Errorf("expected %d, but actual %d", expected, actual)
	}
}

func assertFloat(t *testing.T, expected, actual float64) {
	if expected != actual {
		t.Errorf("expected %f, but actual %f", expected, actual)
	}
}

func assertFloat64(t *testing.T, expected, actual float64) {
	if expected != actual {
		t.Errorf("expected %f, but actual %f", expected, actual)
	}
}
