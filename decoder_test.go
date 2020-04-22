package simplejson

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// This test case was copied from encoding/json/stream_test.go
var streamEncoded = `0.1
"hello"
null
true
false
["a","b","c"]
{"ß":"long s","K":"Kelvin"}
3.14
`

func TestDecoder(t *testing.T) {
	r := strings.NewReader(streamEncoded)
	dec := NewDecoder(r)

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err := dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	f, err := j.Float64()
	if err != nil {
		t.Error(err)
	} else if f != 0.1 {
		t.Errorf(`want %v, but got %v`, 0.1, f)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	s, err := j.String()
	if err != nil {
		t.Error(err)
	} else if s != "hello" {
		t.Errorf(`want %v, but got %v`, "hello", s)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	i := j.Interface()
	if i != nil {
		t.Errorf(`want %v, but got %v`, nil, i)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	b, err := j.Bool()
	if err != nil {
		t.Error(err)
	} else if b != true {
		t.Errorf(`want %v, but got %v`, true, b)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	b, err = j.Bool()
	if err != nil {
		t.Error(err)
	} else if b != false {
		t.Errorf(`want %v, but got %v`, false, b)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	a, err := j.Array()
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(a, []interface{}{"a", "b", "c"}) {
		t.Errorf(`want %v, but got %v`, true, a)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	m, err := j.Map()
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(m, map[string]interface{}{"ß": "long s", "K": "Kelvin"}) {
		t.Errorf(`want %v, but got %v`, true, a)
	}

	if !dec.More() {
		t.Error("More must returns true")
	}
	j, err = dec.Decode()
	if err != nil {
		t.Fatal(err)
	}
	f, err = j.Float64()
	if err != nil {
		t.Error(err)
	} else if f != 3.14 {
		t.Errorf(`want %v, but got %v`, 3.14, f)
	}

	if dec.More() {
		t.Error("More must returns false")
	}
	_, err = dec.Decode()
	if err != io.EOF {
		t.Errorf(`want %v, but got %v`, io.EOF, err)
	}
}
