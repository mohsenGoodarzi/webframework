package persistent

import (
	"database/sql"
	"errors"
	"strings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mohsenGoodarzi/webframework/entities"
	"github.com/mohsenGoodarzi/webframework/utils"
)

type UserPersistant struct {
	DB *sql.DB
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/ACLDB")
	if err != nil {
		return nil, errors.New("connection failed due to " + err.Error())
	}
	return db, nil
}
func Init() (*UserPersistant, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	return &UserPersistant{
		DB: db,
	}, nil

}
func createAclInsertCommand(){

} 
func (up UserPersistant) Inset(user entities.User) (utils.DBOperationResult, error) {
	userResult, userErr := up.DB.Exec("INSERT INTO User (uesr_name,password)values(?,?)", user.Uid, user.Password)
	if userErr != nil {
		return utils.FAILED, userErr
	}
	uid, _ := userResult.LastInsertId()
	// not 
	var sb strings.Builder
	sb.WriteString("INSERT INTO Permission (uid,endpoint,permision)values")
	
	for endpoint,permission := range user.Acl.EndPoints{
		s := fmt.Sprintf("(%s,%s,%d),", "33", "ABOUT", 15)
		sb.WriteString(s)
		
	}
	result := sb.String()
	result = result[0:len]
	_, permissionErr := up.DB.Exec("INSERT INTO Permission (uid,endpoint,permision)values(?,?,?)", uid, endpoint, permission)
		if permissionErr != nil {
			return utils.FAILED, permissionErr
		}
	
	
	pr, _ := permissionResult.RowsAffected()
	if ur > 0 && pr > 0 {
		return utils.SUCCSSED, nil
	}
	return utils.NO_EFFECT, nil
}

func (up UserPersistant) Find(uesr_name string) (*entities.User, error) {
	user := &entities.User{
		Uid:      "",
		Password: "",
		Acl:      nil,
	}
	err := up.DB.QueryRow("SELECT uesr_name,password,permission FROM user where uesr_name = ?", uesr_name).Scan(&user.Uid, &user.Password, &user.Acl)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (up UserPersistant) Update(user entities.User) (int, error) {
	result, err := up.DB.Exec("update User set password = ?,permission = ? where uesr_name = ?", user.Password, user.Acl, user.Uid)
	if err != nil {
		return consts.OPERATION_FAILED, err
	}
	res, res_err := result.RowsAffected()
	if res_err != nil {
		return consts.OPERATION_FAILED, res_err
	}
	if res > 0 {
		return consts.OPERATION_SUCCSSED, nil
	}
	return consts.OPERATION_NO_EFFECT, nil
}

func (up UserPersistant) Delete(userName string) (int, error) {
	result, err := up.DB.Exec("delete User where uesr_name = ?", userName)
	if err != nil {
		return consts.OPERATION_FAILED, err
	}
	res, res_err := result.RowsAffected()
	if res_err != nil {
		return consts.OPERATION_FAILED, res_err
	}
	if res > 0 {
		return consts.OPERATION_SUCCSSED, nil
	}
	return consts.OPERATION_NO_EFFECT, nil
}
