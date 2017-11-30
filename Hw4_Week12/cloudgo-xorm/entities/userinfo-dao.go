package entities

import (
	"fmt"
)

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	_, err := myengine.Insert(u)
	checkErr(err)
	fmt.Println("dao:", *u)
	//u.UID = int(id)
	return err
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserInfo {
	userlist := make([]UserInfo, 0, 0)
	err := myengine.Find(&userlist)
	checkErr(err)
	return userlist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserInfo {
	user := &UserInfo{UID: id}
	has, err := myengine.Get(user)
	checkErr(err)
	if has == false {
		return new(UserInfo)
	}
	return user
}
