package main

import (
	"log"
	"net/http"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-todolist/view"
)

func main() {
	http.HandleFunc("/", view.MainPage)
	http.HandleFunc("/register", view.Register)
	http.HandleFunc("/login", view.Login)
	http.HandleFunc("/logout", view.Logout)
	http.HandleFunc("/todoitemAddition", view.AddTodoItem)
	http.HandleFunc("/todoitemQuery", view.QueryTodoItem)
	http.HandleFunc("/todoitemDeletion", view.DeleteTodoItem)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
