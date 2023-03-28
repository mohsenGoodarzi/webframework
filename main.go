package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	//"github.com/mohsenGoodarzi/webframework/services"
	"github.com/mohsenGoodarzi/webframework/entities/acl"
)

func registerAllEndPoints(r *mux.Router){
	//services.ProductHandleRequests(r)
}
func initServer() {

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
	//initServer()
	endpoint := "About"
	myAcl := acl.CreateACL()
	myAcl.EndPoints["Product"]=345
	myAcl.EndPoints["About"]=450
	myAcl.EndPoints["Remove/ACL"]=450
	for i := 0; i < myAcl.EndPoints.Len(); i++ {
		elem := myAcl.EndPoints.Front()
		if elem.Value == endpoint {
			return elem
		}
    }
}
