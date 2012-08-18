### go-simplejson

a Go package to interact with arbitrary JSON

### importing

    import simplejson github.com/bitly/go-simplejson

### go doc

```go
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

func (j *Json) Array() ([]interface{}, error)
    Array type asserts to an `array`

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
    Encode returns its marshaled data as `[]byte`

func (j *Json) Float64() (float64, error)
    Float64 type asserts to `float64`

func (j *Json) Get(key string) *Json
    Get returns a pointer to a new `Json` object for `key` in its `map`
    representation

    useful for chaining operations (to traverse a nested JSON):

    js.Get("top_level").Get("dict").Get("value").Int()

func (j *Json) Int() (int, error)
    Int type asserts to `float64` then converts to `int`

func (j *Json) Int64() (int64, error)
    Int type asserts to `float64` then converts to `int64`

func (j *Json) Map() (map[string]interface{}, error)
    Map type asserts to `map`

func (j *Json) MustFloat64(args ...float64) float64
    MustFloat64 guarantees the return of a `float64` (with optional default)

    useful when you explicitly want a `float64` in a single value return
    context:

    myFunc(js.Get("param1").MustFloat64(), js.Get("optional_param").MustFloat64(5.150))

func (j *Json) MustInt(args ...int) int
    MustInt guarantees the return of an `int` (with optional default)

    useful when you explicitly want an `int` in a single value return
    context:

    myFunc(js.Get("param1").MustInt(), js.Get("optional_param").MustInt(5150))

func (j *Json) MustString(args ...string) string
    MustString guarantees the return of a `string` (with optional default)

    useful when you explicitly want a `string` in a single value return
    context:

    myFunc(js.Get("param1").MustString(), js.Get("optional_param").MustString("my_default"))

func (j *Json) String() (string, error)
    String type asserts to `string`
```
