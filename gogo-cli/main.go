package main

import (
    "fmt"
)

func main() {
    c := Client{}
    c.ParseFlags()
    fmt.Println("OK")
}
