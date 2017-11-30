package entities

import (
	"database/sql"
	"errors"
	"fmt"
)

type todoItemDao DaoSource

var todoItemInsertStmt = "INSERT todolist SET username=?,thingstodo=?"
var todoItemDeleteStmt = "DELETE from todolist where tid=?"

// AddtodoItem .
func (dao *todoItemDao) AddtodoItem(t *TodoItem) error {
	stmt, err := dao.Prepare(todoItemInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(t.UserName, t.Todo)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	t.TID = int(id)
	return nil
}

// DeletetodoItem .
func (dao *todoItemDao) DeletetodoItem(t *TodoItem) error {
	findItem := dao.FindtodoItemByID(t.TID)
	fmt.Println("FindtodoItemByID123: ", findItem)
	if findItem.TID == 0 {
		errInfo := fmt.Sprintf("没有todoitem的tid是:%d", t.TID)
		return errors.New(errInfo)
	}
	if findItem.UserName != t.UserName {
		return errors.New("添加该todoitem的用户不是当前用户")
	}

	stmt, err := dao.Prepare(todoItemDeleteStmt)
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(t.TID)
	checkErr(err)
	return err
}

var todoItemQueryAll = "SELECT * FROM todolist"
var todoItemQueryAllByUserName = "SELECT * FROM todolist where username = ?"
var todoItemQueryByID = "SELECT * FROM todolist where tid = ?"
var todoItemQueryByUserName = "SELECT * FROM todolist where username = ?"

// FindAll .
func (dao *todoItemDao) FindAlltodoItems() []TodoItem {
	rows, err := dao.Query(todoItemQueryAll)
	checkErr(err)
	defer rows.Close()

	todoItemlist := make([]TodoItem, 0, 0)
	for rows.Next() {
		t := TodoItem{}
		err := rows.Scan(&t.TID, &t.UserName, &t.Todo)
		checkErr(err)
		todoItemlist = append(todoItemlist, t)
	}
	return todoItemlist
}

// FindAlltodoItemsByUserName .
func (dao *todoItemDao) FindAlltodoItemsByUserName(username string) []TodoItem {
	rows, err := dao.Query(todoItemQueryAllByUserName, username)
	checkErr(err)
	defer rows.Close()

	todoItemlist := make([]TodoItem, 0, 0)
	for rows.Next() {
		t := TodoItem{}
		err := rows.Scan(&t.TID, &t.UserName, &t.Todo)
		checkErr(err)
		todoItemlist = append(todoItemlist, t)
	}
	return todoItemlist
}

// FindByID .
func (dao *todoItemDao) FindtodoItemByID(id int) *TodoItem {
	stmt, err := dao.Prepare(todoItemQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	t := TodoItem{}
	err = row.Scan(&t.TID, &t.UserName, &t.Todo)
	if err == sql.ErrNoRows {
		return &t
	}
	checkErr(err)

	return &t
}

// FindByUserName .
func (dao *todoItemDao) FindtodoItemByUserName(username string) *TodoItem {
	stmt, err := dao.Prepare(todoItemQueryByUserName)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(username)
	t := TodoItem{}
	err = row.Scan(&t.TID, &t.UserName, &t.Todo)
	if err == sql.ErrNoRows {
		return &t
	}
	checkErr(err)
	return &t
}
