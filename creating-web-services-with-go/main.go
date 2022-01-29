package main

import (
	"log"
	"net/http"

	"github.com/prasnitt/go/inventoryservice/product"
	"github.com/prasnitt/go/inventoryservice/receipt"
)

const basePath = "/api"
const port = ":5000"

func main() {
	product.Init("products.json")
	product.SetupRoutes(basePath)
	receipt.SetupRoutes(basePath)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
