package main

import (
	"fmt"
	"main/graph"
)

func main() {
	graphInt := graph.Empty[int]()
	graphStr := graph.Empty[string]()

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

	fmt.Printf("\nInt ")
	graphInt.Print()
	fmt.Printf("\nString ")
	graphStr.Print()
	fmt.Println("\nDepth-First Search:")
	graphInt.DFS(2)

	//fixed packages, now these are illegal:
	//graphInt.vertices[1] = make(map[int]struct{})
	//graphInt.vertices[1][99] = struct{}{}
	//graphInt.Print()

}
