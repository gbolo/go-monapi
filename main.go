package main

// Constants go here.
const (
	configName   = "config"
	envVarPrefix = "MONAPI"
	appName      = "monapi"
	version      = "v0.1"
)

// Start main()
func main() {

	initLogging()
	loadConfig()
	InitSensuApiClient()

	logger.Info("[STARTUP] Initializing", appName, version)
	logger.Fatal(httpServer())
}
