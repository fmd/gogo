package main

import (
    "github.com/eaigner/hood"
)

type User struct {
    Id hood.Id
    Email string
    Username string

    Created hood.Created
    Updated hood.Updated
  }

func (m *M) CreateUserTable_1391299892_Up(hd *hood.Hood) {
    hd.CreateTable(&User{})
}

func (m *M) CreateUserTable_1391299892_Down(hd *hood.Hood) {
    hd.DropTable(&User{})
}   