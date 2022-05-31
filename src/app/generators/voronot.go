package generator

// Computing of voronoi diagramms is implemented via Fortune algorithm.
// TL;DR
// Naive algorithm - O(n^4);
// Fortune algorithm - O(nlogn)
// In context of generating ~ 10000 maps, difference is brutal

type Voronoi struct {
	cells []*Cell
	edges []*Edge
}
