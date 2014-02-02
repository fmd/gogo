package backends

import (
    "github.com/fmd/gogo/gogo-server/models"
    "github.com/eaigner/hood"
)

type Postgres struct {
    Hood *hood.Hood
}

func NewPostgres() Backend {
	//Create the object
	p := &Postgres{}
    p.Hood, err := hood.Open("postgres", "user='fareed' dbname='gogo' sslmode='disable' password='Killjoy12345'")
    if err != nil {
        panic(err)
    }

	//Return the object
	return Backend(p)
}

func (p *Postgres) Flag() string {
	return "postgres"
}

func (p *Postgres) CreateUser()