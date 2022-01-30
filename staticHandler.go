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
