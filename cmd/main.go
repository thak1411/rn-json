package main

import (
	"fmt"

	rnjson "github.com/thak1411/rnjson"
)

func main() {
	dat := `
	{
		"foo": "bar",
		"foo2": "bar2",
		"foo3": {
			"foo4": "bar4",
			"foo5": {
				"foo6": "bar6"
			}
		}
	}
	`
	result, err := rnjson.Unmarshal(dat)
	if err != nil {
		panic(err)
	}
	fmt.Println(rnjson.Get(result, "foo"))
	fmt.Println(rnjson.Get(result, "foo3.foo5"))
	fmt.Println(rnjson.Get(result, "foo3.foo5.foo6"))
	fmt.Println(rnjson.Get(result, "foo3.foo5.foo6.foo8"))
	fmt.Println(rnjson.Get(result, "foo3.foo5.foo6.foo8.foo9"))
}
