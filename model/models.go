// Package model defines the data structures and methods for the application.
// New requests can be created using the Requester interface.
// The Request struct holds the path, parameters, and payload for HTTP requests.
// ex r := new(model.NewRequest(model.RoutesRequest))
package model

import "encoding/json"

type Request[T any, Q any] struct {
	Path       string
	Parameters Q
	Payload    T
}

type Requester[T, Q any] interface {
	New() *Request[T, Q]
}

func NewRequest[T, Q any](funcType Requester[T, Q]) *Request[T, Q] {
	return funcType.New()
}

func (r *Request[T, Q]) UnMarshalPayload(data []byte) error {
	return json.Unmarshal(data, &r.Payload)
}
