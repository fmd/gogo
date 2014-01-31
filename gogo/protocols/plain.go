package protocols

type Plain struct{}

func NewPlain() IOProtocol {
	//Create the object
	p := &Plain{}

	//Return the object
	return IOProtocol(p)
}

func (p *Plain) Flag() string {
	return "plain"
}
