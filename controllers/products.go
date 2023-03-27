package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/mohsenGoodarzi/webframework/models"
)

func products(w http.ResponseWriter, r *http.Request) {
	products := models.AllProducts()
	fmt.Fprintf(w, "%v", products)
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["productId"]
	fmt.Println(key)
	i, err := strconv.Atoi(key)
	if err != nil {
		return
	}
	product := models.GetProduct(i)
	fmt.Fprintf(w, "%v", product)
}

func ProductHandleRequests(route *mux.Router) {
	pro := route.PathPrefix("/products").Subrouter()
	pro.Path("/").HandlerFunc(products)
	pro.Path("/{productId}").HandlerFunc(product)
}
