package model

import "github.com/go-openapi/strfmt"

type DeparturePayload struct {
	Departures  []Departures `json:"departures,omitempty"`
	Stops       Stops        `json:"stops"`
	Routes      Routes       `json:"routes"`
	Runs        Runs         `json:"runs"`
	Directions  Directions   `json:"directions"`
	Disruptions Disruptions  `json:"disruptions"`
	Status      Status       `json:"status"`
}

type DeparturesParameters struct {
	RouteType        int32           `path:"route_type"`
	StopID           int32           `path:"stop_id"`
	PlatformNumbers  []int32         `query:"platform_numbers"`
	DirectionID      int32           `query:"direction_id"`
	Gtfs             bool            `query:"gtfs"`
	DateUtc          strfmt.DateTime `query:"date_utc"`
	MaxResults       int32           `query:"max_results"`
	IncludeCancelled bool            `query:"include_cancelled"`
	LookBackwards    bool            `query:"look_backwards"`
	Expand           []int32         `query:"expand"`
	IncludeGeopath   bool            `query:"include_geopath"`
}

type DeparturesParametersByRoute struct {
	RouteType        int32           `path:"route_type"`
	StopID           int32           `path:"stop_id"`
	RouteID          int32           `path:"route_id"`
	PlatformNumbers  []int32         `query:"platform_numbers"`
	DirectionID      int32           `query:"direction_id"`
	Gtfs             bool            `query:"gtfs"`
	DateUtc          strfmt.DateTime `query:"date_utc"`
	MaxResults       int32           `query:"max_results"`
	IncludeCancelled bool            `query:"include_cancelled"`
	LookBackwards    bool            `query:"look_backwards"`
	Expand           []int32         `query:"expand"`
	IncludeGeopath   bool            `query:"include_geopath"`
}

type (
	DeparturesByRouteTypeAndStopIDRequest    struct{}
	DeparturesByRouteTypeAndStopIDAndRouteID struct{}
)

func (DeparturesByRouteTypeAndStopIDRequest) New() *Request[DeparturePayload, DeparturesParameters] {
	return &Request[DeparturePayload, DeparturesParameters]{
		Path:       "/v3/departures/route_type/{route_type}/stop/{stop_id}",
		Parameters: DeparturesParameters{},
		Payload:    DeparturePayload{},
	}
}

func (DeparturesByRouteTypeAndStopIDAndRouteID) New() *Request[DeparturePayload, DeparturesParametersByRoute] {
	return &Request[DeparturePayload, DeparturesParametersByRoute]{
		Path:       "/v3/departures/route_type/{route_type}/stop/{stop_id}/route/{route_id}",
		Parameters: DeparturesParametersByRoute{},
		Payload:    DeparturePayload{},
	}
}
