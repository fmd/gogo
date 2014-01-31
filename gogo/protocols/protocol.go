package protocols

import (
    "fmt"
    "errors"
)

var (
    //To add a new protocol, implement IOProtocolFlag and add its New() to this slice
    Protocols = []IOProtocol{ NewPlain(), NewGtp1() }
)

type IOProtocol interface {
    Flag() string
}

func GetProtocolFlags() []string {
    flags := []string{}
    for p := range Protocols {
        f := Protocols[p]
        flags = append(flags, f.Flag())
    }

    return flags
}

func GetProtocol(name string) (IOProtocol, error) {
    for p := range Protocols {
        f := Protocols[p]
        if f.Flag() == name {
            return f, nil
        }
    }
    msg := fmt.Sprintf("protocols: Could not find protocol '%s'.", name)
    err := errors.New(msg)
    return nil, err
}