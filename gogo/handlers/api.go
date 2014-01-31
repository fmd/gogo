package handlers

import (
    "fmt"
    "net/http"
    "github.com/codegangsta/martini"
    "github.com/fmd/gogo/gogo/backends"
)

type ApiHandler struct {
    Martini *martini.ClassicMartini
    Backend backends.Backend
}

func (a *ApiHandler) loadRoutes() {
    a.Martini.Get("/", func () string {
        return "Hello, world."
    })

    a.Martini.Get("/about", func () string {
        return "About the world."
    })

    a.Martini.Get("/contact", func () string {
        return "Contact the world."
    })
}

func (a *ApiHandler) Route(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, "{\"hello\": \"world\"}")
}

func NewApiHandler(m *martini.ClassicMartini, b backends.Backend) *ApiHandler {
    a := &ApiHandler{}
    a.Backend = b
    a.Martini = m
    
    a.loadRoutes()
    return a
}