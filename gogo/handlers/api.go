package handlers

import (
    "fmt"
    "net/http"
    "github.com/fmd/gogo/gogo/backends"
)

type ApiHandler struct {
    Backend backends.Backend
}

func (a *ApiHandler) Route(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, "Hello, <a href=\"google.com\">world</a>.")
}

func NewApiHandler(b backends.Backend) *ApiHandler {
    a := &ApiHandler{}
    a.Backend = b
    return a
}