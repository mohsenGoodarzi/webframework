package entities

const (
	NOTKEY = "not found"
	NOELEMENT = -1
)
type ACLs struct {
	Acls *[]ACL
}

type ACL struct {
	EndPoints map[string]int
}

func CreateACL() (result *ACL){
 result.EndPoints = make(map[string]int)
 return result
}
func (acl ACL)CalcACLValue(requestType []string) int{
	to := len(requestType)
	for i:=0; i< to; i++{
		reqtype := requestType[i]
	}
} 

func (acl ACL)getACLValue(reqType string) int{
	for i := 0; i < acl.EndPoints.Len(); i++ {
		elem := acl.EndPoints.Front()
		if elem.Value == endpoint {
			return elem.key
		}
} 

func (acl ACL) AddACL(endpoint string, value int) {
	acl.EndPoints[endpoint] = value
}

func (acl ACL) RemoveACL(endpoint string) {
	delete(acl.EndPoints, endpoint)
}

func (acl ACL)FindACL(endpoint string) (string,int){
	for key, element := range acl.EndPoints {
       if key == endpoint {
		return key, element
	   }
    }
return NOTKEY, NOELEMENT
}
