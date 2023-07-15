package main

import (
	"Publisher/activity"
	"Publisher/banner"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var INETAddress string = "0.0.0.0"
	var PublisherPortAddress string = ":65110"

	// Publisher Services Runs on 65110
	// Subscriber Services Runs on 65111
	// Repository Service Runs on 65123                  // INTERNAL USAGE ONLY
	// Persistent Store Service [Database] Runs on 65124 // INTERNAL USAGE ONLY
	// AUTHORIZATION Service Runs on 65420
	// Infrastructure Service Runs on 60001              // INTERNAL USAGE ONLY

	banner.StartBanner()

	// starting publisher service, registering routes and etc

	router := mux.NewRouter()
	router.HandleFunc("/all", activity.GetAllPublisherResponse)

	var ServiceAddress string = INETAddress + PublisherPortAddress
	http.ListenAndServe(ServiceAddress, router)
}
