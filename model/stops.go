package model

import "github.com/go-openapi/strfmt"

type StopsPayload struct {
	Stop        Stop        `json:"stop"`
	Disruptions Disruptions `json:"disruptions"`
	Status      Status      `json:"status"`
}

type StopParametersByStopIDAndRouteType struct {
	StopID            int32 `path:"stop_id"`
	RouteType         int32 `path:"route_id"`
	StopLocation      bool  `query:"stop_location"`
	StpAmenities      bool  `query:"stop_amenities"`
	StopAccessibility bool  `query:"stop_accessibility"`
	StopContact       bool  `query:"stop_contact"`
	StopTicket        bool  `query:"stop_ticket"`
	Gtfs              bool  `query:"gtfs"`
	StopStaffing      bool  `query:"stop_staffing"`
	StopDisruptions   bool  `query:"stop_disruptions"`
}

type StopParametersByRouteIDAndRouteType struct {
	RouteID                      int32           `path:"route_id"`
	RouteType                    int32           `path:"route_type"`
	DrirectionID                 int32           `query:"direction_id"`
	StopDisruptions              bool            `query:"stop_disruptions"`
	IncludeGeopath               bool            `query:"include_geopath"`
	GeoPathUTC                   strfmt.DateTime `query:"geopath_utc"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}

type StopParametersByLatLng struct {
	Latitude        float32 `path:"latitude"`
	Longitude       float32 `path:"longitude"`
	RouteType       int32   `query:"route_type"`
	MaxResults      int32   `query:"max_results"`
	MaxDistance     float64 `query:"max_distance"`
	StopDisruptions bool    `query:"stop_disruptions"`
}

type (
	StopsByStopIDAndRouteType  struct{}
	StopsByRouteIDAndRouteType struct{}
	StopsByLatLng              struct{}
)

func (StopsByStopIDAndRouteType) New() *Request[StopsPayload] {
	return &Request[StopsPayload]{
		Path:       "/v3/stops/{stop_id}/route_type/{route_type}",
		Parameters: StopParametersByStopIDAndRouteType{},
		Payload:    StopsPayload{},
	}
}

func (StopsByRouteIDAndRouteType) New() *Request[StopsPayload] {
	return &Request[StopsPayload]{
		Path:       "/v3/stops/route/{route_id}/route_type/{route_type}",
		Parameters: StopParametersByRouteIDAndRouteType{},
		Payload:    StopsPayload{},
	}
}

func (StopsByLatLng) New() *Request[StopsPayload] {
	return &Request[StopsPayload]{
		Path:       "/v3/stops/location/{latitude},{longitude}",
		Parameters: StopParametersByLatLng{},
		Payload:    StopsPayload{},
	}
}
