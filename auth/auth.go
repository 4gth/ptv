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
	devID  string
	apiKey string
}

// AuthWriter interface
// Uses the HMAC-SHA1 algorithm to create a signature for the URL.
type AuthWriter interface {
	GenerateSignature(path url.URL) *url.URL
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
func NewAuthWriter() *Auth {
	devID, apiKey := loadEnv()
	return &Auth{
		devID:  devID,
		apiKey: apiKey,
	}
}

// GenerateSignature generates a signed path using the HMAC-SHA1 algorithm.
// Takes a url.URL as input and appends the developer ID and signature to it.
// Returns a pointer to the signed URL.
func (a *Auth) GenerateSignature(path url.URL) *url.URL {
	urlToSign := path.Path + "?" + path.RawQuery
	urlToSign = urlToSign + "devid=" + a.devID

	mac := hmac.New(sha1.New, []byte(a.apiKey))
	mac.Write([]byte(urlToSign))

	signature := hex.EncodeToString(mac.Sum(nil))

	signedURL, err := url.Parse(urlToSign + "&signature=" + signature)
	if err != nil {
		fmt.Println("Error parsing signed URL:", err)
	}
	fmt.Println("Signed URL:", signedURL.String())

	return signedURL
}
