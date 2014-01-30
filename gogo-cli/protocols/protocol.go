package protocols

import (
    "fmt"
    "errors"
)

var (
    //To add a new protocol, implement IOProtocol and add to this slice.
    Protocols = []IOProtocol{ NewPlain(), NewGtp1() }
)

type IOProtocol interface {
    Flag() string
}

func GetProtocol(name string) (*IOProtocol, error) {
    for p := range Protocols {
        proto := Protocols[p]
        if proto.Flag() == name {
            return &proto, nil
        }
    }
    msg := fmt.Sprintf("protocols: Could not find protocol '%s'.", name)
    err := errors.New(msg)
    return nil, err
}