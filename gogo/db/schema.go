package db

import (
	"github.com/eaigner/hood"
)

type User struct {
	Id       hood.Id
	Email    string
	Username string
	Password string
	Created  hood.Created
	Updated  hood.Updated
}
