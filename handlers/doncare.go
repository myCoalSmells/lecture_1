package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type DonCare struct {
	l *log.Logger
}

func NewDonCare(l *log.Logger) *DonCare {
	return &DonCare{l}
}

func (dc *DonCare) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	dc.l.Printf("Received message: %s\n", d)
	fmt.Fprintf(rw, "We don't care")
}
