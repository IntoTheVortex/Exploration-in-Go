package main

import "fmt"

type Graph[V comparable] interface {
	Empty() Graph[V]
	AddVertex(V)
	AddEdge(V, V)
	Neighbors(func(V, V) V, V, V) bool
}

type Set[T comparable] map[T]struct{}

//type Vertex[T comparable] struct {
//Data T
//}

type HashGraph[T comparable] struct {
	vertices map[T]map[T]struct{}
}

// Set functions *********************************************
func EmptySet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Graph functions *********************************************
func Empty[T comparable]() *HashGraph[T] {
	return &HashGraph[T]{
		vertices: make(map[T]map[T]struct{}),
	}
}

func (g *HashGraph[T]) AddVertex(v T) error {
	_, exists := g.vertices[v]
	if exists {
		return fmt.Errorf("vertex already in graph")
	} else {
		g.vertices[v] = EmptySet[T]()
	}
	return nil
}

func (g *HashGraph[T]) AddEdge(v1 T, v2 T) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	_set := Set[T](g.vertices[v1])
	_set.Add(v2)
	g.vertices[v1] = _set
}

func (g *HashGraph[T]) Neighbors(v1 T, f func(T)) {
	for u := range g.vertices[v1] {
		f(u)
	}
}

func (g *HashGraph[T]) Print() {
	fmt.Println("Graph:")
	for v, neighbors := range g.vertices {
		fmt.Printf("%v: ", v)
		for neighbor := range neighbors {
			fmt.Printf("%v ", neighbor)
		}
		fmt.Println()
	}
}

//TODO fix
//test Neighbors with dfs
func (g *HashGraph[T]) DFS(v T) {
	fmt.Printf("%v ", v)
	g.Neighbors(v, g.DFS)
}

func main() {
	fmt.Println("hello")

	graphInt := Empty[int]()
	graphStr := Empty[string]()

	c := 'A'
	v_prev := 0
	for i := 0; i < 10; i++ {
		v1 := i
		v2 := string(c)
		c++
		graphInt.AddVertex(v1)
		if i > 0 {
			graphInt.AddEdge(v1, v_prev)
			graphInt.AddEdge(v1, 100)
			graphInt.AddEdge(v1, 200)
			graphInt.AddEdge(v1, 300)
		}
		graphStr.AddVertex(v2)
		v_prev = v1
	}
	graphInt.Print()
	graphStr.Print()
	graphInt.DFS(2)

}
