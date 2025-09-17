package main

import "fmt"

// Node represent a single vertex
type Node struct{
	Value int
	adj  []*Node     // adjacency list : Slice of pointers to connected nodes
}

// Graph holds all the node
type Graph struct{
	nodes []*Node
}

// AddNode adds a new vertex to the graph
func (g *Graph) AddNode(value int)*Node{
	newNode := &Node{Value: value}
	g.nodes = append(g.nodes, newNode)
	return newNode
}

// AddEdge adds an edge between two nodes 
func (g *Graph) AddEdge(from, to *Node){
	from.adj = append(from.adj, to)
}


// Display prints the adjacency list representation of the graph
func (g *Graph) Display(){
	for _, node := range g.nodes{
		fmt.Printf("%d-> ", node.Value)
		for _, neighbor := range node.adj{
			fmt.Printf("%d -> ",neighbor.Value)
		}
		fmt.Println()
	}
}


func main(){
	graph := &Graph{}

	// create node
	n1 := graph.AddNode(1)
	n2 := graph.AddNode(2)
	n3 := graph.AddNode(3)
	n4 := graph.AddNode(4)

	// Add edge
	graph.AddEdge(n1,n2)
	graph.AddEdge(n1,n3)
	graph.AddEdge(n2,n4)
	graph.AddEdge(n3,n4)

	// Display adjacency list
	graph.Display()
}