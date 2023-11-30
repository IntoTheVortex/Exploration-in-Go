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

	buildIntGraph()
	buildStrGraph()

}

func buildIntGraph() {
	//Instantiate the graph
	graphInt := graph.Empty[int]()

	//Build the graph
	v_prev := 0
	for i := 0; i < 10; i++ {
		v1 := i
		graphInt.AddVertex(v1)
		if i > 0 {
			graphInt.AddEdge(v1, v_prev)
			graphInt.AddEdge(v1, 100)
			graphInt.AddEdge(v1, 200)
			graphInt.AddEdge(v1, 300)
		}
		v_prev = v1
	}

	//Print out the current graph
	fmt.Printf("Tree version of Int ")
	graphInt.Print()

	//Add some more edges so it's not a tree
	graphInt.AddEdge(0, 9) 
	graphInt.AddEdge(6, 0) 
	graphInt.AddEdge(100, 200)

	//Print out the graph
	fmt.Println("\nFull Int Graph")
	graphInt.Print()

	//Demonstrate DFS on graphInt, starting on 
	// vertex 2
	fmt.Println("\nDepth-First Search on Int Graph, vertex 2:")
	graphInt.DFS(2)
	fmt.Println()

	//Fixed privacy issues, now these are illegal:
	//graphInt.vertices[1] = make(map[int]struct{})
	//graphInt.vertices[1][99] = struct{}{}
	//graphInt.Print() //just to show the result
}

func buildStrGraph() {
	graphStr := graph.Empty[string]()

	//Build the graph
	c := 'A'
	c_prev := c
	for i := 0; i < 10; i++ {
		graphStr.AddVertex(string(c))
		if i > 0 {
			graphStr.AddEdge(string(c), string(c_prev))
			graphStr.AddEdge(string(c), "Hello")
		}
		if (i % 2 == 0) {
			graphStr.AddEdge(string(c), "Goodbye")
			graphStr.AddEdge(string(c), "So long")
			graphStr.AddEdge(string(c), "Farewell")
		} else {
			graphStr.AddEdge(string(c), "Greetings")
			graphStr.AddEdge(string(c), "Welcome")
			graphStr.AddEdge(string(c), "Hi")
		}
		c_prev = c
		c++
	}

	//Print out the current graph
	fmt.Printf("\nTree version of String ")
	graphStr.Print()

	//Add some more edges so it's not just a tree
	graphStr.AddEdge("A", "B")
	graphStr.AddEdge("C", "A")
	graphStr.AddEdge("A", "C")

	//Print out the graph
	fmt.Printf("\nFull String ")
	graphStr.Print()

	//Demonstrate DFS on graphStr, starting on 
	// vertex 2
	fmt.Println("\nDepth-First Search on Str Graph:")
	graphStr.DFS("B")

}