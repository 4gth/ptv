package main

import (
	"fmt"
	"log"

	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/client"
	"github.com/4gth/ptv/model"
)

const (
	host   = "timetableapi.ptv.vic.gov.au"
	scheme = "https"
)

func main() {
	// Example use
	client := client.NewClient()
	request := model.NewRequest(model.RoutesByRouteID{})

	request.Parameters = model.RoutesParametersByRouteID{
		RouteID: 1,
	}

	authWriter := auth.NewAuthWriter()

	client.SetDefaults(host, "", scheme, authWriter).
		SetQuery(request.Path, request.Parameters)

	resp, err := client.Get()
	if err != nil {
		fmt.Println(err)
	}
	if err := request.UnMarshalPayload(resp); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", request.Payload.Route)
}
