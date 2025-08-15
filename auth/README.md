# `auth` package

Provides HMAC-SHA1 request signing for the PTV API. It appends `devid` and `signature` query parameters to URLs.

## Credentials

Authentication for calls rely on 2 URL parameters.
`devid=<yourdevid>&signature=<HMAC-SHA1 string>`

Credentials can be passed to `NewAuthWriter`
Fallback if nil parameter passed a `.env` file in the project root is loaded automatically via `github.com/joho/godotenv`.

- `PTV_DEV_ID`: PTV developer ID
- `PTV_API_KEY`: PTV API key

Example `.env`:

```env
PTV_DEV_ID=1234567
PTV_API_KEY=your_api_key_here
```

## Usage

```go
package main

import (
    "fmt"
    "net/url"
    "github.com/4gth/ptv/auth"
)
var a *Auth
func main() {
    aw := auth.NewAuthWriter(a)

    u := url.URL{
        Scheme: "https",
        Host:   "timetableapi.ptv.vic.gov.au",
        Path:   "/v3/routes",
    }

    signed := aw.GenerateSignature(u)
    fmt.Println(signed.String())
}
```

## API

- `type Auth struct` – internal holder of credentials
- `func NewAuthWriter(*Auth) *Auth` – constructs an auth signer loading credentials from parameter or env/.env
- `func (a *Auth) GenerateSignature(path url.URL) string` – returns string with `devid` and `signature`

