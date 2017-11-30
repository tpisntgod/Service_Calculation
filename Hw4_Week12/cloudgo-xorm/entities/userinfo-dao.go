package entities

import (
	"fmt"
)

type userInfoDao DaoSource

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	_, err := myengine.Insert(u)
	checkErr(err)
	fmt.Println("dao:", *u)
	return err
}

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
