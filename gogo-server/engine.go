package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/handlers"
	"github.com/fmd/gogo/gogo/protocols"
)

type Engine struct {
	Backend     backends.Backend
	Protocol    protocols.IOProtocol
	ApiHandler  *handlers.ApiHandler
	SiteHandler *handlers.SiteHandler
	Martini     *martini.ClassicMartini
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

// ----------------------------
// --- Accessible functions ---
// ----------------------------

// --- Initialisation functions ---
func NewEngine(p string, b string) (*Engine, error) {
	var err error = nil

	//Create the object
	e := &Engine{}
	e.Martini = martini.Classic()

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
	e.SiteHandler = handlers.NewSiteHandler(e.Martini)
	e.ApiHandler = handlers.NewApiHandler(e.Martini, e.Protocol, e.Backend)

	//Return the object
	return e, nil
}

// --- Serving functions ---
func (e *Engine) Run() {
	e.Martini.Run()
}
