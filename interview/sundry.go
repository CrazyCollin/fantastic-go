package interview

import "fmt"

type Set map[string]struct{}

func TestSet() {
	set := make(Set)
	for _, key := range []string{"a", "b", "c"} {
		set[key] = struct{}{}
	}
	//get "a" from set
	if _, ok := set["a"]; ok {
		fmt.Println("get a from set")
	}
}
