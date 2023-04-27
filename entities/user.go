package entities

import (
	"container/list"
	"errors"

	"github.com/mohsenGoodarzi/webframework/utils"
)

type User struct {
	Uid      string
	Password string
	Acl      *utils.ACL
}

type Users struct {
	AllUsers list.List
}

func InitUsers() (users *Users) {

	users = &Users{AllUsers: *list.New()}
	// users will be fetched when persistent layer is ready
	userOne := &User{
		"a.outlook.com",
		"hashed_Password",
		utils.CreateACL(),
	}
	userOne.Acl.EndPoints["Recordings"] = userOne.Acl.PermissionTable.GetCodeNumber([]string{"GET", "PUT", "POST", "DELETE"})
	userOne.Acl.EndPoints["About"] = userOne.Acl.PermissionTable.GetCodeNumber([]string{"GET"})

	userTwo := &User{
		"b.outlook.com",
		"hashed_Password",
		utils.CreateACL(),
	}
	userOne.Acl.EndPoints["Recordings"] = userOne.Acl.PermissionTable.GetCodeNumber([]string{"GET", "PUT", "DELETE"})
	userOne.Acl.EndPoints["About"] = userOne.Acl.PermissionTable.GetCodeNumber([]string{"POST"})

	users.AllUsers.PushBack(userOne)
	users.AllUsers.PushBack(userTwo)

	return users
}

func (users Users) AddUser(user User) int {
	current := users.AllUsers.Len()
	users.AllUsers.PushBack(user)
	new := users.AllUsers.Len()
	if (new + 1) == current {
		return utils.OPERATION_SUCCSSED
	}
	return utils.OPERATION_FAILED
}
func (users Users) RemoveUser(user User) int {

	users.AllUsers.Remove(
		&list.Element{
			Value: user},
	)
	return utils.OPERATION_SUCCSSED
}
func (users Users) FindUser(user User) (User, error) {
	for e := users.AllUsers.Front(); e != nil; e = e.Next() {
		if e.Value == user {
			return e.Value.(User), nil
		}
	}
	return User{}, errors.New("User not found")
}

func (users Users) Update(oldUser *User, newUser *User) (int, error) {
	user, err := users.FindUser(*oldUser)
	if err != nil {
		return utils.OPERATION_FAILED, err
	}
	user.Uid = newUser.Uid
	user.Password = newUser.Password
	user.Acl = newUser.Acl
	return utils.OPERATION_SUCCSSED, nil
}
