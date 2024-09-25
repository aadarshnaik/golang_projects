package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpint Hit: return all products")
	json.NewEncoder(w).Encode(Products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getProduct)
	log.Fatal(http.ListenAndServe(":3040", myRouter))
}

func main() {
	Products = []Product{
		Product{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		Product{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},
	}
	handleRequests()
}
