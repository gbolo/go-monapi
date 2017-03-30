package sensuapi

import (
	"net/http"
	"strings"
)


// ClientsURI Constant
const clientsURI string = "/clients"

// GetClients gets all clients
// sets 'out' as an array of Client structs as JSON
func (c *API) GetClients(out interface{}) (*http.Response, error) {
	resp, err := c.get(clientsURI, out)
	return resp, err
}

// GetClient gets an individual client
// sets 'out' as a Client struct as JSON
func (c *API) GetClient(out interface{}, clientName string) (*http.Response, error) {
	s := []string{clientsURI, clientName}
	uri := strings.Join(s, "/")
	resp, err := c.get(uri, out)
	return resp, err
}
