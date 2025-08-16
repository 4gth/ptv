# `client` package

Minimal HTTP client that:

- Expands `{path}` placeholders using struct fields tagged with `path:"..."`
- Builds query strings from struct fields tagged with `query:"..."`
- Signs requests via an injected `auth.AuthWriter`

## API overview

- `func NewClient() *Client`
- `func (c *Client) SetDefaults(host, path, scheme string, authWriter auth.AuthWriter) *Client`
- `func (c *Client) SetQuery(path string, query any) *Client` – sets path+params; expands `{}` and builds `?query`
- `func (c *Client) Get() error` – executes GET

## Example: list routes

```go
package main

import (
    "fmt"
    "github.com/4gth/ptv/auth"
    "github.com/4gth/ptv/client"
    "github.com/4gth/ptv/model"
)
var a *auth.Auth
func main() {
    c := client.NewClient()
    aw := auth.NewAuthWriter(a)

    req := model.NewRequest(model.RoutesRequest{})

    c.SetDefaults("timetableapi.ptv.vic.gov.au", "", "https", aw).
        SetQuery(req.Path, req.Parameters)

    if err := c.Get(); err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", req.Payload)
}
```

## Example: stops by lat/lng with query params

```go
req := model.NewRequest(model.StopsByLatLng{})
req.Parameters.Latitude = -37.8183
req.Parameters.Longitude = 144.9671
req.Parameters.RouteType = 0 // train
req.Parameters.MaxResults = 5

c.SetDefaults("timetableapi.ptv.vic.gov.au", "", "https", aw).
  SetQuery(req.Path, req.Parameters)

if err := c.Get(); err != nil {
  // handle error
}
``
Notes:
- Fields with zero values are omitted from the query string.- Path placeholders must have a matching struct field tagged with `path:"name"`.
