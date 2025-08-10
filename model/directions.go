package model

type DirectionsPayload struct {
	Directions []Directions `json:"directions"`
	Status     Status       `json:"status"`
}

type DirectionsParametersByRoute struct {
	RouteID int32 `path:"route_id"`
}

type DirectionsParametersByDirection struct {
	DirecionID int32 `path:"direction_id"`
}

type DirectionsParametersByDirectionAndRouteType struct {
	DirecionID int32 `path:"direction_id"`
	RouteType  int32 `path:"route_type"`
}

type (
	DirectionsByRouteID             struct{}
	DirectionsByDirectionID         struct{}
	DirectionsByRouteIDAndRouteType struct{}
)

func (DirectionsByRouteID) New() *Request {
	return &Request{
		Path:       "/v3/directions/route/{route_id}",
		Parameters: DirectionsParametersByRoute{},
		Payload:    DirectionsPayload{},
	}
}

func (DirectionsByDirectionID) New() *Request {
	return &Request{
		Path:       "/v3/directions/{direction_id}",
		Parameters: DirectionsParametersByDirection{},
		Payload:    DirectionsPayload{},
	}
}

func (DirectionsByRouteIDAndRouteType) New() *Request {
	return &Request{
		Path:       "/v3/directions/{direction_id}/route_type/{route_type}",
		Parameters: DirectionsParametersByDirectionAndRouteType{},
		Payload:    DirectionsPayload{},
	}
}
