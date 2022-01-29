package main

import (
	"log"
	"net/http"

	"github.com/prasnitt/go/inventoryservice/product"
)

const basePath = "/api"

func main() {
	product.Init("products.json")
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
