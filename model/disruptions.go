package model

type DisruptionsPayload struct {
	Disruptions Disruptions `json:"disruptions"`
	Status      Status      `json:"status"`
}

type DisruptionsParameters struct {
	RouteTypes       []int32 `query:"route_types"`
	DisruptionModes  []int32 `query:"disruption_modes"`
	DisruptionStatus string  `query:"disruption_status"`
}

type DisruptionModesPayload struct {
	DisruptionModes []DisruptionModes `json:"disruption_modes"`
	Status          Status            `json:"status"`
}

type DisruptionsParametersByRouteID struct {
	RouteID          int32 `path:"route_id"`
	DisruptionStatus int32 `query:"disruption_status"`
}

type DisruptionsParametersByRouteIDAndStopID struct {
	RouteID          int32 `path:"route_id"`
	StopID           int32 `path:"stop_id"`
	DisruptionStatus int32 `query:"disruption_status"`
}

type DisruptionsParametersByStopID struct {
	StopID           int32 `path:"stop_id"`
	DisruptionStatus int32 `query:"disruption_status"`
}

type DisruptionsParametersByDisruptionID struct {
	DisruptionID int32 `path:"disruption_id"`
}

type (
	Disruption                    struct{}
	DisruptionsByRouteID          struct{}
	DisruptionsByRouteIDAndStopID struct{}
	DisruptionsByStopID           struct{}
	DisruptionByDisruptionID      struct{}
	DisruptionByModes             struct{}
)

func (Disruption) New() *Request[DisruptionsPayload] {
	return &Request[DisruptionsPayload]{
		Path:       "/v3/disruptions",
		Parameters: DisruptionsParameters{},
		Payload:    DisruptionsPayload{},
	}
}

func (DisruptionsByRouteID) New() *Request[DisruptionsPayload] {
	return &Request[DisruptionsPayload]{
		Path:       "/v3/disruptions/route/{route_id}",
		Parameters: DisruptionsParametersByRouteID{},
		Payload:    DisruptionsPayload{},
	}
}

func (DisruptionsByRouteIDAndStopID) New() *Request[DisruptionsPayload] {
	return &Request[DisruptionsPayload]{
		Path:       "/v3/disruptions/route/{route_id}/stop/{stop_id}",
		Parameters: DisruptionsParametersByRouteIDAndStopID{},
		Payload:    DisruptionsPayload{},
	}
}

func (DisruptionsByStopID) New() *Request[DisruptionsPayload] {
	return &Request[DisruptionsPayload]{
		Path:       "/v3/disruptions/stop/{stop_id}",
		Parameters: DisruptionsParametersByStopID{},
		Payload:    DisruptionsPayload{},
	}
}

func (DisruptionByDisruptionID) New() *Request[DisruptionsPayload] {
	return &Request[DisruptionsPayload]{
		Path:       "/v3/disruptions/{disruption_id}",
		Parameters: DisruptionsParametersByDisruptionID{},
		Payload:    DisruptionsPayload{},
	}
}

func (DisruptionByModes) New() *Request[DisruptionModesPayload] {
	return &Request[DisruptionModesPayload]{
		Path:       "/v3/disruptions/modes",
		Parameters: DisruptionsParameters{},
		Payload:    DisruptionModesPayload{},
	}
}
