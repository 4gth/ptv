package model

import "github.com/go-openapi/strfmt"

type RoutePayload struct {
	Route  []Route `json:"route"`
	Status Status  `json:"status"`
}

type RoutesParameters struct {
	RouteTypes []int32 `query:"route_types"`
	RouteName  string  `query:"route_name"`
}

type RoutesParametersByRouteID struct {
	RouteID        int32           `path:"route_id"`
	IncludeGeopath bool            `query:"include_geopath"`
	GeoPathUTC     strfmt.DateTime `query:"geopath_utc"`
}

type (
	RoutesRequest   struct{}
	RoutesByRouteID struct{}
)

func (RoutesRequest) New() *Request {
	return &Request{
		Path:       "/v3/routes",
		Parameters: RoutesParameters{},
		Payload:    RoutePayload{},
	}
}

func (RoutesByRouteID) New() *Request {
	return &Request{
		Path:       "/v3/routes/{route_id}",
		Parameters: RoutesParametersByRouteID{},
		Payload:    RoutePayload{},
	}
}
