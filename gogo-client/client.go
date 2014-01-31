package main

import (
	"flag"
	"fmt"
	"github.com/fmd/gogo/gogo/protocols"
	"strings"
)

var (
	proto *string
)

type Client struct{}

func (c *Client) Init() {
	//Parse the flags
	c.ParseFlags()
}

func (c *Client) ParseFlags() {

	//Get protocols for use in the help text and default protocol
	pdFlag := protocols.GetProtocolFlags()[0]
	pFlags := strings.Join(protocols.GetProtocolFlags(), "', '")
	pMsg := fmt.Sprintf("I/O Protocol to use ('%s')", pFlags)

	//Create and parse the flags
	proto = flag.String("p", pdFlag, pMsg)
	flag.Parse()
}
