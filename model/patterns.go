package model

import "github.com/go-openapi/strfmt"

type PatternsPayload struct {
	Disruptions []Disruptions `json:"disruptions"`
	Departures  []Departures  `json:"departures"`
	Stops       Stops         `json:"stops"`
	Routes      Routes        `json:"routes"`
	Runs        Runs          `json:"runs"`
	Directions  Directions    `json:"directions"`
	Status      Status        `json:"status"`
}

type PatternsParameters struct {
	RunRef                       string          `path:"run_ref"`
	RouteType                    int32           `path:"route_type"`
	Expand                       []int32         `query:"expand"`
	StopID                       int32           `query:"stop_id"`
	DateUtc                      strfmt.DateTime `query:"date_utc"`
	IncludeSkippedStops          bool            `query:"include_skipped_stops"`
	IncludeGeopath               bool            `query:"include_geopath"`
	IncludeAdvertisedInterchange bool            `query:"include_advertised_interchange"`
}

type PatternByRunRefAndRouteType struct{}

func (PatternByRunRefAndRouteType) New() *Request[PatternsPayload] {
	return &Request[PatternsPayload]{
		Path:       "/v3/pattern/run/{run_ref}/route_type/{route_type}",
		Parameters: PatternsParameters{},
		Payload:    PatternsPayload{},
	}
}
