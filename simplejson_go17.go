// +build go1.7 go1.8 go1.9 go1.10

package simplejson

import (
	"bytes"
	"encoding/json"
)

func (j *Json) SetHtmlEscape(escape bool) {
	j.escapeHtml = escape
}

// EncodePretty returns its marshaled data as `[]byte` with indentation
func (j *Json) EncodePretty() ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(j.escapeHtml)
	enc.SetIndent("", "  ")
	err := enc.Encode(j.data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

// Implements the json.Marshaler interface.
func (j *Json) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(j.escapeHtml)
	err := enc.Encode(j.data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
