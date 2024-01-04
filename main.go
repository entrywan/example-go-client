package main

import (
	"context"
	"fmt"

	openapi "github.com/entrywan/api-go"
)

const secret = "iamtoken"

func main() {
	configuration := openapi.NewConfiguration()
	apiClient := openapi.NewAPIClient(configuration)
	auth := context.WithValue(context.Background(),
		openapi.ContextAccessToken,
		secret)

	res, _, err := apiClient.InstanceAPI.InstanceGet(auth).Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range res {
		name := i.Hostname
		fmt.Printf("%s\n", *name)
	}
}
