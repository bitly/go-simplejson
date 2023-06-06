package simplejson

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"testing"
)

func TestNewFromReader(t *testing.T) {
	//Use New Constructor
	buf := bytes.NewBuffer([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807,
			"uint64": 18446744073709551615
		}
	}`))
	js, err := NewFromReader(buf)

	//Standard Test Case
	if js == nil {
		t.Fatal("got nil")
	}
	if err != nil {
		t.Fatalf("got err %#v", err)
	}

	arr, _ := js.Get("test").Get("array").Array()
	if arr == nil {
		t.Fatal("got nil")
	}
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			if err != nil {
				t.Fatalf("got err %#v", err)
			}
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		if iv != i+1 {
			t.Errorf("got %#v expected %#v", iv, i+1)
		}
	}

	if ma := js.Get("test").Get("array").MustArray(); !reflect.DeepEqual(ma, []interface{}{json.Number("1"), "2", json.Number("3")}) {
		t.Errorf("got %#v", ma)
	}

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	if !reflect.DeepEqual(mm, map[string]interface{}{"subkeyone": json.Number("1")}) {
		t.Errorf("got %#v", mm)
	}

	if n := js.Get("test").Get("bignum").MustInt64(); n != int64(9223372036854775807) {
		t.Errorf("got %#v", n)
	}
	if n := js.Get("test").Get("uint64").MustUint64(); n != uint64(18446744073709551615) {
		t.Errorf("got %#v", n)
	}
}

func TestSimplejsonGo11(t *testing.T) {
	js, err := NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807,
			"uint64": 18446744073709551615
		}
	}`))

	if js == nil {
		t.Fatal("got nil")
	}
	if err != nil {
		t.Fatalf("got err %#v", err)
	}

	arr, _ := js.Get("test").Get("array").Array()
	if arr == nil {
		t.Fatal("got nil")
	}
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			if err != nil {
				t.Fatalf("got err %#v", err)
			}
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		if iv != i+1 {
			t.Errorf("got %#v expected %#v", iv, i+1)
		}
	}

	if ma := js.Get("test").Get("array").MustArray(); !reflect.DeepEqual(ma, []interface{}{json.Number("1"), "2", json.Number("3")}) {
		t.Errorf("got %#v", ma)
	}

	mm := js.Get("test").Get("arraywithsubs").GetIndex(0).MustMap()
	if !reflect.DeepEqual(mm, map[string]interface{}{"subkeyone": json.Number("1")}) {
		t.Errorf("got %#v", mm)
	}
	if n := js.Get("test").Get("bignum").MustInt64(); n != int64(9223372036854775807) {
		t.Errorf("got %#v", n)
	}
	if n := js.Get("test").Get("uint64").MustUint64(); n != uint64(18446744073709551615) {
		t.Errorf("got %#v", n)
	}
}
