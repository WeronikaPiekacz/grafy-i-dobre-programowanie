package main

import (
	"fmt"
	"main/bl"
	"main/config"
	"main/infrastructure"
)

var (
	service *bl.GraphService
)

func main() {

	var configService = config.GetInstance("config.json")
	configService = config.GetInstance("config.json")

	configuration := configService.ReadConfig()
	if configuration.IsFileBasedGraphReadingEnabled {
		service = bl.GetGraphServiceInstance(infrastructure.GetInMemoryGraphRepositoryInstance())
		result, err := service.Solve("example_2_improved")
		if err != nil {
			fmt.Printf(err.Error())
		}

		for _, str := range result {
			fmt.Printf(" %s ", str)
		}
	} else {
		service = bl.GetGraphServiceInstance(bl.GetInMemoryGraphRepositoryInstance())
		id, err := service.Create([]string{"0", "1", "2", "3", "4", "5"}, [][]string{
			{"0", "1"}, {"0", "2"}, {"0", "4"}, {"0", "5"}, {"1", "2"}, {"2", "3"}, {"2", "4"}, {"3", "4"}, {"4", "5"}})

		if err != nil {
			fmt.Printf(err.Error())
		}

		result, err := service.Solve(id)

		if err != nil {
			fmt.Printf(err.Error())
		}

		for _, str := range result {
			fmt.Printf(" %s ", str)
		}
	}

}
