package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/tpisntgod/Service_Calculation/cloudgo-io/view"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	/*
		formatter := render.New(render.Options{
			Directory:  "templates",
			Extensions: []string{".html"},
			IndentJSON: true,
		})*/

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	fmt.Println("webroot: " + webRoot)
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	fmt.Println("webroot22: " + webRoot)

	//mainpage中js请求支持
	mx.HandleFunc("/api/mainpage", mainPageHandler(formatter)).Methods("GET")

	mx.HandleFunc("/mainpage", view.MainPage)
	//静态文件服务
	mx.HandleFunc("/register", view.Register)
	mx.HandleFunc("/login", view.Login)

	mx.HandleFunc("/logout", view.Logout)
	mx.HandleFunc("/todolistInformation", view.GetTodoListInformation)
	mx.HandleFunc("/todoitemAddition", view.AddTodoItem)
	mx.HandleFunc("/todoitemQuery", view.QueryTodoItem)
	mx.HandleFunc("/todoitemDeletion", view.DeleteTodoItem)

	//模板输出
	mx.HandleFunc("/result", view.Result)

	//对 /api/unknown 给出开发中的提示
	mx.PathPrefix("/api").Handler(view.NotImplementedHandler())

	//静态文件服务
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}
