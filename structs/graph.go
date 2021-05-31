package structs

// Graph is a struct representing a graph, with int value as nodes and adjancency lists for edges
type Graph struct {

	Visited map[int]bool
	Adj map[int][]int 
}

// AddEdge adds an edge from source to destination
func (g *Graph) AddEdge(source int, destination int){

	g.Adj[source] = append(g.Adj[source], destination)
}