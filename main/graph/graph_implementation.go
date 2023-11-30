package graph

import "fmt"

//lowercase so not exportable
type set[T comparable] map[T]struct{}

//lowercase so not exportable
type hashGraph[T comparable] struct {
	vertices    map[T]map[T]struct{}
	visited_set set[T]
}


// Set functions *********************************************
func EmptySet[T comparable]() set[T] {
	return make(set[T])
}

func (s set[T]) add(item T) {
	s[item] = struct{}{}
}


// Graph functions *********************************************

//This has to be outside of the class in order to 
// provide the intended functionality
func Empty[T comparable]() *hashGraph[T] {
	return &hashGraph[T]{
		vertices:    make(map[T]map[T]struct{}),
		visited_set: EmptySet[T](),
	}
}

//This one has to be in the class in order to ensure 
// hashGraph implements the graph interface
func (g *hashGraph[T]) Empty() Graph[T] {
	return Empty[T]() //call the standalone version to properly implement interface
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
	set[T](g.vertices[v1]).add(v2)
}

func (g *hashGraph[T]) Neighbors(v1 T, f func(T)) {
	for u := range g.vertices[v1] {
		f(u)
	}
}

//Utility print function
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

//Test Neighbors with dfs
func (g *hashGraph[T]) DFS(v T) {
	_, exists := g.visited_set[v]
	if exists {
		return
	}
	fmt.Printf("%v ", v)
	g.visited_set.add(v)
	g.Neighbors(v, g.DFS)
}

//Only need to demonstrate DFS once, so not necessary
// but it is available if need to search again
func (g *hashGraph[T]) ClearVisited() {
	g.visited_set = EmptySet[T]()
}
