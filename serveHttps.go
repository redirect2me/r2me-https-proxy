package main

import (
	"log"
	"net/http"
	
	"github.com/caddyserver/certmagic"
)

func serveHttps(bind string, mux *http.ServeMux) {

	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = *email
	if *staging {
		certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	} else {
		certmagic.DefaultACME.CA = certmagic.LetsEncryptProductionCA
	}
	certmagic.Default.OnDemand = &certmagic.OnDemandConfig{
		DecisionFunc: func(name string) error {
			//LATER: DNS check
			//LATER: algorithm check
			return nil
		},
	}

	magic := certmagic.NewDefault()
	httpsServer := &http.Server{
		Addr:      bind,
		Handler:   mux,
		TLSConfig: magic.TLSConfig(),
	}

	log.Fatal(httpsServer.ListenAndServeTLS("", ""))
}
