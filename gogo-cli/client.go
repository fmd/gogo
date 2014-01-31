package main

import (
	"flag"
	"fmt"
	"github.com/fmd/gogo/gogo"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
	"strings"
)

var (
	proto *string
	backend *string
	verbose *bool
)

type Client struct {
	Engine *gogo.Engine
	Verbose bool
}

func (c *Client) parseVerbose(verbose bool) {
	c.Verbose = verbose
	if verbose {
		fmt.Println("Using verbose mode.")
	}
}

func (c *Client) Init() error {
	
	//Parse the flags to get the engine setup
	c.ParseFlags()
	c.parseVerbose(*verbose)

	//Use our proto and backend variables to load the engine.
	var err error
	c.Engine, err = gogo.NewEngine(*proto, *backend)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ParseFlags() {

	//Get protocols for use in the help text and default protocol
	pdFlag := protocols.GetProtocolFlags()[0]
	pFlags := strings.Join(protocols.GetProtocolFlags(), "', '")
	pMsg := fmt.Sprintf("I/O Protocol to use ('%s')", pFlags)

	//Get backends for use in the help text and default backend
	bdFlag := backends.GetBackendFlags()[0]
	bFlags := strings.Join(backends.GetBackendFlags(), "', '")
	bMsg := fmt.Sprintf("Storage backend to use ('%s')", bFlags)

	//Create and parse the flags
	proto = flag.String("p", pdFlag, pMsg)
	backend = flag.String("s", bdFlag, bMsg)
	verbose = flag.Bool("v", false, "Verbose mode")
	flag.Parse()
}
