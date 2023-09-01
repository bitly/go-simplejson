// +build !go1.1

package simplejson

import (
	"encoding/json"
	"io"
)

func NewDecoder(r io.Reader) *Decoder {
	dec := json.NewDecoder(r)
	return (*Decoder)(dec)
}
