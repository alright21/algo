package search

import (
	"fmt"

	"github.com/alright21/algo/structs"
)

// DFS search a Graph by doing depth first search
func DFS (graph *structs.Graph, node int){
	
	graph.Visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbour := range(graph.Adj[node]){
		if graph.Visited[neighbour] == false {
			DFS(graph, neighbour)
		}
	}
	
}