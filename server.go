package main

import (
	"net/http"
)

func httpServer() error {

	// loadConfig() must be called before this point

	logger.Info("Starting server on:", getServerConfig())

	router := NewRouter()
	return http.ListenAndServe(getServerConfig(), router)

}
