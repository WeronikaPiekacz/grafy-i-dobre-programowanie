package bl

import (
	"errors"
	"testing"
)

func TestCreate(t *testing.T) {
	// given
	mockRepo := &MockGraphRepository{graphs: make(map[string]Graph)}
	instanceUnderTest := &GraphService{repository: mockRepo}
	// when
	_, err := instanceUnderTest.Create([]string{"0", "1", "2", "3", "4", "5"}, [][]string{
		{"0", "1"}, {"0", "2"}, {"0", "4"}, {"0", "5"}, {"1", "2"}, {"2", "3"}, {"2", "4"}, {"3", "4"}, {"4", "5"}})

	// then
	if err != nil {
		t.Errorf("Not expected error thrown %e", err)
		return
	}
	assertNumberOfGraphs(1, mockRepo, t)
}

func TestCreateWhenMalfunctionHappened(t *testing.T) {
	// given
	mockRepo := &MockGraphRepository{graphs: make(map[string]Graph), simulateMalfunction: true}
	instanceUnderTest := &GraphService{repository: mockRepo}
	// when
	_, err := instanceUnderTest.Create([]string{"0", "1", "2", "3", "4", "5"}, [][]string{
		{"0", "1"}, {"0", "2"}, {"0", "4"}, {"0", "5"}, {"1", "2"}, {"2", "3"}, {"2", "4"}, {"3", "4"}, {"4", "5"}})

	// then
	if err == nil {
		t.Errorf("Error not thrown.")
		return
	}
	assertNumberOfGraphs(0, mockRepo, t)
}

func TestSolve(t *testing.T) {
	// given
	mockRepo := &MockGraphRepository{graphs: make(map[string]Graph)}
	instanceUnderTest := &GraphService{repository: mockRepo}
	id, _ := instanceUnderTest.Create([]string{"0", "1", "2", "3", "4", "5"}, [][]string{
		{"0", "1"}, {"0", "2"}, {"0", "4"}, {"0", "5"}, {"1", "2"}, {"2", "3"}, {"2", "4"}, {"3", "4"}, {"4", "5"}})
	// when
	result, err := instanceUnderTest.Solve(id)

	// then
	if err != nil {
		t.Errorf("Not expected error thrown %e", err)
		return
	}
	assertArraysEqual([]string{"0", "1", "2", "0", "4", "2", "3", "4", "5", "0"}, result, t)
}

func assertNumberOfGraphs(expected int, repo *MockGraphRepository, t *testing.T) {
	actual := len(repo.graphs)
	if actual != expected {
		t.Errorf("Number of graphs is different expected: %d actual: %d", expected, actual)
	}
}

func assertArraysEqual(expected, actual []string, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Arrays are not equal expected array len is equal %d but actual array len is %d", len(expected), len(actual))
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Errorf("Arrays are has different element at index: %d", i)
		}
	}
}

type MockGraphRepository struct {
	graphs              map[string]Graph
	simulateMalfunction bool
}

func (repo *MockGraphRepository) Save(graph Graph) error {
	if repo.simulateMalfunction {
		return errors.New("Malfunction happened")
	}
	repo.graphs[graph.GetId()] = graph
	return nil
}

func (repo *MockGraphRepository) Load(graphId string) (*Graph, error) {
	graph := repo.graphs[graphId]
	return &graph, nil
}
