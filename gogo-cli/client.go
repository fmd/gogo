package main

import (
	"flag"
	"fmt"
	"github.com/fmd/gogo/gogo/backends"
	"github.com/fmd/gogo/gogo/protocols"
	"strings"
)

type Client struct {
	Backend  backends.Backend
	Protocol protocols.IOProtocol
	Verbose  bool
}

func (c *Client) parseProtocol(name string) {
	p, err := protocols.GetProtocol(name)
	if err != nil {
		panic(err)
	}

	c.Protocol = p
	fmt.Println(fmt.Sprintf("Using protocol '%s'.", c.Protocol.Flag()))
}

func (c *Client) parseBackend(name string) {
	b, err := backends.GetBackend(name)
	if err != nil {
		panic(err)
	}

	c.Backend = b
	fmt.Println(fmt.Sprintf("Using storage format '%s'.", c.Backend.Flag()))
}

func (c *Client) parseVerbose(verbose bool) {
	c.Verbose = verbose
	if verbose {
		fmt.Println("Using verbose mode.")
	}
}

func (c *Client) Init() {
	c.ParseFlags()
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
	proto := flag.String("p", pdFlag, pMsg)
	backend := flag.String("s", bdFlag, bMsg)
	verbose := flag.Bool("v", false, "Verbose mode")
	flag.Parse()

	//Load the flags into this struct.
	c.parseProtocol(*proto)
	c.parseBackend(*backend)
	c.parseVerbose(*verbose)
}
