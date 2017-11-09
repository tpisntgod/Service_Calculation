package model

import "database/sql"
import "fmt"

//UserAccounts 用户账户类型
type UserAccounts struct {
	uid      int64
	username string
	password string
}

var accounts []UserAccounts
var db *sql.DB

//QueryUser 查询一个用户是否存在
func QueryUser(username string) bool {
	fmt.Println(len(accounts))
	for i := 0; i < len(accounts); i++ {
		if accounts[i].username == username {
			return true
		}
	}
	return false
}

//RegisterUser 注册一个用户
func RegisterUser(username string, password string) error {
	//insertSql := fmt.Sprintf("INSERT userinfo SET username=%s,password=%s", username, password)
	stmt, err := db.Prepare("INSERT userinfo SET username=?,password=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(username, password)
	if err != nil {
		return err
	}
	uid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	var user UserAccounts
	user.uid = uid
	user.username = username
	user.password = password
	accounts = append(accounts, user)
	return nil
}

func init() {
	database, err := sql.Open("mysql", "root:houxi5201314@/todolist_accounts?charset=utf8")
	checkErr(err)
	db = database
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		fmt.Println("rowsnext")
		var user UserAccounts
		err := rows.Scan(&user.uid, &user.username, &user.password)
		checkErr(err)
		fmt.Println(user)
		accounts = append(accounts, user)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
