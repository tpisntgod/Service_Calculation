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

	"github.com/tpisntgod/Service_Calculation/Hw4_Week12/cloudgo-io-update/controller"
)

var htmlpath = "src/github.com/tpisntgod/Service_Calculation/Hw4_Week12/cloudgo-io-update/templates/"
var pathtoAdd = "/src/github.com/tpisntgod/Service_Calculation/Hw4_Week12/cloudgo-io-update/templates/mainpage.html"

//ResultInfo 所有操作redirect到result 用于handler修改模板
var ResultInfo string

//Islogout js中ajax访问时判断是否登出
var Islogout = false

var todoListInfoTemplate *template.Template
var resultTemplate *template.Template

type todoListInfo struct {
	FirstTerm  string
	SecondTerm string
}

type resultInfomation struct {
	Information string
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
		if _, err := os.Stat(filepath.Join(v, pathtoAdd)); !os.IsNotExist(err) {
			return &v
		}
	}
	return nil
}

var i int

//MainPage 主页面
func MainPage(w http.ResponseWriter, r *http.Request) {
	mainpage, err := ioutil.ReadFile(htmlpath + "/mainpage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(mainpage))
}

//Register 注册用户
func Register(w http.ResponseWriter, r *http.Request) {
	_, cookieError := r.Cookie("username")
	if cookieError == nil {
		fmt.Fprintf(w, "please log out first")
		return
	}
	if r.Method == "GET" {
		http.Redirect(w, r, "/html/register.html", http.StatusFound)
	}
	if r.Method == "POST" {
		if err := controller.RegisterUser(r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		r.ParseForm()
		cookie := http.Cookie{Name: "username", Value: r.Form["username"][0], Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/mainpage", http.StatusFound)
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
		http.Redirect(w, r, "/html/login.html", http.StatusFound)
	}
	if r.Method == "POST" {
		if err := controller.LoginUser(r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		r.ParseForm()
		cookie := http.Cookie{Name: "username", Value: r.Form["username"][0], Path: "/", MaxAge: 600}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/mainpage", http.StatusFound)
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
	Islogout = true
	http.Redirect(w, r, "/mainpage", http.StatusFound)
}

//GetTodoListInformation 查看TodoList系统信息
func GetTodoListInformation(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/html/todolistInfoEnquiry.html", http.StatusFound)
	}

	if r.Method == "POST" {
		var todoListInformation todoListInfo
		r.ParseForm()
		if r.Form["queryType"][0] == "2" {
			todoListInformation = todoListInfo{FirstTerm: "开发者信息", SecondTerm: "toupiOfRivia"}
		} else {
			todoListInformation = todoListInfo{FirstTerm: "系统使用方法", SecondTerm: "github的README有写哦_(:з」∠)_"}
		}
		err := todoListInfoTemplate.Execute(w, todoListInformation)
		if err != nil {
			fmt.Fprintf(w, "what222")
			fmt.Fprintf(w, err.Error())
			return
		}
	}
}

//AddTodoItem 用户增加todo条目
func AddTodoItem(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		ResultInfo = "you have not logged in"
		http.Redirect(w, r, "/result", http.StatusFound)
		return
	}
	if r.Method == "GET" {
		http.Redirect(w, r, "/html/todoitemAddition.html", http.StatusFound)
	}
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("add todoitem username:" + cookie.Value)
		if err := controller.AddTodoItem(cookie.Value, r); err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		ResultInfo = "add todoitem successfully!"
		http.Redirect(w, r, "/result", http.StatusFound)
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

	ResultInfo = todoItemInfo
	http.Redirect(w, r, "/result", http.StatusFound)
}

//DeleteTodoItem 用户删除todo条目
func DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("username")
	if cookieError != nil {
		ResultInfo = "you have not logged in"
		http.Redirect(w, r, "/result", http.StatusFound)
		return
	}
	if r.Method == "GET" {
		http.Redirect(w, r, "/html/todoitemDeletion.html", http.StatusFound)
	}
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("delete todoitem username:" + cookie.Value)
		if err := controller.DeleteTodoItem(cookie.Value, r); err != nil {
			ResultInfo = err.Error()
			http.Redirect(w, r, "/result", http.StatusFound)
			return
		}
		ResultInfo = "delete todoitem successfully!"
		http.Redirect(w, r, "/result", http.StatusFound)
		//fmt.Fprintf(w, "delete todoitem successfully!")
	}
}

//Result 显示操作结果
func Result(w http.ResponseWriter, r *http.Request) {
	//t, err := template.ParseFiles(htmlpath + "/result.html")

	var resultInfo resultInfomation

	resultInfo = resultInfomation{Information: ResultInfo}

	err := resultTemplate.Execute(w, resultInfo)
	if err != nil {
		fmt.Fprintf(w, "what222")
		fmt.Fprintf(w, err.Error())
		return
	}
}

//NotImplemented 客户端访问路由是 目前没有开发的功能 的处理函数
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 function Not Implemented", http.StatusNotFound)
}

//NotImplementedHandler 处理函数类型转换成http.Handler
func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(NotImplemented)
}

func init() {
	var err error
	htmlpath = filepath.Join(*GetGOPATH(), htmlpath)
	resultTemplate, err = template.ParseFiles(htmlpath + "/result.html")
	if err != nil {
		fmt.Println("result template parse failed" + err.Error())
		return
	}
	todoListInfoTemplate, err = template.ParseFiles(htmlpath + "/todolistInfo.html")
	if err != nil {
		fmt.Println("result template parse failed" + err.Error())
		return
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
