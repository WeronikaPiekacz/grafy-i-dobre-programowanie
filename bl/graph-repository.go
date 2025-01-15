package bl

type GraphRepository interface {
	Save(graph Graph) error
	Load(graphId string) (*Graph, error)
}
