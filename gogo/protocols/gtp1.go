package protocols

type Gtp1 struct{}

func NewGtp1() IOProtocol {
	//Create the object
	g := &Gtp1{}

	//Return the object
	return IOProtocol(g)
}

func (g *Gtp1) Flag() string {
	return "gtp1"
}
