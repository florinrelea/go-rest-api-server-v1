package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = 9003

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)

	fmt.Printf("Server started and listening on port %d\n", serverPort)

	var serverAddress = fmt.Sprintf(":%d", serverPort)

	log.Fatal(http.ListenAndServe(serverAddress , nil))
}