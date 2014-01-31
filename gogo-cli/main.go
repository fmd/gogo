package main

import (
	"fmt"
)

func main() {
	c := Client{}
	err := c.Init()
    if err != nil {
        panic(err)
    }
    
	fmt.Println("OK")
}
