package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/model"
)

var htmlpath = "src/github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/html_template"

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
	/*
		t := template.Must(template.ParseFiles("mainpage.html"))
		t.Execute(w,{})
	*/

	mainpage, err := ioutil.ReadFile(htmlpath + "/mainpage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(mainpage))
}

//Register 注册用户
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		registerpage, err := ioutil.ReadFile(htmlpath + "/register.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		fmt.Fprintf(w, string(registerpage))
	}
	if r.Method == "POST" {
		r.ParseForm()
		if r.Form["username"][0] == "" {
			fmt.Fprintf(w, "用户名不能为空")
			return
		}
		if r.Form["password"][0] == "" {
			fmt.Fprintf(w, "密码不能为空")
			return
		}
		fmt.Println(r.Form["username"][0])
		fmt.Println(r.Form["password"][0])
		if model.QueryUser(r.Form["username"][0]) {
			fmt.Fprintf(w, "此用户名已经被注册")
			return
		}
		model.RegisterUser(r.Form["username"][0], r.Form["password"][0])
		http.Redirect(w, r, "/", http.StatusFound)
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
