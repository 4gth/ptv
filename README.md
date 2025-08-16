# PTV Go SDK

A lightweight Go client for the Public Transport Victoria (PTV) Timetable API.

- Simple request builder types in `model` with strongly-typed path/query parameters
- An HTTP `client` that expands `{path}` placeholders and builds query strings from struct tags
- An `auth` signer that sends your `devid` and HMAC-SHA1 `signature` on request

## Quick start

- **Requirements**: Go installed, PTV API credentials (Developer ID and API key)
- **Install**:

```bash
go get github.com/4gth/ptv@latest
```

### Configure credentials

Credentials can be set using `Auth` struct passed to `NewAuthWriter(*Auth)`.

Otherwise your credentials can be auto-loaded as environment variables.
A `.env` file is supported via `github.com/joho/godotenv`.

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
var a *auth.Auth{
  devid: "123"
  apiKey: "<your_api_key_here>"
}
func main() {
	// Example use
	client := client.NewClient()
	request := model.NewRequest(model.DeparturesByRouteTypeAndStopIDAndRouteID{})

	request.Parameters.RouteID = 1
	request.Parameters.RouteType = 0 // train
	request.Parameters.StopID = 23

	authWriter := auth.NewAuthWriter(a)

	client.SetDefaults(host, "", scheme, authWriter).
		SetQuery(request.Path, request.Parameters)

	resp, err := client.Get()
	if err != nil {
		fmt.Println(err)
	}
	if err := request.UnMarshalPayload(resp); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", request.Payload.Departures)
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
