package entities

//UserAccounts 用户账户类型
type UserAccounts struct {
	UID      int //之前用的是int64，为什么？ 貌似是LastInsertId()的返回值是int64
	UserName string
	Password string
}

/*
// NewUserInfo .
func NewUserInfo(u UserAccounts) *UserAccounts {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
*/
