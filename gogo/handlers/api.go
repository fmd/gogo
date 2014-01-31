package handlers

import (
	"github.com/codegangsta/martini"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
)

type ApiHandler struct {
	Martini *martini.ClassicMartini
	Protocol protocols.IOProtocol
	Backend backends.Backend
}

func (a *ApiHandler) loadRoutes() {
	a.Martini.Get("/api", func() string {
		return "Hello, world!"
	})

	a.Martini.Get("/api/about", func() string {
		return "About the world!"
	})

	a.Martini.Get("/api/contact", func() string {
		return "Contact the world!"
	})
}

func NewApiHandler(m *martini.ClassicMartini, p protocols.IOProtocol, b backends.Backend) *ApiHandler {
	a := &ApiHandler{}
	a.Backend = b
	a.Protocol = p
	a.Martini = m

	a.loadRoutes()
	return a
}
