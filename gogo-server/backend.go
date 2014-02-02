package main

import (
    "github.com/eaigner/hood"
)

type Backend struct {
    Hood *hood.Hood
}

func NewBackend() (*Backend, error) {
    b := &Backend{}
    b.Hood, err := hood.Open("postgres", "user='fareed' dbname='gogo' sslmode='disable' password='Killjoy12345'")
    if err != nil {
        return nil, err
    }
    return b, nil
}

func (b *Backend) Save(o interface{}) error {
    tx := hd.Begin()

    _, err := tx.Save(o)
    if err != nil {
        return err
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}


