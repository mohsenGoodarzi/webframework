package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us Page!")
	fmt.Println("Endpoint Hit: AboutPage")
}
func homePage(w http.ResponseWriter, r *http.Request) {
	// writes into response
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	
	r.Path("/home").HandlerFunc(homePage).Schemes("http")
	r.Path("/about").HandlerFunc(aboutPage)

	//products:=r.PathPrefix("/products").Subrouter()
	//products.Path("/").HandlerFunc(allProducts)
	//products.Path("/{productId}").HandlerFunc(getProduct)

	url, _ := r.NewRoute().URLHost()
	fmt.Println(url)
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func main() {
	handleRequests()
}
