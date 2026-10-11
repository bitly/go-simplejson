package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestIntNumberRange(t *testing.T) {
	max := int64(^uint(0) >> 1)
	min := -max - 1
	cases := []struct {
		name string
		raw  string
		want int
		err  error
	}{
		{"max", strconv.FormatInt(max, 10), int(max), nil},
		{"min", strconv.FormatInt(min, 10), int(min), nil},
		{"above_max", strconv.FormatUint(uint64(max)+1, 10), int(max), strconv.ErrRange},
		{"below_min", "-" + strconv.FormatUint(uint64(max)+2, 10), int(min), strconv.ErrRange},
		{"fraction", "1.5", 0, strconv.ErrSyntax},
		{"exponent", "1e2", 0, strconv.ErrSyntax},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			js, err := NewJson([]byte(tc.raw))
			if err != nil {
				t.Fatal(err)
			}
			got, err := js.Int()
			if got != tc.want || !errors.Is(err, tc.err) {
				t.Fatalf("Int(%s): got %d, %v; want %d, %v", tc.raw, got, err, tc.want, tc.err)
			}
			if tc.err != nil && js.MustInt(42) != 42 {
				t.Fatal("MustInt did not return the default after a conversion error")
			}
		})
	}
}

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
