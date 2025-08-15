package model

type RouteTypesPayload struct {
	RouteTypes []RouteTypes `json:"route_types,omitempty"`
	Status     Status       `json:"status"`
}

type RouteTypeParameters struct{}

type RouteType struct{}

func (RouteType) New() *Request[RouteTypesPayload, RouteTypeParameters] {
	return &Request[RouteTypesPayload, RouteTypeParameters]{
		Path:    "/v3/route_types",
		Payload: RouteTypesPayload{},
	}
}
