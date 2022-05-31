package generator

import "sort"

type Cell struct {
	Site      Vertex
	Halfedges []*Halfedge
}

func newCell(site Vertex) *Cell {
	return &Cell{Site: site}
}

func (t *Cell) prepare() int {
	halfedges := t.Halfedges
	iHalfedge := len(halfedges) - 1

	for ; iHalfedge >= 0; iHalfedge-- {
		edge := halfedges[iHalfedge].Edge

		if edge.VertexB.Vertex == NULL_VERTEX || edge.VertexA.Vertex == NULL_VERTEX {
			halfedges[iHalfedge] = halfedges[len(halfedges)-1]
			halfedges = halfedges[:len(halfedges)-1]
		}
	}

	sort.Sort(halfEdgesByAngle{halfedges})
	t.Halfedges = halfedges
	return len(halfedges)
}

func (cell *Cell) GetArea() float64 {
	area := float64(0)
	for _, halfedge := range cell.Halfedges {
		s := halfedge.GetStartpoint()
		e := halfedge.GetEndpoint()
		area += s.X * e.Y
		area -= s.Y * e.X
	}

	return area / 2
}

func (cell *Cell) CellCentroid() Vertex {
	x, y := float64(0), float64(0)
	for _, halfedge := range cell.Halfedges {
		s := halfedge.GetStartpoint()
		e := halfedge.GetEndpoint()
		v := s.X*e.Y - e.X*s.Y
		x += (s.X + e.X) * v
		y += (s.Y + e.Y) * v
	}
	v := cell.GetArea() * 6
	return Vertex{x / v, y / v}
}

func (cell *Cell) InsideCell(v Vertex) bool {
	for _, halfedge := range cell.Halfedges {
		a := halfedge.GetStartpoint()
		b := halfedge.GetEndpoint()

		cross := ((b.X-a.X)*(v.Y-a.Y) - (b.Y-a.Y)*(v.X-a.X))

		if cross > 0 {
			return false
		}
	}
	return true
}

func (cell *Cell) EdgeIndex(edge *Edge) int {
	for i, halfedge := range cell.Halfedges {
		if halfedge.Edge == edge {
			return i
		}
	}
	return -1
}

func LloydRelaxation(cells []*Cell) (result []Vertex) {
	result = make([]Vertex, len(cells))
	for id, cell := range cells {
		result[id] = cell.CellCentroid()
	}
	return
}
