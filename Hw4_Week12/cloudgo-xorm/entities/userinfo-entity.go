package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `xorm:"pk autoincr 'uid'"` //语义标签
	UserName   string     `xorm:"varchar(64) notnull 'username'"`
	DepartName string     `xorm:"varchar(64) 'departmentName'"`
	CreatedAt  *time.Time `xorm:"created 'createTime'"`
}

//TableName .
func (UserInfo) TableName() string {
	return "userinformation"
}
