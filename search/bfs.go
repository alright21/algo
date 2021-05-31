package search

import (
	"fmt"

	"github.com/alright21/algo/structs"
)

// BFS search the graph using Breadth First Search
func BFS(graph *structs.Graph, node int){

	queue := structs.Queue{}

	graph.Visited[node] = true

	for queue.Len() != 0 {

		element := queue.Dequeue()

		fmt.Printf("%d ", element)

		for _, neighbour := range(graph.Adj[node]){
			if graph.Visited[neighbour] == false {
				queue.Enqueue(neighbour)
			}
		}

	}
}