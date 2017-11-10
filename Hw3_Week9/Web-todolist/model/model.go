package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//UserAccounts 用户账户类型
type UserAccounts struct {
	uid      int64
	username string
	password string
}

type todolist struct {
	tid        int64
	username   string
	thingstodo string
}

var accounts []UserAccounts
var todolists []todolist
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

//AddTodoItem 给用户添加todoitem
func AddTodoItem(username string, todoitem string) error {
	stmt, err := db.Prepare("INSERT todolist SET username=?,thingstodo=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(username, todoitem)
	if err != nil {
		return err
	}
	tid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	var todo todolist
	todo.tid = tid
	todo.username = username
	todo.thingstodo = todoitem
	todolists = append(todolists, todo)
	return nil
}

//QueryTodoItem 查询用户的所有todoitem
func QueryTodoItem(username string) string {
	var todoItemInfo = "todoitem id(根据tid删除todoitem): todoitem内容:\n"
	var hasTodoItem = false
	for i := 0; i < len(todolists); i++ {
		if todolists[i].username == username {
			hasTodoItem = true
			todoItemInfo += "         " + strconv.FormatInt(todolists[i].tid, 10)+ "    " + todolists[i].thingstodo + "\n"
		}
	}
	if hasTodoItem == false {
		todoItemInfo = ""
	}
	return todoItemInfo
}

//DeleteTodoItem 给用户添加todoitem
func DeleteTodoItem(username string, tid int64) error {
	var isTidExists = false
	for i := 0; i < len(todolists); i++ {
		if todolists[i].tid == tid {
			isTidExists = true
			if todolists[i].username != username {
				return errors.New("添加该todoitem的用户不是当前用户")
			}
			stmt, err := db.Prepare("delete from todolist where tid=?")
			if err != nil {
				return err
			}
			_, err = stmt.Exec(tid)
			if err != nil {
				return err
			}
			todolists = append(todolists[:i], todolists[i+1:]...)
			break
		}
	}
	if isTidExists == false {
		var errInfo = fmt.Sprintf("没有todoitem的tid是:%d", tid)
		return errors.New(errInfo)
	}
	return nil
}

func init() {
	database, err := sql.Open("mysql", "root:houxi5201314@/todolist_accounts?charset=utf8")
	checkErr(err)
	db = database

	//把mysql数据库的user账户信息导入
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var user UserAccounts
		err := rows.Scan(&user.uid, &user.username, &user.password)
		checkErr(err)
		fmt.Println(user)
		accounts = append(accounts, user)
	}

	//把mysql数据库的todolist信息导入
	todorows, err := db.Query("SELECT * FROM todolist")
	checkErr(err)
	fmt.Println("todorows")
	for todorows.Next() {
		var todo todolist
		err := todorows.Scan(&todo.tid, &todo.username, &todo.thingstodo)
		checkErr(err)
		fmt.Println(todo)
		todolists = append(todolists, todo)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
