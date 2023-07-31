package main

import (
	"embed"
	"log"
	"net"
	"net/http"
)

//go:embed static/html
var htmlEmbedFS embed.FS

//go:embed static/txt/*.txt
var txtEmbedFS embed.FS

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/static/html/", http.FileServer(http.FS(htmlEmbedFS)))
	http.Handle("/static/txt/", http.FileServer(http.FS(txtEmbedFS)))

	log.Fatal(http.Serve(listener, nil))
}
