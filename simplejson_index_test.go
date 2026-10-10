package simplejson

import "testing"

func TestGetIndexBounds(t *testing.T) {
	for _, body := range []string{`[]`, `["first", "last"]`, `null`, `{}`} {
		j, err := NewJson([]byte(body))
		if err != nil {
			t.Fatal(err)
		}
		for _, index := range []int{-int(^uint(0)>>1) - 1, -1, 2, int(^uint(0) >> 1)} {
			if got := j.GetIndex(index); got == nil || got.Interface() != nil {
				t.Errorf("GetIndex(%d) on %s = %#v, want a null Json", index, body, got)
			}
		}
	}
	j, err := NewJson([]byte(`["first", "last"]`))
	if err != nil {
		t.Fatal(err)
	}
	for i, want := range []string{"first", "last"} {
		if got := j.GetIndex(i).MustString(); got != want {
			t.Errorf("GetIndex(%d) = %q, want %q", i, got, want)
		}
	}
	if got := j.GetIndex(-1).Get("missing").MustString("fallback"); got != "fallback" {
		t.Fatalf("chained missing index = %q", got)
	}
}
