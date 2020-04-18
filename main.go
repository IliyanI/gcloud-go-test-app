package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm listening"))
	})

	fmt.Printf("Listening on port %s: ", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}
