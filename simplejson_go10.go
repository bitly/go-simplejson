// +build !go1.1

package simplejson

import (
	"encoding/json"
	"io"
	"reflect"
)

// NewFromReader returns a *JSON by decoding from an io.Reader
func NewFromReader(r io.Reader) (*JSON, error) {
	j := new(JSON)
	dec := json.NewDecoder(r)
	err := dec.Decode(&j.data)
	return j, err
}

// Implements the json.Unmarshaler interface.
func (j *JSON) UnmarshalJSON(p []byte) error {
	return json.Unmarshal(p, &j.data)
}

// CheckFloat64 coerces into a float64
func (j *JSON) CheckFloat64() (float64, bool) {
	switch j.data.(type) {
	case float32, float64:
		return reflect.ValueOf(j.data).Float(), true
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(j.data).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(j.data).Uint()), true
	}
	return 0, false
}

// CheckInt coerces into an int
func (j *JSON) CheckInt() (int, bool) {
	switch j.data.(type) {
	case float32, float64:
		return int(reflect.ValueOf(j.data).Float()), true
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(j.data).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(j.data).Uint()), true
	}
	return 0, false
}

// CheckInt64 coerces into an int64
func (j *JSON) CheckInt64() (int64, bool) {
	switch j.data.(type) {
	case float32, float64:
		return int64(reflect.ValueOf(j.data).Float()), true
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(j.data).Int(), true
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(j.data).Uint()), true
	}
	return 0, false
}

// CheckUint64 coerces into an uint64
func (j *JSON) CheckUint64() (uint64, bool) {
	switch j.data.(type) {
	case float32, float64:
		return uint64(reflect.ValueOf(j.data).Float()), true
	case int, int8, int16, int32, int64:
		return uint64(reflect.ValueOf(j.data).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(j.data).Uint(), true
	}
	return 0, false
}
