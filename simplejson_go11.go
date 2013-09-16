// +build go1.1

package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewReader(p))
	dec.UseNumber()

	return dec.Decode(&j.data)
}

// Float64 type asserts to `json.Number` then converts to `float64`
func (j *Json) Float64() (float64, error) {
	if n, ok := (j.data).(json.Number); ok {
		return n.Float64()
	}
	return -1, errors.New("type assertion to float64 failed")
}

// Int type asserts to `json.Number` then converts to `int`
func (j *Json) Int() (int, error) {
	if n, ok := (j.data).(json.Number); ok {
		i, ok := n.Int64()
		return int(i), ok
	}

	return -1, errors.New("type assertion to float64 failed")
}

// Int type asserts to `json.Number` then converts to `int64`
func (j *Json) Int64() (int64, error) {
	if n, ok := (j.data).(json.Number); ok {
		return n.Int64()
	}

	return -1, errors.New("type assertion to float64 failed")
}
