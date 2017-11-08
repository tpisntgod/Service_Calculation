package main

import (
	"log"
	"net/http"

	"github.com/tpisntgod/Service_Calculation/Hw3_Week9/Web-Agenda/controller"
)

func main() {
	http.HandleFunc("/", controller.MainPage)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/cm", controller.CreateMeeting)
	http.HandleFunc("/qm", controller.QueryMeeting)
	http.HandleFunc("/cancelmeeting", controller.CancelMeeting)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
