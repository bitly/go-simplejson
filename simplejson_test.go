package simplejson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSimplejson(t *testing.T) {
	var ok bool
	var err error

	js, err := NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"string_array_null": ["abc", null, "efg"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"string": "simplejson",
			"bool": true,
			"sub_obj": {"a": 1}
		}
	}`))
	if js == nil {
		t.Fatal("got nil")
	}
	if err != nil {
		t.Fatalf("got err %#v", err)
	}

	_, ok = js.CheckGet("test")
	if !ok {
		t.Errorf("test: got %#v expected true", ok)
	}

	_, ok = js.CheckGet("missing_key")
	if ok {
		t.Errorf("missing_key: got %#v expected false", ok)
	}

	aws := js.Get("test").Get("arraywithsubs")
	if aws == nil {
		t.Fatal("got nil")
	}

	if got, _ := aws.GetIndex(0).Get("subkeyone").Int(); got != 1 {
		t.Errorf("got %#v", got)
	}
	if got, _ := aws.GetIndex(1).Get("subkeytwo").Int(); got != 2 {
		t.Errorf("got %#v", got)
	}
	if got, _ := aws.GetIndex(1).Get("subkeythree").Int(); got != 3 {
		t.Errorf("got %#v", got)
	}

	if i, _ := js.Get("test").Get("int").Int(); i != 10 {
		t.Errorf("got %#v", i)
	}

	if f, _ := js.Get("test").Get("float").Float64(); f != 5.150 {
		t.Errorf("got %#v", f)
	}

	if s, _ := js.Get("test").Get("string").String(); s != "simplejson" {
		t.Errorf("got %#v", s)
	}

	if b, _ := js.Get("test").Get("bool").Bool(); b != true {
		t.Errorf("got %#v", b)
	}

	if mi := js.Get("test").Get("int").MustInt(); mi != 10 {
		t.Errorf("got %#v", mi)
	}

	if mi := js.Get("test").Get("missing_int").MustInt(5150); mi != 5150 {
		t.Errorf("got %#v", mi)
	}

	if s := js.Get("test").Get("string").MustString(); s != "simplejson" {
		t.Errorf("got %#v", s)
	}

	if s := js.Get("test").Get("missing_string").MustString("fyea"); s != "fyea" {
		t.Errorf("got %#v", s)
	}

	a := js.Get("test").Get("missing_array").MustArray([]interface{}{"1", 2, "3"})
	if !reflect.DeepEqual(a, []interface{}{"1", 2, "3"}) {
		t.Errorf("got %#v", a)
	}

	msa := js.Get("test").Get("string_array").MustStringArray()
	if !reflect.DeepEqual(msa, []string{"asdf", "ghjk", "zxcv"}) {
		t.Errorf("got %#v", msa)
	}

	msa = js.Get("test").Get("string_array").MustStringArray([]string{"1", "2", "3"})
	if !reflect.DeepEqual(msa, []string{"asdf", "ghjk", "zxcv"}) {
		t.Errorf("got %#v", msa)
	}

	msa = js.Get("test").Get("missing_array").MustStringArray([]string{"1", "2", "3"})
	if !reflect.DeepEqual(msa, []string{"1", "2", "3"}) {
		t.Errorf("got %#v", msa)
	}

	mm := js.Get("test").Get("missing_map").MustMap(map[string]interface{}{"found": false})
	if !reflect.DeepEqual(mm, map[string]interface{}{"found": false}) {
		t.Errorf("got %#v", mm)
	}

	sa, err := js.Get("test").Get("string_array").StringArray()
	if err != nil {
		t.Fatalf("got err %#v", err)
	}
	if !reflect.DeepEqual(sa, []string{"asdf", "ghjk", "zxcv"}) {
		t.Errorf("got %#v", sa)
	}

	sa, err = js.Get("test").Get("string_array_null").StringArray()
	if err != nil {
		t.Fatalf("got err %#v", err)
	}
	if !reflect.DeepEqual(sa, []string{"abc", "", "efg"}) {
		t.Errorf("got %#v", sa)
	}

	if s, _ := js.GetPath("test", "string").String(); s != "simplejson" {
		t.Errorf("got %#v", s)
	}

	if i, _ := js.GetPath("test", "int").Int(); i != 10 {
		t.Errorf("got %#v", i)
	}

	if b := js.Get("test").Get("bool").MustBool(); b != true {
		t.Errorf("got %#v", b)
	}

	js.Set("float2", 300.0)
	if f := js.Get("float2").MustFloat64(); f != 300.0 {
		t.Errorf("got %#v", f)
	}

	js.Set("test2", "setTest")
	if s := js.Get("test2").MustString(); s != "setTest" {
		t.Errorf("got %#v", s)
	}

	js.Del("test2")
	if s := js.Get("test2").MustString(); s == "setTest" {
		t.Errorf("got %#v", s)
	}

	js.Get("test").Get("sub_obj").Set("a", 2)
	if i := js.Get("test").Get("sub_obj").Get("a").MustInt(); i != 2 {
		t.Errorf("got %#v", i)
	}

	js.GetPath("test", "sub_obj").Set("a", 3)
	if i := js.Get("test").Get("sub_obj").Get("a").MustInt(); i != 3 {
		t.Errorf("got %#v", i)
	}
}

func TestStdlibInterfaces(t *testing.T) {
	val := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})
	val2 := new(struct {
		Name   string `json:"name"`
		Params *Json  `json:"params"`
	})

	raw := `{"name":"myobject","params":{"string":"simplejson"}}`

	if err := json.Unmarshal([]byte(raw), val); err != nil {
		t.Fatalf("err %#v", err)
	}
	if val.Name != "myobject" {
		t.Errorf("got %#v", val.Name)
	}
	if val.Params.data == nil {
		t.Errorf("got %#v", val.Params.data)
	}
	if s, _ := val.Params.Get("string").String(); s != "simplejson" {
		t.Errorf("got %#v", s)
	}

	p, err := json.Marshal(val)
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if err = json.Unmarshal(p, val2); err != nil {
		t.Fatalf("err %#v", err)
	}
	if !reflect.DeepEqual(val, val2) { // stable
		t.Errorf("got %#v expected %#v", val2, val)
	}
}

func TestSet(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	js.Set("baz", "bing")

	s, err := js.GetPath("baz").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "bing" {
		t.Errorf("got %#v", s)
	}
}

func TestReplace(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	err = js.UnmarshalJSON([]byte(`{"baz":"bing"}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	s, err := js.GetPath("baz").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "bing" {
		t.Errorf("got %#v", s)
	}
}

