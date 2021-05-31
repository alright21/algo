package search

import (
	"fmt"

	"github.com/alright21/algo/graph"
)

// DFS search a Graph by doing depth first search
func DFS (graph *graph.Graph, node int){
	
	graph.Visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbour := range(graph.Adj[node]){
		if graph.Visited[neighbour] == false {
			DFS(graph, neighbour)
		}
	}
	
}