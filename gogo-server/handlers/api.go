package handlers

import (
	//"fmt"
	"net/http"
	"github.com/fmd/gogo/gogo/protocols"
	"github.com/fmd/gogo/gogo-server/models"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/codegangsta/martini-contrib/sessionauth"
)

type ApiHandler struct {
	Engine *Engine
	Store sessions.CookieStore
}

func (a *ApiHandler) needsAuth(r render.Render, u sessionauth.User, req *http.Request) {
	if u.IsAuthenticated() == false {
		r.JSON(403,map[string]string{"status":"unauthorized"})
		return
	}
}

func (a *ApiHandler) loadRoutes() {
	a.Martini.Get("/api/user/create", func (s sessions.Session, r render.Render) {
		u := &models.User{}
		u.Email = "fareed@3ev.com"
		u.Username = "fmd"

		a.Backend.Save(u)
	})

	a.Martini.Get("/api/login", func(s sessions.Session, u sessionauth.User, r render.Render) {
		err := sessionauth.AuthenticateSession(s, u)
		if err != nil {
			r.JSON(500, "Internal server error")
			return
		}

		r.JSON(200, u)
	})

	a.Martini.Get("/api/logout", a.needsAuth, func(s sessions.Session, u sessionauth.User, r render.Render) {
		sessionauth.Logout(s, u)
		r.JSON(200, map[string]string{"status":"success"})
	})

	a.Martini.Get("/api/game", a.needsAuth, func(r render.Render) {
		r.JSON(200,"Cool the world!")
	})

	a.Martini.Get("/api/beefy", func(r render.Render) {
		r.JSON(200,"Beef the world!")
	})
}

func NewApiHandler(m *martini.ClassicMartini, p protocols.IOProtocol, b *Backend) *ApiHandler {
	a := &ApiHandler{}

	a.Backend = b
	a.Protocol = p
	a.Martini = m
	a.Store = sessions.NewCookieStore([]byte("secret123"))

	a.Martini.Use(sessions.Sessions("gogo_session", a.Store))
	a.Martini.Use(sessionauth.SessionUser(func () sessionauth.User {
		return &models.User{}
	}))

	a.loadRoutes()
	return a
}
