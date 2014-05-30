// +build !go1.1

package simplejson

import (
	"encoding/json"
	"errors"
	"io"
)

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error {
	return json.Unmarshal(p, &j.data)
}

// Float64 type asserts to `float64`
func (j *Json) Float64() (float64, error) {
	if i, ok := (j.data).(float64); ok {
		return i, nil
	}
	return -1, errors.New("type assertion to float64 failed")
}

// NewFromReader returns a *Json by decoding from an io.Reader
func NewFromReader(r io.Reader) (*Json, error) {
	j := new(Json)
	dec := json.NewDecoder(r)
	err := dec.Decode(&j.data)
	return j, err
}

// Int type asserts to `float64` then converts to `int`
func (j *Json) Int() (int, error) {
	if f, ok := (j.data).(float64); ok {
		return int(f), nil
	}
	return -1, errors.New("type assertion to float64 failed")
}

// Int64 type asserts to `float64` then converts to `int64`
func (j *Json) Int64() (int64, error) {
	if f, ok := (j.data).(float64); ok {
		return int64(f), nil
	}
	return -1, errors.New("type assertion to float64 failed")
}

// Uint64 type asserts to `float64` then converts to `uint64`
func (j *Json) Uint64() (uint64, error) {
	if f, ok := (j.data).(float64); ok {
		return uint64(f), nil
	}
	return 0, errors.New("type assertion to float64 failed")
}
