package model

import "github.com/go-openapi/strfmt"

type RunsPayload struct {
	Runs   []Runs `json:"runs,omitempty"`
	Status Status `json:"status"`
}

type RunsParametersByRouteID struct {
	RouteID                      int32           `path:"route_id"`
	Expand                       []int32         `query:"expand"`
	DateUtc                      strfmt.DateTime `query:"date_utc"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}

type RunsParametersByRouteIDAndRouteType struct {
	RouteID                      int32           `path:"route_id"`
	RouteType                    int32           `path:"route_type"`
	Expand                       []int32         `query:"expand"`
	DateUtc                      strfmt.DateTime `query:"date_utc"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}

type RunsParametersByRunRef struct {
	RunRef                       string          `path:"run_ref"`
	Expand                       []int32         `query:"expand"`
	DateUtc                      strfmt.DateTime `query:"date_utc"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}

type RunsParametersByRunRefAndRouteType struct {
	RunRef                       string          `path:"run_ref"`
	RouteType                    int32           `path:"route_type"`
	Expand                       []int32         `query:"expand"`
	DateUtc                      strfmt.DateTime `query:"date_utc"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}
type (
	RunsByRouteID             struct{}
	RunsByRouteIDAndRouteType struct{}
	RunsByRunRef              struct{}
	RunsByRunRefAndRouteType  struct{}
)

func (RunsByRouteID) New() *Request {
	return &Request{
		Path:       "/v3/runs/route/{route_id}",
		Parameters: RunsParametersByRouteID{},
		Payload:    RunsPayload{},
	}
}

func (RunsByRouteIDAndRouteType) New() *Request {
	return &Request{
		Path:       "/v3/runs/route/{route_id}/route_type/{route_type}",
		Parameters: RunsParametersByRouteIDAndRouteType{},
		Payload:    RunsPayload{},
	}
}

func (RunsByRunRef) New() *Request {
	return &Request{
		Path:       "/v3/runs/{run_ref}",
		Parameters: RunsParametersByRunRef{},
		Payload:    RunsPayload{},
	}
}

func (RunsByRunRefAndRouteType) New() *Request {
	return &Request{
		Path:       "/v3/runs/{run_ref}/route_type/{route_type}",
		Parameters: RunsParametersByRunRefAndRouteType{},
		Payload:    RunsPayload{},
	}
}
