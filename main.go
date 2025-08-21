package main

import (
	"context"
	"fmt"

	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/model"
	"github.com/4gth/ptv/ptv"
)

const (
	host   = "timetableapi.ptv.vic.gov.au"
	scheme = "https"
)

var a *auth.Auth

func main() {
	ctx := context.Background()

	svc := ptv.NewFromEnv()

	routes, err := svc.Routes(ctx, func(p *model.RoutesParameters) {
		p.RouteName = "Sandringham"
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(len(routes.Route))
	for _, r := range routes.Route {
		fmt.Println(r)
	}
}
