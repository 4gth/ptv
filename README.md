# PTV Go SDK

A lightweight Go client for the Public Transport Victoria (PTV) Timetable API.

- Simple request builder types in `model` with strongly-typed path/query parameters
- An HTTP `client` that expands `{path}` placeholders and builds query strings from struct tags
- An `auth` signer that sends your `devid` and HMAC-SHA1 `signature` on request
- A high-level `ptv` package that ties everything together behind a single Service

## Quick start

- Requirements: Go installed, PTV API credentials (Developer ID and API key)
- Install:

```bash
go get github.com/4gth/ptv@latest
```

### Configure credentials

- Option 1 (recommended): Environment variables (supports `.env` via github.com/joho/godotenv)

```env
PTV_DEV_ID=1234567
PTV_API_KEY=your_api_key_here
```

- Option 2: Pass credentials directly when constructing the high-level Service

```go
svc := ptv.New("1234567", "your_api_key_here")
```

## High-level usage (ptv.Service)

```go
package main

import (
    "context"
    "fmt"

    "github.com/4gth/ptv/ptv"
    "github.com/4gth/ptv/model"
)

func main() {
    ctx := context.Background()

    // Load from env (or .env)
    svc := ptv.NewFromEnv()

    // Example: list routes filtered by name
    routes, err := svc.Routes(ctx, func(p *model.RoutesParameters) {
        p.RouteName = "Sandringham"
    })
    if err != nil { panic(err) }
    fmt.Println(len(routes.Route))

    // Example: departures for a stop
    deps, err := svc.DeparturesByRouteTypeAndStopID(ctx, func(p *model.DeparturesParameters) {
        p.RouteType = 0 // train
        p.StopID = 1071 // Flinders Street
        p.MaxResults = 10
    })
    if err != nil { panic(err) }
    fmt.Println(len(deps.Departures))
}
```

## Low-level usage (client + model)

```go
package main

import (
    "fmt"
    "github.com/4gth/ptv/auth"
    "github.com/4gth/ptv/client"
    "github.com/4gth/ptv/model"
)

func main() {
    a := &auth.Auth{DevID: "123", APIKey: "<your_api_key_here>"}
    aw := auth.NewAuthWriter(a)

    c := client.NewClient()
    req := model.NewRequest[model.DeparturePayload, model.DeparturesParameters](model.DeparturesByRouteTypeAndStopIDRequest{})

    req.Parameters.RouteType = 0
    req.Parameters.StopID = 23

    c.SetDefaults("timetableapi.ptv.vic.gov.au", "", "https", aw).
        SetQuery(req.Path, req.Parameters)

    resp, err := c.Get()
    if err != nil { panic(err) }
    if err := req.UnMarshalPayload(resp); err != nil { panic(err) }

    fmt.Printf("%+v\n", req.Payload.Departures)
}
```

## Packages

- `auth`: Sign requests using your PTV `devid` and `signature`
- `client`: Minimal HTTP client with path templating and query builder
- `model`: Request builders and payload/parameter types for PTV endpoints
- `ptv`: Unified high-level Service over auth, client, and model

## Environment

- `PTV_DEV_ID`: Your PTV developer ID
- `PTV_API_KEY`: Your PTV API key
- Optional `.env` in repository root is auto-loaded

## License

This project is licensed under the GPL-3.0. See `LICENSE`.
