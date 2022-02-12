package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"

	"os"

	"github.com/caddyserver/certmagic"
)

var (
	verbose   = flag.Bool("verbose", true, "verbose logging")
	bind      = flag.String("bind", ":443", "address to listen on")
	email     = flag.String("email", "r2proxy@mailinator.com", "email address for LetsEncrypt")
	hostname  = flag.String("hostname", "localhost", "hostname for default server")
	serveHttp = flag.Bool("http", false, "also handle http on port 80")
	staging   = flag.Bool("staging", false, "use LetsEncrypt staging server")
	target    = flag.String("target", "localhost", "target (=proxied server hostname)")

	logger = log.New(os.Stdout, "R2PROXY: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

func main() {

	flag.Parse()

	if *verbose {
		logger.Printf("INFO: https listening on %s\n", *bind)
		logger.Printf("INFO: local hostname is %s\n", *hostname)
		logger.Printf("INFO: proxy target is %s\n", *target)
	}

	director := func(req *http.Request) {

		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", *target)
		//LATER: add X-Forwarded-For
		req.URL.Scheme = "http"
		req.URL.Host = *target
		req.Host = *target
	}

	proxy := &httputil.ReverseProxy{Director: director}

	mux := http.NewServeMux()
	mux.Handle(*hostname+"/status.json", &Status{})

	theStaticHandler := staticHandler()
	mux.Handle(*hostname+"/robots.txt", theStaticHandler)
	mux.Handle(*hostname+"/css/pico.min.css", theStaticHandler)
	mux.Handle(*hostname+"/favicon.ico", theStaticHandler)
	mux.Handle(*hostname+"/favicon.svg", theStaticHandler)
	mux.Handle(*hostname+"/", theStaticHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("INFO: proxying %s (%s)\n", r.URL, r.Host)
		proxy.ServeHTTP(w, r)
	})

	if *serveHttp {
		logger.Printf("INFO: proxying http on port 80")
		go http.ListenAndServe(":80", mux)
	}

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
		Addr:      *bind,
		Handler:   mux,
		TLSConfig: magic.TLSConfig(),
	}

	log.Fatal(httpsServer.ListenAndServeTLS("", ""))
}
