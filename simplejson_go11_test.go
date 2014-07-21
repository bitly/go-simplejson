// +build go1.1

package simplejson

import (
	"bytes"
	"encoding/json"
	"github.com/bmizerany/assert"
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
	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	arr, _ := js.Get("test").Get("array").CheckArray()
	assert.NotEqual(t, nil, arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			assert.Equal(t, nil, err)
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		assert.Equal(t, i+1, iv)
	}

	ma := js.Get("test").Get("array").Array()
	assert.Equal(t, ma, []interface{}{json.Number("1"), "2", json.Number("3")})

	mm := js.Get("test").Get("arraywithsubs").Get(0).Map()
	assert.Equal(t, mm, map[string]interface{}{"subkeyone": json.Number("1")})

	assert.Equal(t, js.Get("test").Get("bignum").Int64(), int64(9223372036854775807))
	assert.Equal(t, js.Get("test").Get("uint64").Uint64(), uint64(18446744073709551615))
}

func TestSimplejsonGo11(t *testing.T) {
	js, err := NewJSON([]byte(`{
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

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	arr, _ := js.Get("test").Get("array").CheckArray()
	assert.NotEqual(t, nil, arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case json.Number:
			i64, err := v.(json.Number).Int64()
			assert.Equal(t, nil, err)
			iv = int(i64)
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		assert.Equal(t, i+1, iv)
	}

	ma := js.Get("test").Get("array").Array()
	assert.Equal(t, ma, []interface{}{json.Number("1"), "2", json.Number("3")})

	mm := js.Get("test").Get("arraywithsubs").Get(0).Map()
	assert.Equal(t, mm, map[string]interface{}{"subkeyone": json.Number("1")})

	assert.Equal(t, js.Get("test").Get("bignum").Int64(), int64(9223372036854775807))
	assert.Equal(t, js.Get("test").Get("uint64").Uint64(), uint64(18446744073709551615))
}
