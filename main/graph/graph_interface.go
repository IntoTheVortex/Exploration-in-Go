package graph

type Graph[V comparable] interface {
	Empty() Graph[V]
	AddVertex(V)
	AddEdge(V, V)
	Neighbors(func(V, V) V, V, V) bool
}
