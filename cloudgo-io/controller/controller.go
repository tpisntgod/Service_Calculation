package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/model"
)

//RegisterUser 注册用户
func RegisterUser(r *http.Request) error {
	r.ParseForm()
	//是否需要判断len大于1，url和表单同时提交了对应item
	if r.Form["username"][0] == "" {
		return errors.New("用户名不能为空")
	}
	if r.Form["password"][0] == "" {
		return errors.New("密码不能为空")
	}
	fmt.Println(r.Form["username"][0])
	fmt.Println(r.Form["password"][0])
	if model.QueryUser(r.Form["username"][0]) {
		return errors.New("此用户名已经被注册")
	}
	err := model.RegisterUser(r.Form["username"][0], r.Form["password"][0])
	return err
}

//LoginUser 登录用户
func LoginUser(r *http.Request) error {
	r.ParseForm()
	if r.Form["username"][0] == "" {
		return errors.New("用户名不能为空")
	}
	if r.Form["password"][0] == "" {
		return errors.New("密码不能为空")
	}
	fmt.Println(r.Form["username"][0])
	fmt.Println(r.Form["password"][0])
	if model.QueryUser(r.Form["username"][0]) {
		return nil
	}
	return errors.New("此用户还没有注册")
}

//QueryUser 查询用户是否存在
func QueryUser(r *http.Request) error {
	r.ParseForm()
	if r.Form["username"][0] == "" {
		return errors.New("用户名不能为空")
	}
	fmt.Println(r.Form["username"][0])
	if model.QueryUser(r.Form["username"][0]) {
		return nil
	}
	return errors.New("此用户还没有注册")
}

//AddTodoItem 增加用户todoitem
func AddTodoItem(username string, r *http.Request) error {
	if r.Form["todoitem"][0] == "" {
		return errors.New("todoitem不能为空")
	}
	fmt.Println("add todoitem "+username, r.Form["todoitem"][0])
	err := model.AddTodoItem(username, r.Form["todoitem"][0])
	return err
}

//QueryTodoItem 查找用户todoitem
func QueryTodoItem(username string) (string, error) {
	todoItemInfo := model.QueryTodoItem(username)
	if todoItemInfo == "" {
		return todoItemInfo, errors.New("该用户没有todoitem")
	}
	return todoItemInfo, nil
}

//DeleteTodoItem 删除用户todoitem
func DeleteTodoItem(username string, r *http.Request) error {
	if r.Form["tid"][0] == "" {
		return errors.New("tid不能为空")
	}
	tid, err := strconv.ParseInt(r.Form["tid"][0], 10, 64)
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	if err != nil {
		return err
	}
	fmt.Println(username, tid)
	err = model.DeleteTodoItem(username, tid)
	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
