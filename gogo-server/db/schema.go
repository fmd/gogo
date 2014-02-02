package db

import (
	"github.com/eaigner/hood"
)

type Users struct {
	Id       hood.Id
	Email    string
	Username string
}
