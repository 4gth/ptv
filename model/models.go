// Package model defines the data structures and methods for the application.
// New requests can be created using the Requester interface.
// The Request struct holds the path, parameters, and payload for HTTP requests.
// ex r := new(model.NewRequest(model.RoutesRequest))
package model

type Request struct {
	Path       string
	Parameters any
	Payload    any
}

type Requester interface {
	New() *Request
}

func NewRequest(funcType Requester) *Request {
	return funcType.New()
}
