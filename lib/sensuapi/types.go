package sensuapi

import "net/http"

// used for connection to sensu api

type Config struct {
	Address    string
	Scheme     string
	Username   string
	Password   string
	HTTPClient *http.Client
}

// used for /clients endpoint

type Client struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Subscriptions []string `json:"subscriptions"`
	Version       string   `json:"version"`
	Timestamp     int      `json:"timestamp"`
	Environment   string   `json:"environment"`
	Socket        Csocket  `json:"socket,omitempty"`
}

type Csocket struct {
	Port	      int       `json:"port"`
	Bind	      string    `json:"bind"`
}


// used for /result endpoint

type Check struct {
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	Interval    int      `json:"interval"`
	Issued      int      `json:"issued"`
	Executed    int      `json:"executed"`
	Output      string   `json:"output"`
	Status      int      `json:"status"`
	Duration    float64  `json:"duration"`
	Type	    string   `json:"type"`
	Standalone  bool     `json:"standalone"`
}

type Result struct {
	ClientName string `json:"client"`
	Check      Check  `json:"check"`
}