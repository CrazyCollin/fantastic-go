package main

import "fantastic-go/design-pattern/singleton"

func main() {
	s := singleton.GetInstance()
	s.GetString()
}
