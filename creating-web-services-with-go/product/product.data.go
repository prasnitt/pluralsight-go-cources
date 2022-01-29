package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// used to hold our product list in memory
var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func Init(fileName string) {
	fmt.Println("loading products...")
	prodMap, err := loadProductMap(fileName)
	productMap.m = prodMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d products loaded...\n", len(productMap.m))
}

func loadProductMap(fileName string) (map[int]Product, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file [%s]", fileName)
	}
	products := []Product{}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]Product)
	for i := 0; i < len(products); i++ {
		prodMap[products[i].ProductID] = products[i]
	}
	return prodMap, nil
}

func product(id int) *Product {
	productMap.RLock()
	defer productMap.RUnlock()
	if p, ok := productMap.m[id]; ok {
		return &p
	}
	return nil
}

func remove(id int) {
	productMap.Lock()
	defer productMap.Unlock()
	delete(productMap.m, id)
}

func products() []Product {
	productMap.RLock()
	defer productMap.RUnlock()
	pp := []Product{}
	for _, p := range productMap.m {
		pp = append(pp, p)
	}
	return pp
}

func productIds() []int {
	productMap.RLock()
	defer productMap.RUnlock()
	ids := []int{}
	for id := range productMap.m {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	return ids
}

func nextId() int {
	return len(productMap.m) + 1
}

func addOrUpdate(p Product) (int, error) {
	productId := -1

	if p.ProductID > 0 {
		oldProduct := product(p.ProductID)

		if oldProduct == nil {
			return 0, fmt.Errorf("product id [%d] does'nt exists", p.ProductID)
		}
		productId = p.ProductID
	} else {
		productId = nextId()
		p.ProductID = productId
	}

	productMap.Lock()
	defer productMap.Unlock()
	productMap.m[productId] = p

	return productId, nil
}
