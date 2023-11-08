package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body) // reads the request, r, puts it in d
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}
		log.Printf("Data Received: %s\n", d)
		fmt.Fprintf(rw, "Data returned: %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	log.Println("Listening on port 9090...")
	http.ListenAndServe(":9090", nil)
}
