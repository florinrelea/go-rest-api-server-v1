package backend

import (
	"fmt"
	"log"
	"net/http"
)

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func Run(addr string) {
	http.HandleFunc("/", reqHandler)

	fmt.Println("Server is running on port", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
