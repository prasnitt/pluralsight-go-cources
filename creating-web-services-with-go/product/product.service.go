package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const productsPath = "products"

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products := products()
		err := json.NewEncoder(w).Encode(products)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var p Product
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println("Err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdate(p)
		if err != nil {
			log.Println("Err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", productsPath))

	if len(urlPathSegments) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		p := product(id)
		if p == nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(p)

		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPut:
		p := Product{}
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println("Err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if p.ProductID != id || id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = addOrUpdate(p)
		if err != nil {
			log.Println("Err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		remove(id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func SetupRoutes(apiBasePath string) {
	productsHandler := http.HandlerFunc(handleProducts)
	productHandler := http.HandlerFunc(handleProduct)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), productsHandler)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), productHandler)
}
