package handlers

import (
	"github.com/codegangsta/martini"
)

type SiteHandler struct {
	Martini *martini.ClassicMartini
}

func (s *SiteHandler) loadRoutes() {
	s.Martini.Get("/", func() string {
		return "Hello, world."
	})

	s.Martini.Get("/about", func() string {
		return "About the world."
	})

	s.Martini.Get("/contact", func() string {
		return "Contact the world."
	})
}

func NewSiteHandler(m *martini.ClassicMartini) *SiteHandler {
	s := &SiteHandler{}
	s.Martini = m

	s.loadRoutes()
	return s
}
