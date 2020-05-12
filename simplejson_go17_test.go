// +build go1.7 go1.8 go1.9 go1.10

package simplejson

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleJsonNotEscapeHTML(t *testing.T) {
	//Use New Constructor
	buf := []byte(`{
		"test": {
			"array": [1, "2", 3],
			"arraywithsubs": [
				{"subkeyone": 1},
				{"subkeytwo": 2, "subkeythree": 3}
			],
			"bignum": 9223372036854775807,
			"uint64": 18446744073709551615,
			"u": "a=1234&b=456&c=789"
		}
	}`)
	js, err := NewJson(buf)

	//Standard Test Case
	assert.NotEqual(t, nil, js)
	assert.Equal(t, nil, err)
	js.SetHtmlEscape(false)
	txt, err := js.Encode()
	idx := bytes.Index(txt, []byte(`\u0026`))
	assert.NotEqual(t, true, idx > 0)

	js.SetHtmlEscape(true)
	txt, err = js.Encode()
	idx = bytes.Index(txt, []byte(`\u0026`))
	assert.NotEqual(t, false, idx > 0)
}
