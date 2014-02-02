package handlers

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

type SiteHandler struct {
	Martini *martini.ClassicMartini
}

func (s *SiteHandler) loadRoutes() {
	s.Martini.Get("/", func(r render.Render)  {
		r.HTML(200, "home", "world")
	})

	s.Martini.Get("/about", func(r render.Render)  {
		r.HTML(200, "about", "world")
	})

	s.Martini.Get("/contact", func(r render.Render)  {
		r.HTML(200, "contact", "world")
	})
}

func NewSiteHandler(m *martini.ClassicMartini) *SiteHandler {
	s := &SiteHandler{}
	s.Martini = m

	s.loadRoutes()
	return s
}
