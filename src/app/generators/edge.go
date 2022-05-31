package generator

import "math"

// Edge
type Edge struct {
	LeftCell  *Cell
	RightCell *Cell
	VertexA   EdgeVertex
	VertexB   EdgeVertex
}

func (e *Edge) GetOtherCell(cell *Cell) *Cell {
	if cell == e.LeftCell {
		return e.RightCell
	} else if cell == e.RightCell {
		return e.LeftCell
	}
	return nil
}

func (e *Edge) GetOtherEdgeVertex(v Vertex) EdgeVertex {
	if v == e.VertexA.Vertex {
		return e.VertexB
	} else if v == e.VertexB.Vertex {
		return e.VertexA
	}
	return NULL_EDGE_VERTEX
}

func newEdge(LeftCell, RightCell *Cell) *Edge {
	return &Edge{
		LeftCell:  LeftCell,
		RightCell: RightCell,
		VertexA:   NULL_EDGE_VERTEX,
		VertexB:   NULL_EDGE_VERTEX,
	}
}

//Halfedge
type Halfedge struct {
	Cell  *Cell
	Edge  *Edge
	Angle float64
}

// Sort interface
type HalfEdges []*Halfedge

func (s HalfEdges) Len() int {
	return len(s)
}

func (s HalfEdges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type halfEdgesByAngle struct{ HalfEdges }

func (s halfEdgesByAngle) Less(i, j int) bool {
	return s.HalfEdges[i].Angle > s.HalfEdges[j].Angle
}

func newHalfEdge(edge *Edge, LeftCell, RightCell *Cell) *Halfedge {
	result := &Halfedge{
		Cell: LeftCell,
		Edge: edge,
	}

	if RightCell != nil {
		result.Angle = math.Atan2(RightCell.Site.Y-LeftCell.Site.Y, RightCell.Site.X-LeftCell.Site.X)
	} else {
		vertexA := edge.VertexA
		vertexB := edge.VertexB

		if edge.LeftCell == LeftCell {
			result.Angle = math.Atan2(vertexB.X-vertexA.X, vertexA.Y-vertexB.Y)
		} else {
			result.Angle = math.Atan2(vertexA.X-vertexB.X, vertexB.Y-vertexA.Y)
		}
	}
	return result
}

func (h *Halfedge) GetStartpoint() Vertex {
	if h.Edge.LeftCell == h.Cell {
		return h.Edge.VertexA.Vertex
	}
	return h.Edge.VertexB.Vertex

}

func (h *Halfedge) GetEndpoint() Vertex {
	if h.Edge.LeftCell == h.Cell {
		return h.Edge.VertexB.Vertex
	}
	return h.Edge.VertexA.Vertex
}
