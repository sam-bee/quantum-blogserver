package quantumblogserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func RunWebserver(conf ApplicationConfig) {

	// Create a file server handler for serving static files
	fileServer := http.FileServer(http.Dir(conf.getStaticSiteDir()))

	// Register the file server handler to handle all requests
	http.Handle("/", fileServer)

	// Start the HTTPS server
	log.Println("HTTPS server started...")
	port := ":" + strconv.Itoa(conf.getPortNo())
	log.Fatal(http.ListenAndServeTLS(port, conf.getCertPem(), conf.getKeyPem(), nil))

	fmt.Print("Server started")
}
