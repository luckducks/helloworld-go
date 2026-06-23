package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var static embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(static)))
	http.ListenAndServe(":8080", nil)
}
