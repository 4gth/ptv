package model

type OutletsPayload struct {
	Outlets []Outlets `json:"outlets,omitempty"`
	Status  Status    `json:"status"`
}

type OutletsParameters struct {
	MaxResults int32 `query:"max_results"`
}

type OutletsParametersByLatLng struct {
	Latitude   float32 `path:"latitude"`
	Longitude  float32 `path:"longitude"`
	MaxDisance float64 `query:"max_distance"`
	MaxResults int32   `query:"max_results"`
}

type (
	Outlet          struct{}
	OutletsByLatLng struct{}
)

func (Outlet) New() *Request[OutletsPayload] {
	return &Request[OutletsPayload]{
		Path:       "/v3/outlets",
		Parameters: OutletsParameters{},
		Payload:    OutletsPayload{},
	}
}

func (OutletsByLatLng) New() *Request[OutletsPayload] {
	return &Request[OutletsPayload]{
		Path:       "/v3/outlets/location/{latitude},{longitude}",
		Parameters: OutletsParametersByLatLng{},
		Payload:    OutletsPayload{},
	}
}
