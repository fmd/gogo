package gogo

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fmd/gogo/gogo/handlers"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
)

type Engine struct {
	Backend  backends.Backend
	Protocol protocols.IOProtocol
	ApiHandler *handlers.ApiHandler
	SiteHandler *handlers.SiteHandler
}

// --------------------------
// --- Internal functions ---
// --------------------------

// --- Initialisation functions ---
func (e *Engine) useProtocol(name string) error {
	p, err := protocols.GetProtocol(name)
	if err != nil {
		return err
	}

	e.Protocol = p
	fmt.Println(fmt.Sprintf("Using protocol '%s'.", e.Protocol.Flag()))
	return nil
}

func (e *Engine) useBackend(name string) error {
	b, err := backends.GetBackend(name)
	if err != nil {
		return err
	}

	e.Backend = b
	fmt.Println(fmt.Sprintf("Using storage format '%s'.", e.Backend.Flag()))
	return nil
}

// --- Routing functions ---
func (e *Engine) siteHandler(w http.ResponseWriter, r *http.Request) {
	e.SiteHandler.Route(w, r)
	//Pass request to the siteHandler. Generally we'll be returning static files.
}

func (e *Engine) apiHandler(w http.ResponseWriter, r *http.Request) {
	e.ApiHandler.Route(w, r)
	//1. Decode request POST/GET data using e.Protocol to raw.
	//2. Pass request to ApiHandler.
	//3. Encode response data using e.Protocol.
}

// ----------------------------
// --- Accessible functions ---
// ----------------------------

// --- Initialisation functions ---
func NewEngine(p string, b string) (*Engine, error) {
	var err error = nil

	//Create the object
	e := &Engine{}

	//Set the protocol
	err = e.useProtocol(p)
	if err != nil {
		return nil, err
	}

	//Set the backend
	err = e.useBackend(b)
	if err != nil {
		return nil, err
	}

	//Create the SiteHandler
	e.SiteHandler = handlers.NewSiteHandler()
	e.ApiHandler = handlers.NewApiHandler(e.Backend)

	//Return the object
	return e, nil
}

// --- Serving functions ---
func (e *Engine) Run() {
	r := mux.NewRouter()
    r.HandleFunc("/", e.siteHandler)
    r.HandleFunc("/game", e.siteHandler)
    r.HandleFunc("/api", e.apiHandler)
    http.Handle("/", r)

    log.Fatal(http.ListenAndServe(":3000", nil))
}