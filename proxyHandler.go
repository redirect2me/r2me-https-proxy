package main

import (
	"net/http"
	"net/http/httputil"
)

func getProxyHandler(target string) http.HandlerFunc {
	
	director := func(req *http.Request) {

		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", target)
		//LATER: add X-Forwarded-For
		req.URL.Scheme = "http"
		req.URL.Host = target
		req.Host = target
	}

	proxy := &httputil.ReverseProxy{Director: director}

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("INFO: proxying %s (%s)\n", r.URL, r.Host)
		proxy.ServeHTTP(w, r)
	}
}

