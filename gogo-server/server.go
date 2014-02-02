package main

import (
	"fmt"
	"flag"
	"strings"
	"github.com/fmd/gogo/gogo/protocols"
)

type Server struct {
	Engine  *Engine
	Verbose bool
}

// --------------------------
// --- Internal functions ---
// --------------------------

func (s *Server) parseProtocolFlag() *string {
	//Get protocols for use in the help text and default protocol
	pdFlag := protocols.GetProtocolFlags()[0]
	pFlags := strings.Join(protocols.GetProtocolFlags(), "', '")
	pMsg := fmt.Sprintf("I/O Protocol to use ('%s')", pFlags)

	return flag.String("p", pdFlag, pMsg)
}

func (s *Server) parseVerboseFlag() *bool {
	return flag.Bool("v", false, "Verbose mode")
}

// ----------------------------
// --- Accessible functions ---
// ----------------------------

func (s *Server) Init() error {

	//Parse the flags
	proto := s.parseProtocolFlag()
	s.Verbose = *s.parseVerboseFlag()
	flag.Parse()

	//Use our proto and backend variables to load the engine.
	var err error
	s.Engine, err = NewEngine(*proto)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) ServeForever() {
	s.Engine.Run()
}
