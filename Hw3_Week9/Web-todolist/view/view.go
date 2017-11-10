package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/controller"
)

var htmlpath = "src/github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/html_template"
var islogout = false

type userTemplate struct {
	Username string
}

//GetGOPATH 获得用户环境的gopath
func GetGOPATH() *string {
	var sp string
	if runtime.GOOS == "windows" {
		sp = ";"
	} else {
		sp = ":"
	}
	goPath := strings.Split(os.Getenv("GOPATH"), sp)
	for _, v := range goPath {
		if _, err := os.Stat(filepath.Join(v, "/src/github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/html_template/mainpage.html")); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}

var i int

//MainPage 主页面
func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlpath + "/mainpage.html")
	if err != nil {
		fmt.Fprintf(w, "what")
		fmt.Fprintf(w, err.Error())
		return
	}
	var userMainpage userTemplate

	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		fmt.Println("cookieError:" + cookieError.Error())
	} else {
		fmt.Println("cookie doesn't have error")
	}
	if cookieError != nil {
		if islogout {
			userMainpage = userTemplate{Username: "you have logged out successfully!"}
			islogout = false
		} else {
			userMainpage = userTemplate{Username: "please sign in to use todolist"}
		}
	} else {
		userMainpage = userTemplate{Username: "welcome " + cookie.Value + "!"}
	}

	err = t.Execute(w, userMainpage)
	if err != nil {
		fmt.Fprintf(w, "what222")
		fmt.Fprintf(w, err.Error())
		return
	}
}

//Register 注册用户
func Register(w http.ResponseWriter, r *http.Request) {
	_, cookieError := r.Cookie("username")
	if cookieError == nil {
		fmt.Fprintf(w, "please log out first")
		return
	}
	if r.Method == "GET" {
		registerpage, err := ioutil.ReadFile(htmlpath + "/register.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, string(registerpage))
	}
	if r.Method == "POST" {
		if err := controller.RegisterUser(r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		r.ParseForm()
		expiration := time.Now()
		expiration = expiration.Add(time.Minute * 10)
		cookie := http.Cookie{Name: "username", Value: r.Form["username"][0], Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//Login 用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	_, cookieError := r.Cookie("username")
	if cookieError == nil {
		fmt.Fprintf(w, "please log out first")
		return
	}
	if r.Method == "GET" {
		registerpage, err := ioutil.ReadFile(htmlpath + "/login.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, string(registerpage))
	}
	if r.Method == "POST" {
		if err := controller.LoginUser(r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		r.ParseForm()
		cookie := http.Cookie{Name: "username", Value: r.Form["username"][0], Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//Logout 用户登录
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		fmt.Fprintf(w, "you have not logged in")
		return
	}
	//fmt.Fprintf(w, "user:"+cookie.Value+"logged out successfully")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	islogout = true
	http.Redirect(w, r, "/", http.StatusFound)
}

//AddTodoItem 用户增加todo条目
func AddTodoItem(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		fmt.Fprintf(w, "you have not logged in")
		return
	}
	if r.Method == "GET" {
		addTodoItemPage, err := ioutil.ReadFile(htmlpath + "/todoitemAddition.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, string(addTodoItemPage))
	}
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("add todoitem username:" + cookie.Value)
		if err := controller.AddTodoItem(cookie.Value, r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, "add todoitem successfully!")
	}
}

//QueryTodoItem 用户查询自己的所有todo条目
func QueryTodoItem(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		fmt.Fprintf(w, "you have not logged in")
		return
	}
	fmt.Println("query todoitem username:" + cookie.Value)
	todoItemInfo, err := controller.QueryTodoItem(cookie.Value)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, todoItemInfo)
}

//DeleteTodoItem 用户删除todo条目
func DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		fmt.Fprintf(w, "you have not logged in")
		return
	}
	if r.Method == "GET" {
		deleteTodoItemPage, err := ioutil.ReadFile(htmlpath + "/todoitemDeletion.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, string(deleteTodoItemPage))
	}
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("delete todoitem username:" + cookie.Value)
		if err := controller.DeleteTodoItem(cookie.Value, r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		fmt.Fprintf(w, "delete todoitem successfully!")
	}
}

func init() {
	htmlpath = filepath.Join(*GetGOPATH(), htmlpath)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
