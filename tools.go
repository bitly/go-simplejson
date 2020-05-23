package simplejson

import (
	"bytes"
	"fmt"
	"strings"
)

var logbuff bytes.Buffer

func writeLog(frm string, arg ...interface{}) {
	logbuff.WriteString(fmt.Sprintf(frm, arg...))
	logbuff.WriteString("\n")
}

func (j *Json) Merge(data *Json) error {
	dmp, err := data.Map()
	if err != nil {
		return err
	}

	for k, v := range dmp {
		switch v.(type) {
		case map[string]interface{}:
			if _, f := j.CheckGet(k); !f {
				j.Set(k, make(map[string]interface{}))
			}
			j.Get(k).Merge(Wrap(v))
		default:
			j.Set(k, v)
		}
	}
	return nil
}

func (j *Json) compilePath(path string, stack []string) (*Json, error) {
	for _, v := range stack {
		if v == path {
			return nil, fmt.Errorf("Circular reference: %s in %s", path, stack)
		}
	}
	stack = append(stack, path)
	if branch := j.GetPath(strings.Split(path, "/")...); !branch.Empty() {
		return branch.compile(j, stack), nil
	}
	return nil, fmt.Errorf("Path not found: %s", path)
}

func (j *Json) compile(root *Json, stack []string) *Json {
	switch j.data.(type) {
	case map[string]interface{}:
		out := New()
		acc := New()
		mp, _ := j.Map()
		for k, v := range mp {
			if k != "$include" {
				if str, ok := v.(string); ok && len(str) > 0 && str[0] == 36 {
					path := str[1:len(str)]
					br, err := root.compilePath(path, stack)
					if err != nil {
						writeLog(err.Error())
						continue
					}
					out.Set(k, br.Interface())
				} else {
					br := Wrap(v).compile(root, stack)
					out.Set(k, br.Interface())
				}
			}
		}
		if fld, f := mp["$include"]; f {
			if str, ok := fld.(string); ok {
				fld = make([]interface{}, 1)
				(fld.([]interface{}))[0] = str
			}
			if arr, ok := fld.([]interface{}); ok {
				for _, v := range arr {
					if path, ok := v.(string); ok {
						br, err := root.compilePath(path, stack)
						if err != nil {
							writeLog(err.Error())
							continue
						}
						acc.Merge(br)
					}
				}
			}
		}
		if acc.Len() > 0 {
			acc.Merge(out)
			out = acc
		}
		return out
	case []interface{}:
		out := NewArray(j.Len())
		mp, _ := j.Array()
		for _, v := range mp {
			br := Wrap(v).compile(root, stack)
			out.AddArray(br)
		}
		return out
	case string:
		if str, _ := j.String(); len(str) > 0 && str[0] == 36 {
			path := str[1:len(str)]
			br, err := root.compilePath(path, stack)
			if err != nil {
				writeLog(err.Error())
				return j
			}
			return br
		}
	}
	return j
}

func (j *Json) Compile() (*Json, error) {
	logbuff.Reset()
	stack := make([]string, 0, 10)
	var err error
	out := j.compile(j, stack)
	if logbuff.Len() > 0 {
		err = fmt.Errorf(logbuff.String())
	}
	return out, err
}
