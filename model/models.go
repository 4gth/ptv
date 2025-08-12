// Package model defines the data structures and methods for the application.
// New requests can be created using the Requester interface.
// The Request struct holds the path, parameters, and payload for HTTP requests.
// ex r := new(model.NewRequest(model.RoutesRequest))
package model

import "encoding/json"

type Request[T any] struct {
	Path       string
	Parameters any
	Payload    T
}

type Requester[T any] interface {
	New() *Request[T]
}

func NewRequest[T any](funcType Requester[T]) *Request[T] {
	return funcType.New()
}

func (r *Request[T]) UnMarshalPayload(data []byte) error {
	return json.Unmarshal(data, &r.Payload)
}
