package bl

import (
	"errors"
	"github.com/google/uuid"
)

var ErrOddNumberOfEdges = errors.New("Graph is not eulerian circle - Number of edges are not even.")
var ErrGraphDisconnected = errors.New("Graph is not eulerian circle - Graph is Disconnected")

// graph.go
type Graph struct {
	nodes []string
	edges []Edge
	id    string
}

func NewGraph(nodes []string, edges []Edge, graphId ...string) *Graph {
	uniqueEdges := unique(edges)

	if len(graphId) == 1 {
		return &Graph{id: graphId[0], nodes: nodes, edges: uniqueEdges}
	}
	return &Graph{id: uuid.New().String(), nodes: nodes, edges: uniqueEdges}
}

func unique(edges []Edge) []Edge {
	var uniqueEdges []Edge
	for _, edge := range edges {
		if !isInArray(edge, uniqueEdges) {
			uniqueEdges = append(uniqueEdges, edge)
		}
	}
	return uniqueEdges
}

func isInArray(element Edge, array []Edge) bool {
	for _, edge := range array {
		if edge.equals(element) {
			return true
		}
	}
	return false
}

func (graph *Graph) removeEdge(edgeToRemove Edge) {
	var newEdges []Edge
	for _, value := range graph.edges {
		if !value.equals(edgeToRemove) {
			newEdges = append(newEdges, value)
		}
	}

	graph.edges = newEdges
}

func (graph *Graph) contains(edge Edge) bool {

	for _, value := range graph.edges {
		if value.equals(edge) {
			return true
		}
	}
	return false
}

func (graph *Graph) FindCircuit() ([]string, error) {
	var path []string
	if graph.hasEachNodeOddDegree() {
		return path, ErrOddNumberOfEdges
	}
	startNode := graph.findStartVert()
	circuitFound := false
	result := graph.solve(startNode, &path, &circuitFound)
	if !graph.isAllNodesInCircuit(result) {
		return []string{}, ErrGraphDisconnected
	}
	return result, nil
}

func (graph *Graph) hasEachNodeOddDegree() bool {
	for _, node := range graph.nodes {
		if len(graph.getNeighboursOf(node))%2 != 0 {
			return true
		}
	}
	return false
}

func (graph *Graph) findStartVert() string {
	for _, node := range graph.nodes {
		numberOfEdges := len(graph.getNeighboursOf(node))
		if numberOfEdges%2 != 0 {
			return node
		}
	}
	return graph.nodes[0]
}

func (graph *Graph) solve(startNode string, path *[]string, circuitFound *bool) []string {
	numberOfEdges := len(graph.edges)
	for _, node := range graph.nodes {
		edge := *NewEdge(startNode, node)
		if graph.contains(edge) {
			if numberOfEdges <= 1 || !graph.isBridge(node) {
				*path = append(*path, startNode)
				graph.removeEdge(edge)
				graph.solve(node, path, circuitFound)
			}
		}
	}
	if numberOfEdges == 0 && !*circuitFound {
		*circuitFound = true
		*path = append(*path, startNode)
	}

	return *path
}

func (graph *Graph) getNeighboursOf(node string) []string {
	neighbours := []string{}
	for _, value := range graph.edges {
		if value.RightPoint == node {
			neighbours = append(neighbours, value.LeftPoint)
		}
		if value.LeftPoint == node {
			neighbours = append(neighbours, value.RightPoint)
		}
	}
	return neighbours
}

func (graph *Graph) isBridge(node string) bool {
	return len(graph.getNeighboursOf(node)) <= 1
}

func (graph *Graph) isAllNodesInCircuit(result []string) bool {
	for _, node := range graph.nodes {
		if !graph.isIn(node, result) {
			return false
		}
	}
	return true
}

func (graph *Graph) isIn(element string, list []string) bool {
	for _, el := range list {
		if el == element {
			return true
		}
	}
	return false
}

func (graph *Graph) GetId() string {
	return graph.id
}
