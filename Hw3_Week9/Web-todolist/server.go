package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/controller"
)

func main() {
	http.HandleFunc("/", controller.MainPage)
	http.HandleFunc("/register", controller.Register)
	/*
		http.HandleFunc("/login", controller.Login)
		http.HandleFunc("/logout", controller.Logout)
		http.HandleFunc("/add", controller.CreateMeeting)
		http.HandleFunc("/query", controller.QueryMeeting)
		http.HandleFunc("/delete", controller.CancelMeeting)
	*/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
