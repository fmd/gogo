package backends

import (
	"errors"
	"fmt"
)

var (
	//To add a new protocol, implement Backend and add its New() to this slice
	Backends = []Backend{NewFlatFile(), NewPostgres()}
)

type Backend interface {
	Flag() string
}

func GetBackendFlags() []string {
	flags := []string{}
	for p := range Backends {
		f := Backends[p]
		flags = append(flags, f.Flag())
	}

	return flags
}

func GetBackend(name string) (Backend, error) {
	for p := range Backends {
		f := Backends[p]
		if f.Flag() == name {
			return f, nil
		}
	}
	msg := fmt.Sprintf("backends: Could not find backend '%s'.", name)
	err := errors.New(msg)
	return nil, err
}
