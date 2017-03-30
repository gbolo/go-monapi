package sensuapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)


// API Client is used as a handle for all client methods
type API struct {
	config Config
}


// DefaultConfig sets up a default configuration struct
func DefaultConfig() *Config {
	config := &Config{
		Scheme:     "http",
		Address:    "127.0.0.1:4567",
		HTTPClient: http.DefaultClient,
	}
	return config
}

// NewAPIClient gets a new Sensu API client
func NewAPIClient(config *Config) (*API, error) {
	defConfig := DefaultConfig()

	if len(config.Scheme) == 0 {
		config.Scheme = defConfig.Scheme
	}

	if len(config.Address) == 0 {
		config.Address = defConfig.Address
	}

	if config.HTTPClient == nil {
		config.HTTPClient = defConfig.HTTPClient
	}

	apiClient := &API{
		config: *config,
	}

	return apiClient, nil
}

// Build a http request
func (c *API) buildRequest(method, path string) (*http.Request, error) {
	url := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.Address,
		Path:   path,
	}

	req, err := http.NewRequest(method, url.String(), nil)

	if c.config.Username != "" && c.config.Password != "" {
		req.SetBasicAuth(c.config.Username, c.config.Password)
	}

	return req, err
}

// Send the request to the server
func (c *API) doRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.config.HTTPClient.Do(req)
	return resp, err
}

// Decode JSON payload
func jsonDecode(out interface{}, data io.ReadCloser) error {
	d := json.NewDecoder(data)
	error := d.Decode(out)
	return error
}

// Generic GET request. Decoded JSON is set in the out interface{} passed in.
func (c *API) get(uri string, out interface{}) (*http.Response, error) {
	request, _ := c.buildRequest("GET", uri)
	resp, err := c.doRequest(request)

	if err != nil {
		return nil, err
	}

	if out != nil {
		err = jsonDecode(out, resp.Body)
	}

	return resp, err
}