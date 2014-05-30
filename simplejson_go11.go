// +build go1.1

package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Implements the json.Unmarshaler interface.
func (j *Json) UnmarshalJSON(p []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(p))
	dec.UseNumber()
	return dec.Decode(&j.data)
}

// NewFromReader returns a *Json by decoding from an io.Reader
func NewFromReader(r io.Reader) (*Json, error) {
	j := new(Json)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	err := dec.Decode(&j.data)
	return j, err
}

// Float64 type asserts to `json.Number` then converts to `float64`
func (j *Json) Float64() (float64, error) {
	if n, ok := (j.data).(json.Number); ok {
		return n.Float64()
	}
	if f, ok := (j.data).(float64); ok {
		return f, nil
	}
	return -1, errors.New("type assertion to json.Number failed")
}

// Int type asserts to `json.Number` then converts to `int`
func (j *Json) Int() (int, error) {
	if n, ok := (j.data).(json.Number); ok {
		i, ok := n.Int64()
		return int(i), ok
	}
	if f, ok := (j.data).(float64); ok {
		return int(f), nil
	}
	return -1, errors.New("type assertion to json.Number failed")
}

// Int64 type asserts to `json.Number` then converts to `int64`
func (j *Json) Int64() (int64, error) {
	if n, ok := (j.data).(json.Number); ok {
		return n.Int64()
	}
	if f, ok := (j.data).(float64); ok {
		return int64(f), nil
	}
	return -1, errors.New("type assertion to json.Number failed")
}

// Uint64 type asserts to `json.Number` then converts to `uint64`
func (j *Json) Uint64() (uint64, error) {
	if n, ok := (j.data).(json.Number); ok {
		u, err := strconv.ParseUint(n.String(), 10, 64)
		return u, err
	}
	if f, ok := (j.data).(float64); ok {
		return uint64(f), nil
	}
	return 0, errors.New("type assertion to json.Number failed")
}
