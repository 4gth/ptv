// Package ptv provides a convinent interface for creating instances of a full PTV client.
package ptv

import (
	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/client"
)

var a auth.Auth

const (
	host   = "timetableapi.ptv.vic.gov.au"
	scheme = "https"
)

type Client struct {
	Client client.Client
}

func PTVClient() *Client {
	return &Client{}
}

func (c *Client) InitClient(aw auth.AuthWriter) *Client {
	c.Client.SetDefaults(host, "", scheme, aw)
	return c
}
