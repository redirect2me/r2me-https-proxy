package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed assets
var embeddedFiles embed.FS

func staticHandler() http.Handler {

	fsys, err := fs.Sub(embeddedFiles, "assets")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(fsys))
}

func getAssetMux(hostname string) *http.ServeMux {

	mux := http.NewServeMux()
	mux.Handle(hostname+"/status.json", &Status{})

	theStaticHandler := staticHandler()
	mux.Handle(hostname+"/robots.txt", theStaticHandler)
	mux.Handle(hostname+"/css/pico.min.css", theStaticHandler)
	mux.Handle(hostname+"/favicon.ico", theStaticHandler)
	mux.Handle(hostname+"/favicon.svg", theStaticHandler)
	mux.Handle(hostname+"/", theStaticHandler)

	return mux;
}
