package main

import (
	"fmt"
)

func main() {
	// Publisher Services Runs on 65110
	// Subscriber Services Runs on 65111
	// Repository Service Runs on 65123                  // INTERNAL USAGE ONLY
	// Persistent Store Service [Database] Runs on 65124 // INTERNAL USAGE ONLY
	// AUTHORIZATION Service Runs on 65420
	// Infrastructure Sercive Runs on 60001              // INTERNAL USAGE ONLY

	fmt.Println("ENTRYPOINT: RUNNING ON 65108")
	fmt.Println("This Service has only HELP Endpoint. GET /help to get all informations")

	fmt.Println("********** OTHER SERVICES **********")
	fmt.Println("ENTRYPOINT: Publisher Services Runs on 65110")
	fmt.Println("ENTRYPOINT: Subscriber Services Runs on 65111")
	fmt.Println("ENTRYPOINT: AUTHORIZATION Service Runs on 65420")
}
