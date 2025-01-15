package bl

import "fmt"

type InMemoryGraphRepository struct {
	graphs map[string]Graph
}

func (repo *InMemoryGraphRepository) Save(graph Graph) error {
	repo.graphs[graph.GetId()] = graph
	return nil
}

func (repo *InMemoryGraphRepository) Load(graphId string) (*Graph, error) {
	graph := repo.graphs[graphId]
	return &graph, nil
}

func GetInMemoryGraphRepositoryInstance() *InMemoryGraphRepository {
	repoOnce.Do(func() {
		graphs := make(map[string]Graph)
		inMemoryRepoInstance = &InMemoryGraphRepository{graphs: graphs}
		fmt.Println("InMemoryGraphRepository instance created")
	})
	return inMemoryRepoInstance
}
