package rnjson

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
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
	if _, err := Unmarshal(dat); err != nil {
		t.Errorf("Unmarshal Error %q\n", err)
	}
}

func TestUnmarshal2(t *testing.T) {
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "foo5",
			"foo5": {
				"foo6": "foo7",
			}
		}
	}`
	if _, err := Unmarshal(dat); err == nil {
		t.Errorf("Unmarshal Error %q\n", err)
	}
}

func TestUnmarshal3(t *testing.T) {
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "foo5"
			"foo5": {
				"foo6": "foo7"
			}
		}
	}`
	if _, err := Unmarshal(dat); err == nil {
		t.Errorf("Unmarshal Error %q\n", err)
	}
}

func TestUnmarshal4(t *testing.T) {
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "foo5"
			"foo5": {
				"foo6": "foo7"
			}
		}
	`
	if _, err := Unmarshal(dat); err == nil {
		t.Errorf("Unmarshal Error %q\n", err)
	}
}

func TestGet(t *testing.T) {
	var res map[string]interface{}
	dat := `{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "bar4",
			"foo5": {
				"foo6": "bar6"
			}
		}
	}`
	res, _ = Unmarshal(dat)
	if get := Get(res, "foo"); get != "bar" {
		t.Errorf("Get(res, \"foo\") = %q, want = %q\n", get, "bar")
	}
	if get := Get(res, "foo2"); get != "bar2" {
		t.Errorf("Get(res, \"foo2\") = %q, want = %q\n", get, "bar2")
	}
	{
		get := Get(res, "foo3")
		if get := Get(get, "foo5.foo6"); get != "bar6" {
			t.Errorf("Get(get, \"foo5.foo6\") = %q, want = %q\n", get, "bar6")
		}
		get = Get(get, "foo4")
		if get != "bar4" {
			t.Errorf("Get(get, \"foo4\") = %q, want = %q\n", get, "bar4")
		}
	}
	if get := Get(res, "foo3.foo5.foo6"); get != "bar6" {
		t.Errorf("Get(res, \"foo3.foo5.foo6\") = %q, want = %q\n", get, "bar6")
	}
	if get := Get(res, "foo3.foo5.foo7"); get != nil {
		t.Errorf("Get(res, \"foo3.foo5.foo7\") = %q, want = %q\n", get, "nil")
	}
	if get := Get(res, "foo3.foo5.foo7.foo8"); get != nil {
		t.Errorf("Get(res, \"foo3.foo5.foo7.foo8\") = %q, want = %q\n", get, "nil")
	}
	if get := Get(res, "foo3.foo5.foo7.foo8.foo9"); get != nil {
		t.Errorf("Get(res, \"foo3.foo5.foo7.foo8.foo9\") = %q, want = %q\n", get, "nil")
	}
	if get := Get(res, "foo2.foo3"); get != nil {
		t.Errorf("Get(res, \"foo2.foo3\") = %q, want = %q\n", get, "nil")
	}
	if get := Get(res, "foo4"); get != nil {
		t.Errorf("Get(res, \"foo4\") = %q, want = %q\n", get, "nil")
	}
}
