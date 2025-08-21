// Package auth provides functionality for generating signed URLs for API requests.
// Provides an interface for generating signatures based on the developer ID and API key.
package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

/*----Auth----*/

type Auth struct {
	DevID  string
	APIKey string
}

// AuthWriter interface
// Uses the HMAC-SHA1 algorithm to create a signature for the URL.
type AuthWriter interface {
	GenerateSignature(path url.URL) string
}

// Load environment variables from a .env file.
// Returns the developer ID and API key.
func loadEnv() (string, string) {
	godotenv.Load()
	return os.Getenv("PTV_DEV_ID"), os.Getenv("PTV_API_KEY")
}

// NewAuthWriter Creates a new instance of Auth with the developer ID and API key
// loaded from environment variables.
// Returns a pointer to the Auth object.
func NewAuthWriter(auth *Auth) *Auth {
	if auth == nil {
		devID, apiKey := loadEnv()
		return &Auth{
			DevID:  devID,
			APIKey: apiKey,
		}
	}
	return &Auth{
		DevID:  auth.DevID,
		APIKey: auth.APIKey,
	}
}

// GenerateSignature generates a signed path using the HMAC-SHA1 algorithm.
// Takes a url.URL as input and appends the developer ID and signature to it.
// returns the devID and generated signature as url parameter strings
func (a *Auth) GenerateSignature(path url.URL) string {
	urlToSign := path.Path + "?" + path.RawQuery
	var devID string
	if path.RawQuery == "" {
		devID = "?devid=" + a.DevID
		urlToSign = urlToSign + "devid=" + a.DevID

	} else {
		devID = "&devid=" + a.DevID
		urlToSign = urlToSign + devID
	}
	mac := hmac.New(sha1.New, []byte(a.APIKey))
	mac.Write([]byte(urlToSign))

	signature := hex.EncodeToString(mac.Sum(nil))
	devIDSigString := fmt.Sprintf("%s&signature=%s", devID, signature)
	return devIDSigString
}
