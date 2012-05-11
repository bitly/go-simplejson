### go-simplejson

a dead simple library to interact with arbitrary JSON in Go

### importing

    import simplejson github.com/bitly/go-simplejson


### go doc

```
FUNCTIONS

func Version() string
    returns the current implementation version


TYPES

type Json struct {
    // contains filtered or unexported fields
}

func NewJson(body []byte) (*Json, error)
    NewJson returns a pointer to a new `Json` object after unmarshaling
    `body` bytes

func (j *Json) Bytes() ([]byte, error)
    Bytes type asserts to `[]byte`

func (j *Json) CheckGet(key string) (*Json, bool)
    CheckGet returns a pointer to a new `Json` object and a `bool`
    identifying success or failure

    useful for chained operations when success is important:

	if data, ok := js.Get("top_level").CheckGet("inner"); ok {
	    log.Println(data)
	}

func (j *Json) Encode() ([]byte, error)
    Encode returns it's marshaled data as `[]byte`

func (j *Json) Float64() (float64, error)
    Float64 type asserts to `float64`

func (j *Json) Get(key string) *Json
    Get returns a pointer to a new `Json` object for `key` in it's `map`
    representation

    useful for chaining operations (to traverse a nested JSON):

	js.Get("top_level").Get("dict").Get("value").Int()

func (j *Json) Int() (int, error)
    Int type asserts to `int`

func (j *Json) Int64() (int64, error)
    Int type asserts to `int64`

func (j *Json) Map() (map[string]interface{}, error)
    Map type asserts to `map`

func (j *Json) String() (string, error)
    String type asserts to `string`
```
