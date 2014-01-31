package backends

type Postgres struct{}

func NewPostgres() Backend {
	//Create the object
	p := &Postgres{}

	//Return the object
	return Backend(p)
}

func (p *Postgres) Flag() string {
	return "postgres"
}
