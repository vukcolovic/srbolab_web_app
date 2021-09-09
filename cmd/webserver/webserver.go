package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {

	appConfig, err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	// create http server
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port),
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	log.Printf("Starting HTTP server on port %s....", appConfig.Port)

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
