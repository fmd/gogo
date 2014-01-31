package handlers

import (
    "fmt"
    "net/http"
    "github.com/codegangsta/martini"
)

type SiteHandler struct {
    Martini *martini.ClassicMartini
}

func (s *SiteHandler) Route(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "Hello, <a href=\"google.com\">world</a>.")
}

func (s *SiteHandler) loadRoutes() {
    s.Martini.Get("/api", func () string {
        return "Hello, world!"
    })

    s.Martini.Get("/api/about", func () string {
        return "About the world!"
    })

    s.Martini.Get("/api/contact", func () string {
        return "Contact the world!"
    })
}

func NewSiteHandler(m *martini.ClassicMartini) *SiteHandler {
    s := &SiteHandler{}
    s.Martini = m
    
    s.loadRoutes()
    return s
}