package main

import (
    "strconv"
    "net/http"
    //"github.com/eaigner/hood"
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/render"
    "github.com/codegangsta/martini-contrib/sessions"
    //auth "github.com/codegangsta/martini-contrib/sessionauth"
)
var b *Backend
var s sessions.CookieStore
var m *martini.ClassicMartini

func GetHomepage(req *http.Request, r render.Render) {
    r.HTML(200,"index",nil)
}

func GetUser(req *http.Request, p martini.Params, r render.Render) {
    id, err := strconv.Atoi(p["id"])

    if err != nil {
        r.JSON(500,"Internal server error")
    }

    var users []User
    b.Hd.Where("id", "=", id).Limit(1).Find(&users)

    if len(users) != 1 {
        r.JSON(500,"Internal server error")   
    }
    
    r.JSON(200,users[0])
}

func UpdateUser(req *http.Request, r render.Render) {

}

func CreateUser(req *http.Request, r render.Render) {
    u := &User{}
    u.Email = "fareeddudhia@gmail.com"
    u.Username = "fmd"
    err := b.Save(u)
    if err != nil {
        r.JSON(500,"Internal server error")
    }
    r.JSON(200,u)
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

    m.Get(`/`, GetHomepage)
    m.Post(`/api/user/create`, CreateUser)
    m.Get(`/api/user/:id`, GetUser)
    m.Post(`/api/user/:id/update`, UpdateUser)

    m.Run()
}