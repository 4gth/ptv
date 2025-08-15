package client

import (
	"fmt"
	"testing"

	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/model"
)

const (
	host   = "timetableapi.ptv.vic.gov.au"
	scheme = "https"
)

var a *auth.Auth

func TestDeparturesGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DeparturesByRouteTypeAndStopIDRequest{})
	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.RouteType = 0
	r.Parameters.StopID = 1002

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Departures != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Println(r.Payload.Status.Health)
}

func TestDeparturesRouteGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DeparturesByRouteTypeAndStopIDAndRouteID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.RouteType = 0
	r.Parameters.StopID = 1002
	r.Parameters.RouteID = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Departures != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Println(r.Payload.Status.Health)
}

func TestDirectionsByRouteGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DirectionsByRouteID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.RouteID = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Directions != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Println(r.Payload.Status.Health)
}

func TestDirectionsByDirectionIDGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DirectionsByDirectionID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.DirecionID = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Directions != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Println(r.Payload.Status.Health)
}

func TestDirectionsByDirectionIDAndRouteTypeGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DirectionsByRouteIDAndRouteType{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.DirecionID = 1
	r.Parameters.RouteType = 0

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Directions != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Println(r.Payload.Status.Health)
}

func TestDisruptionsGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.Disruption{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Disruptions.General != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestDisruptionsByRouteIDGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DisruptionsByRouteID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.RouteID = 1
	r.Parameters.DisruptionStatus = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Disruptions.General != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestDisruptionsByRouteIDAndStopIDGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DisruptionsByRouteIDAndStopID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.RouteID = 1
	r.Parameters.StopID = 3
	r.Parameters.DisruptionStatus = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Disruptions.General != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestDisruptionsByStopIDGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DisruptionsByStopID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.StopID = 3
	r.Parameters.DisruptionStatus = 1

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Disruptions.General != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestDisruptionsByDisruptionIDGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DisruptionByDisruptionID{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.DisruptionID = 341423

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.Disruptions.General != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestDisruptionsByModeGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.DisruptionByModes{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 && r.Payload.DisruptionModes != nil {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestFairEstimateGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.FairEstimate{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.MaxZone = 6
	r.Parameters.MinZone = 1
	r.Parameters.IsJourneyInFreeTramZone = false
	r.Parameters.TravelledRouteTypes = []int32{0, 1, 2}

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.FareEstimateResultStatus.StatusCode != 0 {
		t.Error("Status Code != 1 Path: ", r.Path)
	}
	fmt.Printf("1 %d\n", r.Payload.FareEstimateResultStatus.StatusCode)
}

func TestOutletsGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.Outlet{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.MaxResults = 15 // feature seems broken on server end

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}

func TestOutletsByLatLngGetFunctions(t *testing.T) {
	c := NewClient()
	r := model.NewRequest(model.OutletsByLatLng{})

	aw := auth.NewAuthWriter(a)

	c.SetDefaults(host, "", scheme, aw)

	r.Parameters.MaxResults = 15 // feature seems broken on server end
	r.Parameters.Latitude = -37.8501
	r.Parameters.Longitude = 145.0002
	r.Parameters.MaxDisance = 200

	c.SetQuery(r.Path, r.Parameters)
	resp, err := c.Get()
	if err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}

	if err := r.UnMarshalPayload(resp); err != nil {
		t.Error(err)
		fmt.Println("Path: ", r.Path)
	}
	if r.Payload.Status.Health != 1 {
		t.Error("Path: ", r.Path)
	}
	fmt.Printf("%d\n", r.Payload.Status.Health)
}
