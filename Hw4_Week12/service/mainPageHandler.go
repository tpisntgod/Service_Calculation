package service

import (
	"fmt"
	"net/http"

	"github.com/tpisntgod/Service_Calculation/Hw4_Week12/view"
	"github.com/unrolled/render"
)

func mainPageHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("path", r.URL.Path)
		fmt.Println("mainPageHandler")

		var userMainPage string
		cookie, cookieError := r.Cookie("username")
		if cookieError != nil {
			fmt.Println("cookieError:" + cookieError.Error())
		} else {
			fmt.Println("cookie doesn't have error")
		}
		if cookieError != nil {
			if view.Islogout {
				userMainPage = "you have logged out successfully!"
				view.Islogout = false
			} else {
				userMainPage = "please sign in to use todolist"
			}
		} else {
			userMainPage = "welcome " + cookie.Value + "!"
		}

		fmt.Println("mainPageHandler")
		formatter.JSON(w, http.StatusOK, struct {
			Username string `json:"username"`
		}{Username: userMainPage})
	}
}
