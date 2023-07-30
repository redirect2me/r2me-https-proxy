package main

import (
	"flag"
	"log"
	//"net/http"
	//"net/http/httputil"

	"os"

	//"github.com/caddyserver/certmagic"
)

var (
	verbose   = flag.Bool("verbose", true, "verbose logging")
	bind      = flag.String("bind", ":443", "address to listen on")
	email     = flag.String("email", "r2proxy@mailinator.com", "email address for LetsEncrypt")
	hostname  = flag.String("hostname", "localhost", "hostname for default server")
	bindHttp = flag.String("http", "none", "bind address for handling plain http (or \"none\")")
	staging   = flag.Bool("staging", false, "use LetsEncrypt staging server")
	target    = flag.String("target", "localhost", "target (=proxied server hostname)")

	logger = log.New(os.Stdout, "R2PROXY: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

func main() {

	flag.Parse()

	if *verbose {
		logger.Printf("xINFO: https listening on %s\n", *bind)
		logger.Printf("INFO: local hostname is %s\n", *hostname)
		logger.Printf("INFO: proxy target is %s\n", *target)
		logger.Printf("INFO: http handling is %s\n", *bindHttp)
	}

	var done = make(chan bool)

	mux := getAssetMux(*hostname)
	mux.Handle("/", getProxyHandler(*target))

	go serveHttps(*bind, mux)
	go serveHttp(*bindHttp, mux)

	<-done

	logger.Printf("INFO: done")
}
