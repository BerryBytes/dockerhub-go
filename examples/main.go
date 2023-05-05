package main

import (
	"context"
	"fmt"

	dockerhub "github.com/berrybytes/dockerhub-go"
)

func main() {
	client := dockerhub.NewClient(nil)
	err := client.Auth.Login(context.Background(), "sugamdocker35", "594c9bff-b4f9-4c2f-b056-81dd7559ba97")

	if err != nil {
		return
	}
	list, err := client.Repositories.GetRepositories(context.Background(), "sugamdocker35", &dockerhub.ListOptions{
		PageSize: 20,
	})
	fmt.Println(len(list.Results), list.Count)
	for _, data := range list.Results {
		fmt.Println(data.Name)
	}
}
