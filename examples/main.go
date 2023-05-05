package main

import (
	"context"
	"fmt"

	dockerhub "github.com/berrybytes/dockerhub-go"
)

func main() {
	client := dockerhub.NewClient(nil)
	err := client.Auth.Login(context.Background(), "<DOCKER_USERNAME>", "<DOCKER_PERSONAL_TOKEN>")
	if err != nil {
		return
	}
	list, err := client.Repositories.GetRepositories(context.Background(), "namespace", &dockerhub.ListOptions{
		Page:     1,
		PageSize: 20,
	})
	if err != nil {
		return
	}
	fmt.Println(len(list.Results), list.Count)
	for _, data := range list.Results {
		fmt.Println(data.Name)
	}
}
