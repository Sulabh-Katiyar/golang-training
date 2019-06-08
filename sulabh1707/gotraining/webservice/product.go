package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The person Type (more like an object)
type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
	// Address   *Address `json:"address,omitempty"`
}

// type Address struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }

var prod []Product

// Display all from the people var
func GetProd(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(prod)
}

// Display a single data
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range prod {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

// create a new item
func AddProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	prod = append(prod, product)
	json.NewEncoder(w).Encode(prod)
}

// Delete an item
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range prod {
		if item.ID == params["id"] {
			prod = append(prod[:index], prod[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(prod)
	}
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	prod = append(prod, Product{ID: "1", Name: "Mango", Price: "Rs.60"})
	prod = append(prod, Product{ID: "2", Name: "Banana", Price: "Rs.50"})
	router.HandleFunc("/prod", GetProd).Methods("GET")
	router.HandleFunc("/prod/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/prod/{id}", AddProduct).Methods("POST")
	router.HandleFunc("/prod/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
