// +build go1.1

package simplejson

import (
	"encoding/json"
	"io"
)

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	dec := json.NewDecoder(r)
	dec.UseNumber()
	return (*Decoder)(dec)
}
