package banner

import (
	"fmt"
	"io/ioutil"
	"log"
)

func StartBanner() {
	banner := "./banner/banner.txt"
	content, err := ioutil.ReadFile(banner)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))

	fmt.Println("\nWELCOME TO PUBLISHER SERVICES")
	fmt.Println("\n\n********** OTHER SERVICES **********")
	fmt.Println("ENTRYPOINT: Subscriber Services Runs on 65111")
	fmt.Println("ENTRYPOINT: AUTHORIZATION Service Runs on 65420")
}
