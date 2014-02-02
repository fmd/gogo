package main

import (
    //"github.com/eaigner/hood"
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "github.com/codegangsta/martini-contrib/sessions"
    //auth "github.com/codegangsta/martini-contrib/sessionauth"
)
var b *Backend
var s sessions.CookieStore
var m *martini.ClassicMartini

func routes(m *martini.ClassicMartini) {
    m.Get("/api", func(r render.Render) {
        r.JSON(200,"Cool the world!")
    })

    m.Get("/api/user/create", func(r render.Render) {
        u := &User{}
        u.Email = "fareeddudhia@gmail.com"
        u.Username = "fmd"
        err := b.Save(u)
        if err != nil {
            panic(err)
        }
        r.JSON(200,map[string]int{"id":int(u.Id)})
    })
}

func main() {
    var err error
    b, err = NewBackend()
    if err != nil {
        panic(err)
    }

    s = sessions.NewCookieStore([]byte("secret123"))

    m = martini.Classic()
    m.Use(sessions.Sessions("gogo_session", s))
    m.Use(render.Renderer())

    routes(m)

    m.Run()
}