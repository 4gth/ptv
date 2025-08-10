# PTV Go SDK

A lightweight Go client for the Public Transport Victoria (PTV) Timetable API. It provides:

- Simple request builder types in `model` with strongly-typed path/query parameters
- An HTTP `client` that expands `{path}` placeholders and builds query strings from struct tags
- An `auth` signer that appends your `devid` and HMAC-SHA1 `signature` to each request

## Quick start

- **Requirements**: Go installed, PTV API credentials (Developer ID and API key)
- **Install**:

```bash
go get github.com/4gth/ptv@latest
```

### Configure credentials

Set your credentials as environment variables. A `.env` file is supported via `github.com/joho/godotenv`.

```env
PTV_DEV_ID=1234567
PTV_API_KEY=your_api_key_here
```

### Example

```go
package main

import (
    "fmt"
    "github.com/4gth/ptv/auth"
    "github.com/4gth/ptv/client"
    "github.com/4gth/ptv/model"
)

const (
    host   = "timetableapi.ptv.vic.gov.au"
    scheme = "https"
)

func main() {
    c := client.NewClient()
    req := model.NewRequest(model.RoutesRequest{})
    aw := auth.NewAuthWriter()

    c.SetDefaults(host, "", scheme, aw).
        SetQuery(req.Path, req.Parameters)

    if err := c.Get(&req.Payload); err != nil {
        panic(err)
    }

    fmt.Printf("%+v\n", req.Payload)
}
```

## Packages

- [`auth`](auth/README.md): Sign requests using your PTV `devid` and `signature`
- [`client`](client/README.md): Minimal HTTP client with path templating and query builder
- [`model`](model/README.md): Request builders and payload/parameter types for PTV endpoints

## Environment

- `PTV_DEV_ID`: Your PTV developer ID
- `PTV_API_KEY`: Your PTV API key
- Optional `.env` in repository root is auto-loaded

## License

This project is licensed under the GPL-3.0. See [`LICENSE`](LICENSE).