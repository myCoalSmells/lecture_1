package handlers

import (
	"lecture-1/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle POST (update)

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("client POSTed new product")

	prod := &data.Product{}
	err := prod.FromJSONToProduct(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}
