package route

import (
	"log"
	"net/url"
	"container/list"
)
 

type Route struct{
	url.URL
	Endpoint *list.List
}

func (route Route )NewRoute(scope string){
}

func (route Route) validate(urlString string) *url.URL{
	u, err := route.URL.Parse(urlString)
	if err != nil{
		log.Fatal("The given string is not url")
	}
	return u
}

func (route Route)AddEndPoint(endpoint string){
	route.Endpoint.PushBack(endpoint)
}



