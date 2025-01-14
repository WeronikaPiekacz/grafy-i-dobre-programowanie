package bl

import (
	"errors"
	"testing"
)

func TestRemove(t *testing.T) {
	graph := NewGraph([]string{"1", "2", "3"}, []Edge{
		*NewEdge("1", "2"),
		*NewEdge("2", "3"),
		*NewEdge("3", "1"),
	})

	graph.removeEdge(*NewEdge("2", "3"))

	if len(graph.edges) != 2 {
		t.Errorf("Graph size not equal")
	}
	if !graph.contains(*NewEdge("1", "2")) {
		t.Errorf("Edge is missing")
	}
	if !graph.contains(*NewEdge("3", "1")) {
		t.Errorf("Edge is missing")
	}
	if graph.contains(*NewEdge("2", "3")) {
		t.Errorf("Edge shouldn't exists")
	}
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

	if !arraysEquals(result, []string{"0", "1", "2", "0", "4", "2", "3", "4", "5", "0"}) {
		t.Errorf("Should be euler cycle")
	}

}

func TestFleuryWhenGraphPointsHaveOddNumberOfEdges(t *testing.T) {

	graph := NewGraph([]string{"0", "1", "2", "3", "4", "5"}, []Edge{
		*NewEdge("0", "4"),
		*NewEdge("0", "5"),
		*NewEdge("1", "2"),
	})

	_, err := graph.FindCircuit()

	if err == nil {
		t.Errorf("Error not thrown.")
		return
	}

	if !errors.Is(err, ErrOddNumberOfEdges) {
		t.Errorf("Wrong error thrown.")
	}

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
