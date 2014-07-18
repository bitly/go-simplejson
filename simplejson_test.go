package simplejson

import (
	"encoding/json"
	"testing"

	"github.com/bmizerany/assert"
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

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	_, ok = js.CheckGet("test")
	assert.Equal(t, true, ok)

	_, ok = js.CheckGet("missing_key")
	assert.Equal(t, false, ok)

	aws := js.Get("test").Get("arraywithsubs")
	assert.NotEqual(t, nil, aws)
	var awsval int
	awsval, _ = aws.Get(0).Get("subkeyone").CheckInt()
	assert.Equal(t, 1, awsval)
	awsval, _ = aws.Get(1).Get("subkeytwo").CheckInt()
	assert.Equal(t, 2, awsval)
	awsval, _ = aws.Get(1).Get("subkeythree").CheckInt()
	assert.Equal(t, 3, awsval)

	i, _ := js.Get("test").Get("int").CheckInt()
	assert.Equal(t, 10, i)

	f, _ := js.Get("test").Get("float").CheckFloat64()
	assert.Equal(t, 5.150, f)

	s, _ := js.Get("test").Get("string").CheckString()
	assert.Equal(t, "simplejson", s)

	b, _ := js.Get("test").Get("bool").CheckBool()
	assert.Equal(t, true, b)

	mi := js.Get("test").Get("int").Int()
	assert.Equal(t, 10, mi)

	mi2 := js.Get("test").Get("missing_int").Int(5150)
	assert.Equal(t, 5150, mi2)

	ms := js.Get("test").Get("string").String()
	assert.Equal(t, "simplejson", ms)

	ms2 := js.Get("test").Get("missing_string").String("fyea")
	assert.Equal(t, "fyea", ms2)

	ma2 := js.Get("test").Get("missing_array").Array([]interface{}{"1", 2, "3"})
	assert.Equal(t, ma2, []interface{}{"1", 2, "3"})

	mm2 := js.Get("test").Get("missing_map").Map(map[string]interface{}{"found": false})
	assert.Equal(t, mm2, map[string]interface{}{"found": false})

	gp, _ := js.Get("test", "string").CheckString()
	assert.Equal(t, "simplejson", gp)

	gp2, _ := js.Get("test", "int").CheckInt()
	assert.Equal(t, 10, gp2)

	gpa, _ := js.Get("test", "string_array", 0).CheckString()
	assert.Equal(t, "asdf", gpa)

	gpa2, _ := js.Get("test", "arraywithsubs", 1, "subkeythree").CheckInt()
	assert.Equal(t, 3, gpa2)

	jm, ok := js.Get("test").CheckJSONMap()
	assert.Equal(t, ok, true)
	jmbool, _ := jm["bool"].CheckBool()
	assert.Equal(t, true, jmbool)

	ja, ok := js.Get("test", "string_array").CheckJSONArray()
	assert.Equal(t, ok, true)
	jastr, _ := ja[0].CheckString()
	assert.Equal(t, "asdf", jastr)

	assert.Equal(t, js.Get("test").Get("bool").Bool(), true)

	js.Set("float2", 300.0)
	assert.Equal(t, js.Get("float2").Float64(), 300.0)

	js.Set("test2", "setTest")
	assert.Equal(t, "setTest", js.Get("test2").String())

	js.Del("test2")
	assert.NotEqual(t, "setTest", js.Get("test2").String())

	js.Get("test").Get("sub_obj").Set("a", 2)
	assert.Equal(t, 2, js.Get("test").Get("sub_obj").Get("a").Int())

	js.Get("test", "sub_obj").Set("a", 3)
	assert.Equal(t, 3, js.Get("test", "sub_obj", "a").Int())

	jmm := js.Get("missing_map").JSONMap(map[string]*Json{"js1": js})
	assert.Equal(t, js, jmm["js1"])

	jma := js.Get("missing_array").JSONArray([]*Json{js})
	assert.Equal(t, js, jma[0])
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

	assert.Equal(t, nil, json.Unmarshal([]byte(raw), val))

	assert.Equal(t, "myobject", val.Name)
	assert.NotEqual(t, nil, val.Params.data)
	s, _ := val.Params.Get("string").CheckString()
	assert.Equal(t, "simplejson", s)

	p, err := json.Marshal(val)
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, json.Unmarshal(p, val2))
	assert.Equal(t, val, val2) // stable
}

func TestSet(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	assert.Equal(t, nil, err)

	js.Set("baz", "bing")

	s, ok := js.Get("baz").CheckString()
	assert.Equal(t, true, ok)
	assert.Equal(t, "bing", s)
}

func TestReplace(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	assert.Equal(t, nil, err)

	err = js.UnmarshalJSON([]byte(`{"baz":"bing"}`))
	assert.Equal(t, nil, err)

	s, ok := js.Get("baz").CheckString()
	assert.Equal(t, true, ok)
	assert.Equal(t, "bing", s)
}

func TestSetPath(t *testing.T) {
	js, err := NewJson([]byte(`{}`))
	assert.Equal(t, nil, err)

	js.SetPath([]string{"foo", "bar"}, "baz")

	s, ok := js.Get("foo", "bar").CheckString()
	assert.Equal(t, true, ok)
	assert.Equal(t, "baz", s)
}

func TestSetPathNoPath(t *testing.T) {
	js, err := NewJson([]byte(`{"some":"data","some_number":1.0,"some_bool":false}`))
	assert.Equal(t, nil, err)

	f := js.Get("some_number").Float64(99.0)
	assert.Equal(t, f, 1.0)

	js.SetPath([]string{}, map[string]interface{}{"foo": "bar"})

	s, ok := js.Get("foo").CheckString()
	assert.Equal(t, true, ok)
	assert.Equal(t, "bar", s)

	f = js.Get("some_number").Float64(99.0)
	assert.Equal(t, f, 99.0)
}

func TestPathWillAugmentExisting(t *testing.T) {
	js, err := NewJson([]byte(`{"this":{"a":"aa","b":"bb","c":"cc"}}`))
	assert.Equal(t, nil, err)

	js.SetPath([]string{"this", "d"}, "dd")

	cases := []struct {
		path    []interface{}
		outcome string
	}{
		{
			path:    []interface{}{"this", "a"},
			outcome: "aa",
		},
		{
			path:    []interface{}{"this", "b"},
			outcome: "bb",
		},
		{
			path:    []interface{}{"this", "c"},
			outcome: "cc",
		},
		{
			path:    []interface{}{"this", "d"},
			outcome: "dd",
		},
	}

	for _, tc := range cases {
		s, ok := js.Get(tc.path...).CheckString()
		assert.Equal(t, true, ok)
		assert.Equal(t, tc.outcome, s)
	}
}

func TestPathWillOverwriteExisting(t *testing.T) {
	// notice how "a" is 0.1 - but then we'll try to set at path a, foo
	js, err := NewJson([]byte(`{"this":{"a":0.1,"b":"bb","c":"cc"}}`))
	assert.Equal(t, nil, err)

	js.SetPath([]string{"this", "a", "foo"}, "bar")

	s, ok := js.Get("this", "a", "foo").CheckString()
	assert.Equal(t, true, ok)
	assert.Equal(t, "bar", s)
}
