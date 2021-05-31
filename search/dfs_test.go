package search

import (
	"testing"

	"github.com/alright21/algo/graph"
)


func TestDFS(t* testing.T){

	// not a proper test, but it is used to check the functionality of the algorithm
	testGraph := &graph.Graph{Visited: make(map[int]bool), Adj: make(map[int][]int)}

	testGraph.AddEdge(0,1)
	testGraph.AddEdge(0,2)
	testGraph.AddEdge(1,3)
	testGraph.AddEdge(3,2)

	DFS(testGraph,0)
}