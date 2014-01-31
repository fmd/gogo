package handlers

import (
	//"fmt"
	//"net/http"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/codegangsta/martini-contrib/sessionauth"
)

type ApiHandler struct {
	Martini *martini.ClassicMartini
	Protocol protocols.IOProtocol
	Backend backends.Backend
	Store sessions.CookieStore
}

type MyUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	authenticated bool `json:"-"`
}

func (u *MyUser) IsAuthenticated() bool {
	return u.authenticated
}

func (u *MyUser) Login() {
	u.authenticated = true
}

func (u *MyUser) Logout() {
	u.authenticated = false
}

func (u *MyUser) UniqueId() interface{} {
	return u.Id
}

func (u *MyUser) GetById(id interface{}) error {
	u.Id = id.(int)
	u.Name = "My Test User"
	u.Age = 42

	return nil
}


func (a *ApiHandler) loadRoutes() {
	a.Martini.Get("/api/login", func(r render.Render) {
		r.JSON(200,"Hello, world!")
	})

	a.Martini.Get("/api/about", func(r render.Render) {
		r.JSON(200,"About the world!")
	})

	a.Martini.Get("/api/contact", func(r render.Render) {
		r.JSON(200,"Cool the world!")
	})

	a.Martini.Get("/api/beefy", func(r render.Render) {
		r.JSON(200,"Beef the world!")
	})
}

func NewApiHandler(m *martini.ClassicMartini, p protocols.IOProtocol, b backends.Backend) *ApiHandler {
	a := &ApiHandler{}
	a.Backend = b
	a.Protocol = p
	a.Martini = m
	a.Store = sessions.NewCookieStore([]byte("secret123"))

	a.Martini.Use(sessions.Sessions("go_session", a.Store))
	a.Martini.Use(sessionauth.SessionUser(func () sessionauth.User {
		return &MyUser{}
	}))

	a.loadRoutes()
	return a
}
