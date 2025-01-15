package infrastructure

import "sync"

var (
	fileGraphRepoInstance *FileGraphRepository
	once                  sync.Once
)
