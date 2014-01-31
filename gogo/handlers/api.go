package handlers

import (
	//"fmt"
	//"net/http"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

type ApiHandler struct {
	Martini *martini.ClassicMartini
	Protocol protocols.IOProtocol
	Backend backends.Backend
}

func (a *ApiHandler) loadRoutes() {
	a.Martini.Get("/api", func(r render.Render) {
		r.JSON(200,"Hello, world!")
	})

	a.Martini.Get("/api/about", func(r render.Render) {
		r.JSON(200,"About the world!")
	})

	a.Martini.Get("/api/contact", func(r render.Render) {
		r.JSON(200,"Cool the world!")
	})
}

func NewApiHandler(m *martini.ClassicMartini, p protocols.IOProtocol, b backends.Backend) *ApiHandler {
	a := &ApiHandler{}
	a.Backend = b
	a.Protocol = p
	a.Martini = m

	a.Martini.Use(render.Renderer())
	a.loadRoutes()
	return a
}
