package bl

import "sync"

var (
	graphServiceInstance *GraphService
	inMemoryRepoInstance *InMemoryGraphRepository
	serviceOnce          sync.Once
	repoOnce             sync.Once
)
