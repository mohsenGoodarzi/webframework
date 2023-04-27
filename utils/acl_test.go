package utils

import (
	"testing"
)
func TestGetCodeNumber( t *testing.T) {
	
	myAcl := CreateACL()
	myAcl.EndPoints["One"]= myAcl.PermissionTable.GetCodeNumber([]string{"GET"})
	myAcl.EndPoints["Some"]= myAcl.PermissionTable.GetCodeNumber([]string{"PATCH","PUT","CONNECT"})
	myAcl.EndPoints["All"]= myAcl.PermissionTable.GetCodeNumber([]string{"PATCH","GET","POST","OPTION","TRACE","HEAD","CONNECT","DELETE","PUT"})
   expected:= make(map[string]int)
   expected["One"]=1
   expected["Some"]=84
   expected["All"]=511
	for k,v:= range myAcl.EndPoints{
	if  v!=expected[k]{
		t.Errorf("got %q, wanted %q", v, expected[k])
	}
	}	
    }
	func TestGetPermissions(t *testing.T) {
	
		myAcl := CreateACL()
		myAcl.EndPoints["One"]= myAcl.PermissionTable.GetCodeNumber([]string{"GET"})
		myAcl.EndPoints["Some"]= myAcl.PermissionTable.GetCodeNumber([]string{"PATCH","PUT","CONNECT"})
		myAcl.EndPoints["All"]= myAcl.PermissionTable.GetCodeNumber([]string{"PATCH","GET","POST","OPTION","TRACE","HEAD","CONNECT","DELETE","PUT"})
	   got:= [][]string{
		myAcl.PermissionTable.GetPermissions(myAcl.EndPoints["One"]),
		myAcl.PermissionTable.GetPermissions(myAcl.EndPoints["Some"]),
		myAcl.PermissionTable.GetPermissions(myAcl.EndPoints["All"]),
	}
	expected:=[][]string{
		{"GET"},
		{"PUT","CONNECT","PATCH"},
		{"GET","POST","PUT","DELETE", "CONNECT","HEAD", "PATCH","OPTION","TRACE"},
	}
		for i:=0; i<len(expected);i++{
			for j:=0; j<len(expected[i]);j++{
				if  got[i][j]!=expected[i][j]{
					t.Errorf("got %q, wanted %q", got[i][j], expected[i][j])
				}
			}
		}	
		}