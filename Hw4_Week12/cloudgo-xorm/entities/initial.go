package entities

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var myengine *xorm.Engine

func init() {
	engine, err := xorm.NewEngine("mysql", "root:houxi5201314@/xorm_userinfo?charset=utf8")
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(UserInfo))
	if err != nil {
		panic(err)
	}
	myengine = engine
}

// DaoSource Data Access Object Source
type DaoSource struct{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
