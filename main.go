package main

import (
	"flag"
	//	"fmt"
	"log"
	//	"net"
	"net/http"
	"net/http/httputil"

	//    "net/url"
	"os"
	//    "strconv"
	//	"strings"
	//	"sync"
	//	"sync/atomic"
	//	"time"
)

var (
	verbose   = flag.Bool("verbose", true, "verbose logging")
	bind      = flag.String("bind", ":443", "address to listen on")
	email     = flag.String("email", "r2proxy@mailinator.com", "email address for LetsEncrypt")
	hostname  = flag.String("hostname", "localhost", "hostname for default server")
	serveHttp = flag.Bool("http", false, "also handle http on port 80")
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
		go http.ListenAndServe(":80", mux)
	}

	log.Fatal(http.ListenAndServe(*bind, mux))
}
