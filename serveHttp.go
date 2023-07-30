package main

import (
	"log"
	"net/http"
)

func serveHttp(bind string, mux *http.ServeMux) {
	if bind == "none" {
		logger.Fatal("INFO: http not handled")
	} else {
		logger.Printf("INFO: proxying http on %s", bind)
		log.Fatal(http.ListenAndServe(bind, mux))
	}
}
