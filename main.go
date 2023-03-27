package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mohsenGoodarzi/webframework/controllers"
)
func registerAllEndPoints(r *mux.Router){
	controllers.ProductHandleRequests(r)
}
func handleRequest() {

	r := mux.NewRouter()
	r.Host("http://localhost:8000/")
	r.Schemes("http://")
	r.Methods("GET , POST")
	registerAllEndPoints(r)
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
	handleRequest()
}
