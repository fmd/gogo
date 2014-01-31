package gogo

import (
    "fmt"
    "github.com/fmd/gogo/gogo/backends"
    "github.com/fmd/gogo/gogo/protocols"
)

type Engine struct {
    Backend  backends.Backend
    Protocol protocols.IOProtocol
}

func NewEngine(p string, b string) (*Engine, error) {
    var err error = nil

    //Create the object
    e := &Engine{}

    //Set the protocol
    err = e.UseProtocol(p)
    if err != nil {
        return nil, err
    }

    //Set the backend
    err = e.UseBackend(b)
    if err != nil {
        return nil, err
    }

    //Return the object
    return e, nil
}

func (e *Engine) UseProtocol(name string) error {
    p, err := protocols.GetProtocol(name)
    if err != nil {
        return err
    }

    e.Protocol = p
    fmt.Println(fmt.Sprintf("Using protocol '%s'.", e.Protocol.Flag()))
    return nil
}

func (e *Engine) UseBackend(name string) error {
    b, err := backends.GetBackend(name)
    if err != nil {
        return err
    }

    e.Backend = b
    fmt.Println(fmt.Sprintf("Using storage format '%s'.", e.Backend.Flag()))
    return nil
}