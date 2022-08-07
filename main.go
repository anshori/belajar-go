package main

import (
	"belajar/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HandlerIndex)
	http.HandleFunc("/api", handler.HandlerIndex)
	http.HandleFunc("/gallery", handler.HandlerGallery)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	address := "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
