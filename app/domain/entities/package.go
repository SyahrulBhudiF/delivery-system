package entities

type Package struct {
	ID       int
	Weight   float64
	Contents string
}

func NewPackage(id int, weight float64, contents string) *Package {
	return &Package{
		ID:       id,
		Weight:   weight,
		Contents: contents,
	}
}
