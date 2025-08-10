package main

import (
	"fmt"

	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/client"
	"github.com/4gth/ptv/model"
)

const (
	host   = "timetableapi.ptv.vic.gov.au"
	scheme = "https"
)

func main() {
	client := client.NewClient()
	request := model.NewRequest(model.RoutesRequest{})
	authWriter := auth.NewAuthWriter()

	client.SetDefaults(host, "", scheme, authWriter).
		SetQuery(request.Path, request.Parameters)

	client.Get(&request.Payload)

	fmt.Println("Response:", request.Payload)
}
