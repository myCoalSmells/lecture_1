package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := io.ReadAll(r.Body) // reads the request, r, puts it in d
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data Received: %s\n", d)
	fmt.Fprintf(rw, "Data returned: %s\n", d)

}
