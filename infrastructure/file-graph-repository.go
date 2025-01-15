package infrastructure

import (
	"bufio"
	"fmt"
	"main/bl"
	"os"
	"path/filepath"
	"strings"
)

type FileGraphRepository struct {
	DirectoryPath string
}

func (repo *FileGraphRepository) Save(graph bl.Graph) error {
	file, err := os.Create(filepath.Join(repo.DirectoryPath, graph.GetId()+".txt"))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	result, err := graph.FindCircuit()
	_, err = writer.WriteString(fmt.Sprintf("%s\n", strings.Join(result, " ")))

	for _, edge := range graph.GetEdges() {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", edge))
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush file writer: %w", err)
	}
	return nil
}

func (repo *FileGraphRepository) Load(graphId string) (*bl.Graph, error) {
	file, err := os.Open(filepath.Join(repo.DirectoryPath, graphId+".txt"))
	if err != nil {
		return &bl.Graph{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
	}

	nodes := make(map[string]bool)
	edges := []bl.Edge{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return &bl.Graph{}, fmt.Errorf("invalid edge format: %s", line)
		}
		nodes[parts[0]] = true
		nodes[parts[1]] = true
		edges = append(edges, *bl.NewEdge(parts[0], parts[1]))
	}

	if err := scanner.Err(); err != nil {
		return &bl.Graph{}, fmt.Errorf("failed to read file: %w", err)
	}

	unique_nodes := []string{}
	for key := range nodes {
		unique_nodes = append(unique_nodes, key)
	}
	return bl.NewGraph(unique_nodes, edges, graphId), nil
}

func GetInMemoryGraphRepositoryInstance() *FileGraphRepository {
	once.Do(func() {
		fileGraphRepoInstance = &FileGraphRepository{DirectoryPath: ""}
		fmt.Println("InMemoryGraphRepository inMemoryRepoInstance created")
	})
	return fileGraphRepoInstance
}
