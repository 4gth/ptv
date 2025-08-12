package model

type SearchPayload struct {
	Stops   []Stops   `json:"stops,omitempty"`
	Routes  []Routes  `json:"routes,omitempty"`
	Outlets []Outlets `json:"outlets,omitempty"`
	Status  Status    `json:"status"`
}

type SearchParameters struct {
	SearchTerm            string  `path:"search_term"`
	RouteTypes            []int32 `query:"route_types"`
	Latitude              float32 `query:"latitude"`
	Longitude             float32 `query:"longitude"`
	MaxDistance           float64 `query:"max_distance"`
	IncludeAddresses      bool    `query:"include_addresses"`
	IncludeOutlets        bool    `query:"include_outlets"`
	MatchStopBySuburb     bool    `query:"match_stop_by_suburb"`
	MatchRouteBySuburb    bool    `query:"match_route_by_suburb"`
	MatchStopByGtfsStopID bool    `query:"match_stop_by_gtfs_stop_id"`
}
type SearchTerm struct{}

func (SearchTerm) New() *Request[SearchPayload] {
	return &Request[SearchPayload]{
		Path:       "/v3/search/{search_term}",
		Parameters: SearchParameters{},
		Payload:    SearchPayload{},
	}
}
