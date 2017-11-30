package entities

import (
	"database/sql"
)

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT userinfo SET username=?,password=?"

// RegistUser .
func (dao *userInfoDao) RegistUser(u *UserAccounts) error {
	stmt, err := dao.Prepare(userInfoInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(u.UserName, u.Password)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.UID = int(id)
	return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"
var userInfoQueryByUserName = "SELECT * FROM userinfo where username = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserAccounts {
	rows, err := dao.Query(userInfoQueryAll)
	checkErr(err)
	defer rows.Close()

	userlist := make([]UserAccounts, 0, 0)
	for rows.Next() {
		u := UserAccounts{}
		err := rows.Scan(&u.UID, &u.UserName, &u.Password)
		checkErr(err)
		userlist = append(userlist, u)
	}
	return userlist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserAccounts {
	stmt, err := dao.Prepare(userInfoQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	u := UserAccounts{}
	err = row.Scan(&u.UID, &u.UserName, &u.Password)
	if err == sql.ErrNoRows {
		return &u
	}
	checkErr(err)

	return &u
}

// FindByUserName .
func (dao *userInfoDao) FindByUserName(username string) *UserAccounts {
	stmt, err := dao.Prepare(userInfoQueryByUserName)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := UserAccounts{}
	err = row.Scan(&u.UID, &u.UserName, &u.Password)
	if err == sql.ErrNoRows {
		return &u
	}
	checkErr(err)
	return &u
}
