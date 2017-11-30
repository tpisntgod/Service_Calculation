package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	dao := userInfoDao{}
	err := dao.Save(u)
	return err
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{}
	return dao.FindAll()
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{}
	return dao.FindByID(id)
}
