package utils

import (
	"math"
	"strconv"
	"errors"
)
type PermissionTable struct{
	 Values map[string]int 
}

func CreatePermissionTable() (permissionTable *PermissionTable){
	permissionTable = &PermissionTable{nil}
	permissionTable.Values = make(map[string]int)
	permissionTable.Values["GET"] = 1
	permissionTable.Values["POST"] = 2
	permissionTable.Values["PUT"] = 4
	permissionTable.Values["DELETE"] = 8
	permissionTable.Values["CONNECT"] = 16
	permissionTable.Values["HEAD"] = 32
	permissionTable.Values["PATCH"] = 64
	permissionTable.Values["OPTION"] = 128
	permissionTable.Values["TRACE"] = 256
	return permissionTable
}
func (permissionTable PermissionTable) GetCodeNumber(permissions []string) int {
	var sum int = 0
	for i := 0; i < len(permissions); i++ {
		sum = sum + permissionTable.Values[permissions[i]]
	}
	return sum
}
func  (permissionTable PermissionTable) getKey(value int) (string,error){
	for k , v := range permissionTable.Values{
		if v == value{
			return k, nil
		}
	}
	return "", errors.New("no key for the given value found")
}

func (permissionTable PermissionTable) GetPermissions(codeNumber int) []string {
	
	var permissionsCodeStr string = strconv.FormatInt(int64(codeNumber), 2)
	list:= []rune(permissionsCodeStr)
	var permissions []string
	index:= 0
	for i := len(list)-1; i >= 0 ; i-- {
        item := string(list[i])
        if item == "1"{
			base := math.Pow(2,float64(index))
			key, _:= permissionTable.getKey(int(base))
			permissions=append(permissions,key)	
		}
		index++
    }
	return permissions
	}
	