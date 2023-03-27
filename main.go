package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/mohsenGoodarzi/webframework/controllers"
	"github.com/gorilla/mux"
)


func handleRequests() {
	

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
