package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/artbegolli/grafeas"
)

var configuration = &grafeas.Configuration{
	BasePath:   "http://localhost:8080",
	HTTPClient: &http.Client{},
}

var occurrence = grafeas.V1beta1Occurrence{
	Resource: &grafeas.V1beta1Resource{
		Uri: "http://dockerhub.io/myimage:0.1.0",
	},
	NoteName: "production",
}

func main() {

	cli := grafeas.NewAPIClient(configuration)
	res, httpRes, err := cli.GrafeasV1Beta1Api.CreateOccurrence(context.Background(), "projects/image-signing", occurrence)

	if httpRes.StatusCode != http.StatusOK {
		fmt.Println("status code: ", httpRes.Status)
	}

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("res: ", res)
}
