// +build !go1.1

package simplejson

import (
	"bytes"
	"github.com/bmizerany/assert"
	"strconv"
	"testing"
)

func TestNewFromReader(t *testing.T) {
	buf := bytes.NewBuffer([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 8000000000
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
		case float64:
			iv = int(v.(float64))
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		assert.Equal(t, i+1, iv)
	}

	ma := js.Get("test").Get("array").Array()
	assert.Equal(t, ma, []interface{}{float64(1), "2", float64(3)})

	mm := js.Get("test").Get("arraywithsubs").Get(0).Map()
	assert.Equal(t, mm, map[string]interface{}{"subkeyone": float64(1)})

	assert.Equal(t, js.Get("test").Get("bignum").Int64(), int64(8000000000))
}

func TestSimplejsonGo10(t *testing.T) {
	js, err := NewJSON([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 8000000000
		}
	}`))

	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)

	arr, _ := js.Get("test").Get("array").CheckArray()
	assert.NotEqual(t, nil, arr)
	for i, v := range arr {
		var iv int
		switch v.(type) {
		case float64:
			iv = int(v.(float64))
		case string:
			iv, _ = strconv.Atoi(v.(string))
		}
		assert.Equal(t, i+1, iv)
	}

	ma := js.Get("test").Get("array").Array()
	assert.Equal(t, ma, []interface{}{float64(1), "2", float64(3)})

	mm := js.Get("test").Get("arraywithsubs").Get(0).Map()
	assert.Equal(t, mm, map[string]interface{}{"subkeyone": float64(1)})

	assert.Equal(t, js.Get("test").Get("bignum").Int64(), int64(8000000000))
}
