package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func allProducts(w http.ResponseWriter, r *http.Request) {

}

func getProduct(w http.ResponseWriter, r *http.Request) {

}
func handleRequests(route *mux.Route) {
	products := route.PathPrefix("/products").Subrouter()
	products.Path("/").HandlerFunc(allProducts)
	products.Path("/{productId}").HandlerFunc(getProduct)
}
