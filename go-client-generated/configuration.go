/*
 * Prometheus HTTP API
 *
 * The current stable HTTP API is reachable under /api/v1 on a Prometheus server. Any non-breaking additions will be added under that endpoint.  # Format overview The API response format is JSON. Every successful API request returns a ```2xx``` status code.  Invalid requests that reach the API handlers return a JSON error object and one of the following HTTP response codes:  ```400 Bad Request``` when parameters are missing or incorrect. ```422 Unprocessable Entity``` when an expression can't be executed ([RFC4918](https://datatracker.ietf.org/doc/html/rfc4918#page-78)). ```503 Service Unavailable``` when queries time out or abort.  Other non-```2xx``` codes may be returned for errors occurring before the API endpoint is reached.  An array of warnings may be returned if there are errors that do not inhibit the request execution. All of the data that was successfully collected will be returned in the data field.  The JSON response envelope format is as follows:  ``` {   \"status\": \"success\" | \"error\",   \"data\": <data>,    // Only set if status is \"error\". The data field may still hold   // additional data.   \"errorType\": \"<string>\",   \"error\": \"<string>\",    // Only if there were warnings while executing the request.   // There will still be data in the data field.   \"warnings\": [\"<string>\"] } ``` # Generic placeholders:  ```<rfc3339 | unix_timestamp>```: Input timestamps may be provided either in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format or as a Unix timestamp in seconds, with optional decimal places for sub-second precision. Output timestamps are always represented as Unix timestamps in seconds.  ```<series_selector>```: Prometheus [time series selectors](https://prometheus.io/docs/prometheus/latest/querying/basics/#time-series-selectors) like ```http_requests_total``` or ```http_requests_total{method=~\"(GET|POST)\"}``` and need to be URL-encoded.  ```<duration>```: [Prometheus duration strings](https://prometheus.io/docs/prometheus/latest/querying/basics/#time_durations). For example, ```5m``` refers to a duration of 5 minutes.  ```<bool>```: boolean values (strings ```true``` and ```false```).  **Note**: Names of query parameters that may be repeated end with ```[]```. 
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
