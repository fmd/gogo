package main

import (
    "fmt"
    "flag"
    "strings"
    "github.com/fmd/gogo/gogo-cli/protocols"
)

type Client struct {
    Protocol protocols.IOProtocol
    Verbose bool
}

func (c *Client) parseProtocol(name string) {
    p, err := protocols.GetProtocol(name)
    if err != nil {
        panic(err)
    }

    c.Protocol = p
    fmt.Println(fmt.Sprintf("Using protocol: '%s'.",c.Protocol.Flag()))
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
    
    //Protocol flag
    firstFlag := protocols.GetProtocolFlags()[0]
    pFlags := strings.Join(protocols.GetProtocolFlags(), ", ")
    pMsg := fmt.Sprintf("I/O Protocol to use (%s)",pFlags)
    var proto = flag.String("p", firstFlag, pMsg)
    
    //Verbose flag
    var verbose = flag.Bool("v", false, "Verbose mode")

    //Parse the flags
    flag.Parse()

    c.parseProtocol(*proto)
    c.parseVerbose(*verbose)
}