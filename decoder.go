package simplejson

import (
	"encoding/json"
	"io"
)

// Decoder is a simple wrapper of encoding/json's Decoder.
type Decoder json.Decoder

// Decode reads the next JSON-encoded value from its input and returns it.
func (dec *Decoder) Decode() (*Json, error) {
	j := new(Json)
	err := (*json.Decoder)(dec).Decode(&j.data)
	return j, err
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (dec *Decoder) Buffered() io.Reader {
	return (*json.Decoder)(dec).Buffered()
}

// More reports whether there is another element in the
// current array or object being parsed.
func (dec *Decoder) More() bool {
	return (*json.Decoder)(dec).More()
}
