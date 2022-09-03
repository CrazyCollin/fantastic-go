package interview

import (
	"bytes"
	"fmt"
	"strings"
)

func TestStringJoin() {
	var s string
	for _, key := range []string{"a", "b", "c"} {
		s += key
	}
	fmt.Println(s)

	s = ""
	s = s + "a"
	s = s + "b"
	s = s + "c"
	fmt.Println(s)

	s = strings.Join([]string{"a", "b", "c"}, "")
	fmt.Println(s)

	byteBuffer := new(bytes.Buffer)
	byteBuffer.Write([]byte("byte buffer"))
	fmt.Println(byteBuffer.String())
}
