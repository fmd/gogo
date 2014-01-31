package main

import (
	"flag"
	"fmt"
	"github.com/fmd/gogo/gogo"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
	"strings"
)

type Server struct {
	Engine  *gogo.Engine
	Verbose bool
}

func (s *Server) Init() error {

	//Parse the flags
	proto := s.ParseProtocol()
	backend := s.ParseBackend()
	s.Verbose = *s.parseVerbose()
	flag.Parse()

	//Use our proto and backend variables to load the engine.
	var err error
	s.Engine, err = gogo.NewEngine(*proto, *backend)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) ParseProtocol() *string {
	//Get protocols for use in the help text and default protocol
	pdFlag := protocols.GetProtocolFlags()[0]
	pFlags := strings.Join(protocols.GetProtocolFlags(), "', '")
	pMsg := fmt.Sprintf("I/O Protocol to use ('%s')", pFlags)

	return flag.String("p", pdFlag, pMsg)
}

func (s *Server) ParseBackend() *string {
	//Get backends for use in the help text and default backend
	bdFlag := backends.GetBackendFlags()[0]
	bFlags := strings.Join(backends.GetBackendFlags(), "', '")
	bMsg := fmt.Sprintf("Storage backend to use ('%s')", bFlags)

	return flag.String("s", bdFlag, bMsg)
}

func (s *Server) parseVerbose() *bool {
	return flag.Bool("v", false, "Verbose mode")
}
