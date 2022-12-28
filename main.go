package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HalloHandler)
	mux.HandleFunc("/list", handler.BukuHandler)
	mux.HandleFunc("/buku", handler.ListbukuHandler)
	log.Println(("starting web di port 8080"))
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
