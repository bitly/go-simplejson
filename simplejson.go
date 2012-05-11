package simplejson

import (
	"encoding/json"
	"errors"
)

// returns the current implementation version
func Version() string {
	return "0.1"
}

type Json struct {
	data interface{}
}

// NewJson returns a pointer to a new `Json` object
// after unmarshaling `body` bytes
func NewJson(body []byte) (*Json, error) {
	j := new(Json)
	err := json.Unmarshal(body, &j.data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// Encode returns it's marshaled data as `[]byte`
func (j *Json) Encode() ([]byte, error) {
	return json.Marshal(&j.data)
}

// Get returns a pointer to a new `Json` object 
// for `key` in it's `map` representation
// 
// useful for chaining operations (to traverse a nested JSON):
//    js.Get("top_level").Get("dict").Get("value").Int()
func (j *Json) Get(key string) *Json {
	m, err := j.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Json{val}
		}
	}
	return &Json{nil}
}

// CheckGet returns a pointer to a new `Json` object and
// a `bool` identifying success or failure
//
// useful for chained operations when success is important:
//    if data, ok := js.Get("top_level").CheckGet("inner"); ok {
//        log.Println(data)
//    }
func (j *Json) CheckGet(key string) (*Json, bool) {
	m, err := j.Map()
	if err == nil {
		if val, ok := m[key]; ok {
			return &Json{val}, true
		}
	}
	return nil, false
}

// Map type asserts to `map`
func (j *Json) Map() (map[string]interface{}, error) {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

// String type asserts to `string`
func (j *Json) String() (string, error) {
	if s, ok := (j.data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

// Float64 type asserts to `float64`
func (j *Json) Float64() (float64, error) {
	if i, ok := (j.data).(float64); ok {
		return i, nil
	}
	return -1, errors.New("type assertion to float64 failed")
}

// Int type asserts to `int`
func (j *Json) Int() (int, error) {
	if f, ok := (j.data).(float64); ok {
		i := int(f)
		return i, nil
	}
	return -1, errors.New("type assertion to int failed")
}

// Int type asserts to `int64`
func (j *Json) Int64() (int64, error) {
	if f, ok := (j.data).(float64); ok {
		i := int64(f)
		return i, nil
	}
	return -1, errors.New("type assertion to int failed")
}

// Bytes type asserts to `[]byte`
func (j *Json) Bytes() ([]byte, error) {
	if b, ok := (j.data).([]byte); ok {
		return b, nil
	}
	return nil, errors.New("type assertion to []byte failed")
}
