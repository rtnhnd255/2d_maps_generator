package generator

type Map struct {
	Cells []*Cell
	Edges []*Edge
}

func field2map(field DSField) Map {
	return Map{}
}
