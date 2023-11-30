package graph

type Graph[V comparable] interface {
	Empty() Graph[V]
	AddVertex(V)
	AddEdge(V, V)
	Neighbors(V, func(V))
	DFS(V)
}
