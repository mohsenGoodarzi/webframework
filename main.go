package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
	//"github.com/mohsenGoodarzi/webframework/services"
	acl "github.com/mohsenGoodarzi/webframework/utils"
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
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func main() {
	initServer()
	
}
func main_() {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
