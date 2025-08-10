package model

type RouteTypesPayload struct {
	RouteTypes []RouteTypes `json:"route_types,omitempty"`
	Status     Status       `json:"status"`
}

type RouteType struct{}

func (RouteType) New() *Request {
	return &Request{
		Path:    "/v3/route_types",
		Payload: RouteTypesPayload{},
	}
}
