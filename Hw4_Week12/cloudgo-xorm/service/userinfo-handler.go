package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tpisntgod/Service_Calculation/Hw4_Week12/cloudgo-xorm/entities"
	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"No username input,Bad Input!"})
			return
		}
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Username is blank,Bad Input!"})
			return
		}
		u := entities.UserInfo{UserName: req.Form["username"][0]}
		if len(req.Form["department"]) > 0 {
			fmt.Println(req.Form["department"][0])
			u.DepartName = req.Form["department"][0]
		}
		entities.UserInfoService.Save(&u)
		fmt.Println("handler:", u)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println("get method para uid:", req.Form["userid"])

		if len(req.Form["userid"]) == 0 {
			userlist := entities.UserInfoService.FindAll()
			formatter.JSON(w, http.StatusOK, userlist)
			return
		}
		if len(req.Form["userid"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Userid is blank,Bad Input!"})
			return
		}
		fmt.Println("get method userid:", req.Form["userid"][0])
		i, err := strconv.Atoi(req.Form["userid"][0])
		checkErr(err)
		user := entities.UserInfoService.FindByID(i)
		if user.UID == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"the user doesn't exist"})
			return
		}
		formatter.JSON(w, http.StatusOK, user)
		return
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
