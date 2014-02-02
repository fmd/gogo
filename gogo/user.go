package main

import (
    "github.com/eaigner/hood"
)

type User struct {
    Id hood.Id `sql:"id" json:"id"`
    Email string `sql:"email" json:"email"`
    Username string `sql:"username" json:"username"`
    Password string `sql:"password"`
    //authenticated bool `json:"-"`

    Created hood.Created
    Updated hood.Updated
}
/*
func (u *User) IsAuthenticated() bool {
    return u.authenticated
}

func (u *User) Login() {
    u.authenticated = true
}

func (u *User) Logout() {
    u.authenticated = false
}
*/
func (u *User) UniqueId() interface{} {
    return u.Id
}

func (u *User) GetById(id interface{}) error {
    u.Id = id.(hood.Id)
    return nil
}