package handlers

import (
    "fmt"
    "net/http"
)

type SiteHandler struct {}

func (s *SiteHandler) Route(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "Hello, <a href=\"google.com\">world</a>.")
}

func NewSiteHandler() *SiteHandler {
    return &SiteHandler{}
}