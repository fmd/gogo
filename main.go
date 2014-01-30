package main

import (
    "github.com/codegangsta/martini"
    "github.com/fmd/gogo/gogo"
)

func main() {
      m := martini.Classic()
        m.Get("/", func() string {
                return gogo.Board()
                  })
          m.Run()
}
