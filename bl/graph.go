package bl

import (
	"errors"
)

var ErrOddNumberOfEdges = errors.New("number of edges are not even")

// graph.go
type Graph struct {
	nodes []string
	edges []Edge
}

type Result struct {
	isEulerCycle bool
	cycle        []string
}

func NewGraph(nodes []string, edges []Edge) *Graph {
	return &Graph{nodes: nodes, edges: edges}
}

func (graph *Graph) removeEdge(edgeToRemove Edge) {
	// TODO: Implement
}

func (graph *Graph) contains(edge Edge) bool {
	// TODO: Implement
	return false
}

func (graph *Graph) FindCircuit() ([]string, error) {
	// TODO: Implement
	return []string{}, nil
}
