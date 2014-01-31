package backends

type FlatFile struct {}

func NewFlatFile() Backend {
    //Create the object
    p := &FlatFile{}

    //Return the object
    return Backend(p)
}

func (p *FlatFile) Flag() string {
    return "flatfile"
}