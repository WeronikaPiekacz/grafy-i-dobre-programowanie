package bl

import (
	"testing"
)

func TestRemove(t *testing.T) {
	graph := NewGraph([]string{"1", "2", "3"}, []Edge{
		*NewEdge("1", "2"),
		*NewEdge("2", "3"),
		*NewEdge("3", "1"),
	})

	graph.removeEdge(*NewEdge("2", "3"))

	assertNumberOfEdges(2, len(graph.edges), t)
	assertEdgeExists(*NewEdge("1", "2"), graph, t)
	assertEdgeExists(*NewEdge("3", "1"), graph, t)
	assertEdgeNotExists(*NewEdge("2", "3"), graph, t)
}

func TestFleury(t *testing.T) {
	graph := NewGraph([]string{"0", "1", "2", "3", "4", "5"}, []Edge{
		*NewEdge("0", "1"),
		*NewEdge("0", "2"),
		*NewEdge("0", "4"),
		*NewEdge("0", "5"),
		*NewEdge("1", "2"),
		*NewEdge("2", "3"),
		*NewEdge("2", "4"),
		*NewEdge("3", "4"),
		*NewEdge("4", "5"),
	})

	result, _ := graph.FindCircuit()

	assertEulerianCircuitEquals([]string{"0", "1", "2", "0", "4", "2", "3", "4", "5", "0"},result,t)

}

func TestFleuryWhenGraphPointsHaveOddNumberOfEdges(t *testing.T) {

	graph := NewGraph([]string{"0", "1", "2", "3", "4", "5"}, []Edge{
		*NewEdge("0", "4"),
		*NewEdge("0", "5"),
		*NewEdge("1", "2"),
	})

	_, err := graph.FindCircuit()

	assertErrorIs(ErrOddNumberOfEdges, err, t)

}

func arraysEquals(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func assertNumberOfEdges(expectedNumberOfEdges, actualNumberOfEdges int, t *testing.T) {
	if expectedNumberOfEdges != actualNumberOfEdges {
		t.Errorf("Expected number of edges: %d, actual number of edges: %d", expectedNumberOfEdges, actualNumberOfEdges)
	}
}

func assertEdgeExists(edge Edge, graph *Graph, t *testing.T) {
	if !graph.contains(edge) {
		t.Errorf("Edge is missing")
	}
}

func assertEdgeNotExists(edge Edge, graph *Graph, t *testing.T) {
	if graph.contains(edge) {
		t.Errorf("Edge shouldn't exists")
	}
}

func assertEulerianCircuitEquals(expected []string, actual []string, t *testing.T) {
	if !arraysEquals(expected, actual) {
		t.Errorf("Arrays are not equal")
	}
}

func assertErrorIs(expected error, actual error, t *testing.T) {
	if actual == nil {
		t.Errorf("Error not thrown.")
		return
	}
	if expected != actual {
		t.Errorf("Expected error: %v, actual error: %v", expected, actual)
	}
}
