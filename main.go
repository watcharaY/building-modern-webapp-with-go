package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
