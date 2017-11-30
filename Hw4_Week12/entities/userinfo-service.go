package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

//RegistUser .
func (*UserInfoAtomicService) RegistUser(u *UserAccounts) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{tx}
	err = dao.RegistUser(u)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

//FindAll .
func (*UserInfoAtomicService) FindAll() []UserAccounts {
	dao := userInfoDao{mydb}
	return dao.FindAll()
}

//FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserAccounts {
	dao := userInfoDao{mydb}
	return dao.FindByID(id)
}

//FindByUserName .
func (*UserInfoAtomicService) FindByUserName(username string) *UserAccounts {
	dao := userInfoDao{mydb}
	return dao.FindByUserName(username)
}
