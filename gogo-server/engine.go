package main

import (
	"fmt"
	"github.com/eaigner/hood"
	"github.com/codegangsta/martini"
	"github.com/fmd/gogo/gogo/protocols"
	"github.com/fmd/gogo/gogo-server/handlers"
	"github.com/codegangsta/martini-contrib/render"
)

type Engine struct {
	Backend     *Backend
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

// ----------------------------
// --- Accessible functions ---
// ----------------------------

// --- Initialisation functions ---
func NewEngine(protocol string) (*Engine, error) {
	var err error = nil

	//Create the object
	e := &Engine{}
	e.Martini = martini.Classic()
	e.Martini.Use(render.Renderer())

	//Set the protocol
	err = e.useProtocol(protocol)
	if err != nil {
		return nil, err
	}

	e.Backend = NewBackend()
	e.SiteHandler = handlers.NewSiteHandler(e)
	e.ApiHandler = handlers.NewApiHandler(e)

	//Return the object
	return e, nil
}

// --- Serving functions ---
func (e *Engine) Run() {
	e.Martini.Run()
}
