package interview

import (
	"fmt"
	"reflect"
)

type Author struct {
	Name  string   `json:"name"`
	Books []string `json:"books"`
}

func TestReflect() {
	t := reflect.TypeOf(Author{})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag.Get("json"))
	}
}
