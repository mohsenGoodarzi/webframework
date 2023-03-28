package persistentlayer

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/mohsenGoodarzi/webframework/entities/acl"
)

func GetACL(id string) *acl.ACL {
	
	return *acl.ACL
}