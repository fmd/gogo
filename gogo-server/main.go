package main

import (
	"fmt"
)

func main() {
	s := Server{}
	err := s.Init()
	if err != nil {
		panic(err)
	}
	s.ServeForever()
}
