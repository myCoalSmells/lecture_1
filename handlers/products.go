package handlers

import (
	"lecture-1/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	// handle POST (add)

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// handle PUT (update)
	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)

		p.updateProduct(id, rw, r)
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

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("client PUT wants to update product")

	prod := &data.Product{}
	err := prod.FromJSONToProduct(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	e := data.UpdateProduct(id, prod)
	if e != nil {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
}
