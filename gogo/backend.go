package main

import (
    "github.com/eaigner/hood"
)

type Backend struct {
    Hd *hood.Hood
}

func NewBackend() (*Backend, error) {
    b := &Backend{}
    h, err := hood.Open("postgres", "user='fareed' dbname='gogo' sslmode='disable' password='Killjoy12345'")
    b.Hd = h
    if err != nil {
        return nil, err
    }
    return b, nil
}

func (b *Backend) Save(u *User) error {
    tx := b.Hd.Begin()

    id, err := tx.Save(u)
    if err != nil {
        return err
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    u.Id = hood.Id(id)
    return nil
}