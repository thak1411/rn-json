package rnjson

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var res Rnjson
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "foo5",
			"foo5": {
				"foo6": "foo7"
			}
		}
	}`
	res, _ = Unmarshal(dat)
	if get, ok := Get(res, "foo"); !ok {
		t.Errorf("Get(res, \"foo\") = %q, want = %q", get, "bar")
	}
}

func TestGet(t *testing.T) {
	var res Rnjson
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "foo5",
			"foo5": {
				"foo6": "foo7"
			}
		}
	}`
	res, _ = Unmarshal(dat)
	if get, ok := Get(res, "foo"); !ok {
		t.Errorf("Get(res, \"foo\") = %q, want = %q", get, "bar")
	}
}
