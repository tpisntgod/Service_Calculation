package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/tpisntgod/Service_Calculation/Hw4_Week12/entities"
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

	userFind := entities.UserInfoService.FindByUserName(r.Form["username"][0])
	fmt.Println("regis userfind:", userFind)
	if userFind.UID != 0 {
		return errors.New("该用户已经被注册")
	}

	u := entities.UserAccounts{UserName: r.Form["username"][0], Password: r.Form["password"][0]}
	err := entities.UserInfoService.RegistUser(&u)
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
	userFind := entities.UserInfoService.FindByUserName(r.Form["username"][0])
	fmt.Println("login userfind:", *userFind)
	if userFind.UID == 0 {
		return errors.New("该用户还没有被注册")
	}
	return nil
}

//AddTodoItem 增加用户todoitem
func AddTodoItem(username string, r *http.Request) error {
	if r.Form["todoitem"][0] == "" {
		return errors.New("todoitem不能为空")
	}
	fmt.Println("add todoitem "+username, r.Form["todoitem"][0])

	t := entities.TodoItem{UserName: username, Todo: r.Form["todoitem"][0]}
	err := entities.TodoItemService.RegistTodoItem(&t)
	return err
}

//QueryTodoItem 查找用户todoitem
func QueryTodoItem(username string) (string, error) {
	todoItemInfo := "todoitem id(根据tid删除todoitem): todoitem内容:\n"
	todolists := entities.TodoItemService.FindAllByUserName(username)
	if len(todolists) == 0 {
		return "", errors.New("该用户没有todoitem")
	}
	for i := 0; i < len(todolists); i++ {
		if todolists[i].UserName == username {
			todoItemInfo += "         " + strconv.Itoa(todolists[i].TID) + "    " + todolists[i].Todo + "\n"
		}
	}
	return todoItemInfo, nil
}

//DeleteTodoItem 删除用户todoitem
func DeleteTodoItem(username string, r *http.Request) error {
	if r.Form["tid"][0] == "" {
		return errors.New("tid不能为空")
	}
	tid, err := strconv.Atoi(r.Form["tid"][0])
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	if err != nil {
		return err
	}
	fmt.Println(username, tid)

	t := entities.TodoItem{TID: tid, UserName: username}
	err = entities.TodoItemService.DeleteTodoItem(&t)
	return err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
