package main

import (
	"fmt"
	"main/graph"
)

//This function shows that the Graph interface is 
// implemented by hashGraph. If there is an error here,
// it means that the interface is *not* implemented correctly.
func DemonstrateInterface[T comparable](t T) graph.Graph[T] {
	return graph.Empty[T]()
}

func main() {
	//Instantiate the graphs
	graphInt := graph.Empty[int]()
	graphStr := graph.Empty[string]()

	//Build the two graphs
	c := 'A'
	v_prev := 0
	c_prev := c
	for i := 0; i < 10; i++ {
		v1 := i
		graphInt.AddVertex(v1)
		if i > 0 {
			graphInt.AddEdge(v1, v_prev)
			graphInt.AddEdge(v1, 100)
			graphInt.AddEdge(v1, 200)
			graphInt.AddEdge(v1, 300)
		}
		if i < 5 {
			graphStr.AddVertex(string(c))
			graphStr.AddEdge(string(c), string(c_prev))
		}
		c_prev = c
		c++
		v_prev = v1
	}
	graphInt.AddEdge(0, 9)
	graphInt.AddEdge(100, 100)

	//Print out the graphs
	fmt.Printf("\nInt ")
	graphInt.Print()
	fmt.Printf("\nString ")
	graphStr.Print()

	//Demonstrate DFS on graphInt, starting on 
	// vertex 2
	fmt.Println("\nDepth-First Search on Int Graph:")
	graphInt.DFS(2)

	//Fixed privacy issues, now these are illegal:
	//graphInt.vertices[1] = make(map[int]struct{})
	//graphInt.vertices[1][99] = struct{}{}
	//graphInt.Print()
}
