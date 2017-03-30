package main

import (
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"time"
)

var logger *logging.Logger
var format = logging.MustStringFormatter(
	`%{time:2006/01/02 15:04:05} %{shortfile} [%{level:.4s}](%{id:03x}) %{message}`,
)

func logSensitive(m string) error {

	viper.SetDefault("logging.private", "private.log")

	f, err := os.OpenFile(viper.GetString("logging.private"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		logger.Error("Error opening private log file:", viper.GetString("logging.private"))
	} else {
		// assign to standard logger
		log.SetOutput(f)

		// log the message
		log.Println(m)
	}
	defer f.Close()

	return err
}

func initLogging() {
	// set up logger
	logger = logging.MustGetLogger(appName)
	generalLog := logging.NewLogBackend(os.Stdout, "", 0)
	generalLogBackEnd := logging.NewBackendFormatter(generalLog, format)

	logging.SetBackend(generalLogBackEnd)
}

func AccessLogger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Infof(
			"%s %s %s %s %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
