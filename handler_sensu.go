package main

import (
	"fmt"
	"github.com/gbolo/go-monapi/lib/sensuapi"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

const (
	SensuApiError = "<span class='badge badge-pill badge-danger'>Sensu API Error</span>\n"
	SensuNoChecks = "<span class='badge badge-pill badge-default'>No Checks Found</span>\n"
)

var (
	SensuApiClient *sensuapi.API
)

func SensuChecksToHtml(w http.ResponseWriter, r *http.Request) {

	status := make(map[int]string)
	status[0] = "success"
	status[1] = "warning"
	status[2] = "danger"

	// needed for ajax
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)

	// get sensu agent name
	clientName := vars["clientName"]
	logger.Debug("Sensu clientName:", clientName)

	// get api client
	c := SensuApiClient

	// get all check results for specific sensu agent
	var results []sensuapi.Result
	if _, err := c.GetResultsByClient(&results, clientName); err != nil {
		// still send 200, but with SensuApiError
		w.WriteHeader(http.StatusOK)
		logger.Error(err)
		fmt.Fprint(w, SensuApiError)
		return
	}

	// combine all checks
	checks := make([]sensuapi.Check, 0, 3)
	for _, result := range results {
		checks = append(checks, result.Check)
	}

	if len(checks) > 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, ChecksToHtml(checks))
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, SensuNoChecks)
	}

}

func InitSensuApiClient() {

	viper.SetDefault("sensu.api_address", "127.0.0.1:4576")
	viper.SetDefault("sensu.api_scheme", "http")
	viper.SetDefault("sensu.api_username", "")
	viper.SetDefault("sensu.api_password", "")

	config := sensuapi.DefaultConfig()
	config.Address = viper.GetString("sensu.api_address")
	config.Scheme = viper.GetString("sensu.api_scheme")
	config.Username = viper.GetString("sensu.api_username")
	config.Password = viper.GetString("sensu.api_password")

	var err error
	SensuApiClient, err = sensuapi.NewAPIClient(config)
	if err != nil {
		panic(err)
	}

	logger.Infof(
		"Sensu API Config: %s://%s %s\n",
		config.Scheme,
		config.Address,
		config.Username,
	)

}

func ChecksToHtml(checks []sensuapi.Check) string {

	status := make(map[int]string)
	status[0] = "success"
	status[1] = "warning"
	status[2] = "danger"

	ss := []string{}
	for _, check := range checks {

		class := "info"
		if check.Status >= 0 && check.Status <= 2 {
			class = status[check.Status]
		}

		ss = append(ss,
			fmt.Sprintf("<span class='badge badge-pill badge-%s'>%s</span>\n",
				class,
				check.Name,
			),
		)

	}

	return strings.Join(ss, "")
}
