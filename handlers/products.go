package handlers

import (
	"lecture-1/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// want to make a GET request to serve.http and get our product list. productlist must be converted to json to be sent
	p.l.Println("client GET requested product list")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	// 	rw.Write(d)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("client POSTed new product")

	prod := &data.Product{}
	err := prod.FromJSONToProduct(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "bad id", http.StatusBadRequest)
	}
	p.l.Println("client PUT wants to update product", id)

	prod := &data.Product{}
	e2 := prod.FromJSONToProduct(r.Body)

	if e2 != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	e := data.UpdateProduct(id, prod)
	if e != nil {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
}
