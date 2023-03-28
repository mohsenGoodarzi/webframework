package main

import (
	"github.com/mohsenGoodarzi/webframework/entities/acl"
)


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
