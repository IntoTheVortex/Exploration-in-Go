package main

import (
	"fmt"
	"main/graph"
)

func main() {
	fmt.Println("hello")

	graphInt := graph.Empty[int]()
	graphStr := graph.Empty[string]()

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
	//fixed packages, now these are illegal:
	//graphInt.vertices[1] = make(map[int]struct{})
	//graphInt.vertices[1][99] = struct{}{}
	//graphInt.Print()

}
