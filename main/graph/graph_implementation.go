package graph

import "fmt"

//lowercase so not exportable
type set[T comparable] map[T]struct{}

//lowercase so not exportable
type hashGraph[T comparable] struct {
	vertices map[T]map[T]struct{}
}

// Set functions *********************************************
func EmptySet[T comparable]() set[T] {
	return make(set[T])
}

func (s set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s set[T]) Remove(item T) {
	delete(s, item)
}

func (s set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Graph functions *********************************************
func Empty[T comparable]() *hashGraph[T] {
	return &hashGraph[T]{
		vertices: make(map[T]map[T]struct{}),
	}
}

func (g *hashGraph[T]) AddVertex(v T) {
	_, exists := g.vertices[v]
	if exists {
		return
	} else {
		g.vertices[v] = EmptySet[T]()
	}
}

func (g *hashGraph[T]) AddEdge(v1 T, v2 T) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	set[T](g.vertices[v1]).Add(v2)
}

func (g *hashGraph[T]) Neighbors(v1 T, f func(T)) {
	for u := range g.vertices[v1] {
		f(u)
	}
}

func (g *hashGraph[T]) Print() {
	fmt.Println("Graph:")
	for v, neighbors := range g.vertices {
		fmt.Printf("%v: ", v)
		for neighbor := range neighbors {
			fmt.Printf("%v ", neighbor)
		}
		fmt.Println()
	}
}

//test Neighbors with dfs
func (g *hashGraph[T]) DFS(v T) {
	fmt.Printf("%v ", v)
	g.Neighbors(v, g.DFS)
}