func TestSetPath(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	js.SetPath([]string{"foo", "bar"}, "baz")

	s, err := js.GetPath("foo", "bar").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "baz" {
		t.Errorf("got %#v", s)
	}
}

func TestSetPathNoPath(t *testing.T) {
	js, err := NewJson([]byte(`{"some":"data","some_number":1.0,"some_bool":false}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	if f := js.GetPath("some_number").MustFloat64(99.0); f != 1.0 {
		t.Errorf("got %#v", f)
	}

	js.SetPath([]string{}, map[string]interface{}{"foo": "bar"})

	s, err := js.GetPath("foo").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "bar" {
		t.Errorf("got %#v", s)
	}

	if f := js.GetPath("some_number").MustFloat64(99.0); f != 99.0 {
		t.Errorf("got %#v", f)
	}
}

func TestPathWillAugmentExisting(t *testing.T) {
	js, err := NewJson([]byte(`{"this":{"a":"aa","b":"bb","c":"cc"}}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	js.SetPath([]string{"this", "d"}, "dd")

	cases := []struct {
		path    []string
		outcome string
	}{
		{
			path:    []string{"this", "a"},
			outcome: "aa",
		},
		{
			path:    []string{"this", "b"},
			outcome: "bb",
		},
		{
			path:    []string{"this", "c"},
			outcome: "cc",
		},
		{
			path:    []string{"this", "d"},
			outcome: "dd",
		},
	}

	for _, tc := range cases {
		s, err := js.GetPath(tc.path...).String()
		if err != nil {
			t.Fatalf("err %#v", err)
		}
		if s != tc.outcome {
			t.Errorf("got %#v expected %#v", s, tc.outcome)
		}
	}
}

func TestPathWillOverwriteExisting(t *testing.T) {
	// notice how "a" is 0.1 - but then we'll try to set at path a, foo
	js, err := NewJson([]byte(`{"this":{"a":0.1,"b":"bb","c":"cc"}}`))
	if err != nil {
		t.Fatalf("err %#v", err)
	}

	js.SetPath([]string{"this", "a", "foo"}, "bar")

	s, err := js.GetPath("this", "a", "foo").String()
	if err != nil {
		t.Fatalf("err %#v", err)
	}
	if s != "bar" {
		t.Errorf("got %#v", s)
	}
}
