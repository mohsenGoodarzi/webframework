package utils

const (
	NOTKEY    = "not found"
	NOELEMENT = -1
)

type ACLs struct {
	Acls            *[]ACL
	PermissionTable *PermissionTable
}

type ACL struct {
	EndPoints       map[string]int
	PermissionTable *PermissionTable
}

func CreateACL() (acl *ACL) {
	acl = &ACL{
		EndPoints:       make(map[string]int),
		PermissionTable: CreatePermissionTable(),
	}
	return acl
}
func (acl ACL) CalcAclValue(requestType []string) int {
	return acl.PermissionTable.GetCodeNumber(requestType)
}

func (acl ACL) AddAcl(key string, value int) {
	acl.EndPoints[key] = value
}

func (acl ACL) RemoveAcl(endpoint string) {
	delete(acl.EndPoints, endpoint)
}

func (acl ACL) FindAcl(endpoint string) (string, int) {
	for key, element := range acl.EndPoints {
		if key == endpoint {
			return key, element
		}
	}
	return NOTKEY, NOELEMENT
}
