package bl

import "fmt"

type GraphService struct {
	repository GraphRepository
}

func (service GraphService) Create(nodes []string, edges [][]string) (string, error) {
	graph_edges := []Edge{}
	for _, edge := range edges {
		graph_edges = append(graph_edges, *NewEdge(edge[0], edge[1]))
	}
	graph := NewGraph(nodes, graph_edges)
	err := service.repository.Save(*graph)
	if err != nil {
		return "", fmt.Errorf("failed to create graph: %w", err)
	}
	return graph.GetId(), nil
}

func (service GraphService) Solve(graphId string) ([]string, error) {
	graph, _ := service.repository.Load(graphId)
	result, err := graph.FindCircuit()
	if err != nil {
		return nil, fmt.Errorf("failed to create graph: %w", err)
	}
	return result, nil
}

func GetGraphServiceInstance(repository GraphRepository) *GraphService {
	serviceOnce.Do(func() {
		graphServiceInstance = &GraphService{repository: repository}
		fmt.Println("GraphService instance created")
	})
	return graphServiceInstance
}
