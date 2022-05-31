package generator

import "math"

// Vertex
type Vertex struct {
	X float64
	Y float64
}

var NULL_VERTEX = Vertex{math.Inf(1), math.Inf(1)}

// Sort interface
type Vertices []Vertex

func (s Vertices) Len() int {
	return len(s)
}

func (s Vertices) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type VerticesByY struct{ Vertices }

func (s VerticesByY) Less(i, j int) bool {
	return s.Vertices[i].Y < s.Vertices[j].Y
}

func Distance(a, b Vertex) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2.0) + math.Pow(a.Y-b.Y, 2.0))
}

// EdgeVertex
type EdgeVertex struct {
	Vertex
	Edges []*Edge
}

var NULL_EDGE_VERTEX = EdgeVertex{NULL_VERTEX, nil}
