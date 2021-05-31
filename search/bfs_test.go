package search

import (
	"testing"

	"github.com/alright21/algo/structs"
)

func TestBFS(t* testing.T){

	// not a proper test, but it is used to check the functionality of the algorithm
	testGraph := &structs.Graph{Visited: make(map[int]bool), Adj: make(map[int][]int)}

	testGraph.AddEdge(0,1)
	testGraph.AddEdge(0,2)
	testGraph.AddEdge(1,3)
	testGraph.AddEdge(3,2)

	BFS(testGraph,0)
}