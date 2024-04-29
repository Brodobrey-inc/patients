package main

import (
	"patients/database"
	"patients/logging"
	"patients/webserver"
)

func main() {
	logging.Initialize("debug")

	err := database.Initialize("./data/list_patients.json")
	if err != nil {
		logging.LogFatalError(err, "Failed initialize data")
	}

	r := webserver.Initialize()
	if err := r.Run(); err != nil {
		logging.LogError(err, "Failed to start webserver")
	}
}
