package main

import (
	"github.com/eaigner/hood"
)

func (m *M) CreateUserTable_1391299892_Up(hd *hood.Hood) {
  type Users struct {
    Id hood.Id
    Email string
    Username string
  }
  hd.CreateTable(&Users{})
}

func (m *M) CreateUserTable_1391299892_Down(hd *hood.Hood) {
	// TODO: implement
}