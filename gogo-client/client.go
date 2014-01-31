package main

import (
	"flag"
	"fmt"
	curl "github.com/andelf/go-curl"
	"github.com/fmd/gogo/gogo/protocols"
	"strings"
)

var (
	proto *string
)

type Client struct{}

func (c *Client) ParseFlags() {

	//Get protocols for use in the help text and default protocol
	pdFlag := protocols.GetProtocolFlags()[0]
	pFlags := strings.Join(protocols.GetProtocolFlags(), "', '")
	pMsg := fmt.Sprintf("I/O Protocol to use ('%s')", pFlags)

	//Create and parse the flags
	proto = flag.String("p", pdFlag, pMsg)
	flag.Parse()
}

func (c *Client) Init() {
	//Parse the flags
	c.ParseFlags()
}

func (c *Client) Run() {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "localhost:3000")

	// make a callback function
	fooTest := func(buf []byte, userdata interface{}) bool {
		println("DEBUG: size=>", len(buf))
		println("DEBUG: content=>", string(buf))
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)

	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
