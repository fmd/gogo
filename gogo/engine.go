package gogo

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
)

type Engine struct {
	Backend  backends.Backend
	Protocol protocols.IOProtocol
	ApiHandler handlers.ApiHandler
	SiteHandler handlers.SiteHandler
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
func siteHandler(w http.ResponseWriter, r *http.Request) {
	//Pass request to the siteHandler. Generally we'll be returning static files.
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	//1. Decode request POST/GET data using e.Protocol to raw.
	//2. Pass request to ApiHandler with the backend to apply the changes to. Get raw response.
	//3. Encode response data using e.Protocol.
}

// ----------------------------
// --- Accessible functions ---
// ----------------------------

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

	//Return the object
	return e, nil
}

func (e *Engine) Run() {
	r := mux.NewRouter()
    r.HandleFunc("/", siteHandler)
    r.HandleFunc("/game", siteHandler)
    r.HandleFunc("/api", apiHandler)
    http.Handle("/", r)

    log.Fatal(http.ListenAndServe(":3000", nil))
}