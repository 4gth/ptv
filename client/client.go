// Package client provides a simple HTTP client for making requests to the PTV API.
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/4gth/ptv/auth"
)

/*-- Client ----*/

// TODO Add Context support to the HTTP client methods for better control over request lifetimes
// TODO Add error handling for HTTP status codes other than 200 OK.

type Client struct {
	client     *http.Client
	url        url.URL
	params     any
	AuthWriter auth.AuthWriter
	Context    context.Context
}

// NewClient creates a new instance of Client with the default HTTP client.
// It initializes the client with default settings and returns a pointer to the Client object.
// It does not set any URL or parameters initially.
// Use this method to create a new client instance before configuring it with SetDefaults or other methods
func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

// SetDefaults sets the default values for the Client.
func (c *Client) SetDefaults(host, path, scheme string, authWriter auth.AuthWriter) *Client {
	c.url = url.URL{
		Host:   host,
		Path:   path,
		Scheme: scheme,
	}
	c.AuthWriter = authWriter
	return c
}

// SetQuery sets the query parameters for the Client's URL.
func (c *Client) SetQuery(path string, query any) *Client {
	c.params = query
	c.setPath(path)
	if query != nil {
		c.buildQueryString()
	}
	return c
}

// Get makes a GET request to the PTV API using the Client's URL and parameters.
func (c *Client) Get() ([]byte, error) {
	c.signURL()

	req, err := http.NewRequest("GET", c.url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}
	fmt.Println("Status code: ", resp.StatusCode)
	return io.ReadAll(resp.Body)
}

// setPath sets the path for the Client's URL.
// Calls setPathVariables.
// It iterates over the fields of the params struct, checking for "path" tags.
// Parses for placeholders in the path and replaces them with values from the params struct.
func (c *Client) setPath(path string) {
	c.url.Path = path
	err := c.setPathVariables()
	if err != nil {
		fmt.Println("Error setting path variables:", err)
	}
}

var flattenKeys = []string{
	"departure",
	"departures",
	"directions",
	"disruption",
	"disruptions",
	"outlets",
	"route",
	"route_types",
	"routes",
	"runs",
	"stop",
	"stops",
}

func flattenMap(m map[string]any, keys []string) {
	for _, k := range keys {
		if inner, ok := m[k].(map[string]any); ok {
			for ik, iv := range inner {
				m[ik] = iv
			}
			delete(m, k)
		}
	}
}

// buildQueryString constructs the query string from the params struct.
// It iterates over the fields of the struct, checking for "query" tags.
// Checks for nil or zero values to avoid adding empty query parameters.
func (c *Client) buildQueryString() {
	params := reflect.ValueOf(c.params)
	if params.Kind() == reflect.Ptr {
		params = params.Elem()
	}

	typ := params.Type()
	query := url.Values{}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := params.Field(i)

		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}

		if !value.IsValid() || value.IsZero() {
			continue
		}
		query.Add(tag, fmt.Sprintf("%v", value.Interface()))
	}
	c.url.RawQuery = query.Encode()
}

// setPathVariables replaces placeholders in the path with values from the params struct.
// It iterates over each to access the fields of the struct and their tags.
func (c *Client) setPathVariables() error {
	val := reflect.ValueOf(c.params)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("params must be a pointer to a struct, type passed: %v", val.Kind())
	}

	typ := val.Type()
	path := c.url.Path

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		tag := field.Tag.Get("path")

		if tag == "" {
			continue
		}
		placeholder := fmt.Sprintf("{%s}", tag)
		valueStr := fmt.Sprintf("%v", value.Interface())
		path = strings.ReplaceAll(path, placeholder, valueStr)
		fmt.Printf("Replacing %s with %s in path\n", placeholder, valueStr)
	}
	fmt.Printf("Final path after replacements: %s\n", path)
	c.url.Path = path
	return nil
}

// signURL generates a signed URL using the AuthWriter.
// It appends the developer ID and signature to the URL's query string.
func (c *Client) signURL() {
	c.url.RawQuery = c.AuthWriter.GenerateSignature(c.url).Query().Encode()
}
