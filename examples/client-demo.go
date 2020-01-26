package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ocibuilder/gofeas"
)

var configuration = &gofeas.Configuration{
	BasePath:   "http://localhost:8080",
	HTTPClient: &http.Client{},
}

var occurrence = gofeas.V1beta1Occurrence{
	Resource: &gofeas.V1beta1Resource{
		Uri: "http://dockerhub.io/myimage:0.1.1",
	},
	NoteName: "projects/image-signing/notes/production",
	Attestation: &gofeas.V1beta1attestationDetails{
		Attestation: &gofeas.AttestationAttestation{
			PgpSignedAttestation: &gofeas.AttestationPgpSignedAttestation{
				Signature: "this-is-the-signature",
				PgpKeyId:  "key-id",
			},
		}},
}

func main() {

	cli := gofeas.NewAPIClient(configuration)
	res, httpRes, err := cli.CreateOccurrence(context.Background(), "projects/image-signing", occurrence)

	if httpRes != nil && httpRes.StatusCode != http.StatusOK {
		fmt.Println("status code: ", httpRes.Status)
	}

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("res: ", res)
}
