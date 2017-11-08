package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tpisntgod/Agenda/entity/meeting"
	"github.com/tpisntgod/Agenda/entity/user"
)

//MainPage 主页面
func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!This is a Agenda Web App\nHere is usage:")
	fmt.Fprintf(w, `
	to register an account: http://localhost:8080/register?name=qwe1&password=qwe
	to login an account: http://localhost:8080/login?name=qwe1&password=qwe
	to logout an account: http://localhost:8080/logout
	to create a meeting: http://localhost:8080/cm?title=webmeet1&p=aa&p=cc&stime=2017-10-28 09:30:00&etime=2017-10-28 10:30:00
	to query meetings in an interval: http://localhost:8080/qm?stime=2017-10-01 09:30:00&etime=2017-10-31 10:30:00
	to cancel a meeting that you host: http://localhost:8080/cancelmeeting?title=meet`)
}

func checkUsernameandPassword(r *http.Request) string {
	var s string
	if len(r.Form["name"]) > 1 {
		s = s + "only one username should be input\n"
	}
	if len(r.Form["name"]) == 0 {
		s = s + "Username of your account can't be blank. Please input the username of your account in url as query\n"
		s = s + "Example:http://localhost:8080/register?name=bbb\n"
	}
	if len(r.Form["password"]) > 1 {
		s = s + "only one username should be input"
	}
	if len(r.Form["password"]) == 0 {
		s = s + "Password of your account can't be blank. Please input the password of your account in url as query"
		s = s + "Example:http://localhost:8080/register?password=psw\n"
	}
	return s
}

//Register 注册用户
func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := checkUsernameandPassword(r)
	if err != "" {
		fmt.Fprintf(w, err)
		return
	}
	err2 := user.RegisterUser(r.Form["name"][0], r.Form["password"][0], r.Form["name"][0]+"@qq.com", "123")
	if err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}
	fmt.Fprintf(w, "a new account is registered\nname: %s password: %s", r.Form["name"][0], r.Form["password"][0])
}

//Login 登录用户
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	err := checkUsernameandPassword(r)
	if err != "" {
		fmt.Fprintf(w, err)
		return
	}
	err2 := user.LoginUser(r.Form["name"][0], r.Form["password"][0])
	if err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}
	fmt.Fprintf(w, "%s has logined successfully", r.Form["name"][0])
}

//Logout 登出用户
func Logout(w http.ResponseWriter, r *http.Request) {
	err := user.LogoutUser()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "logout successfully")
}

func checkTimeIntervalValid(r *http.Request) string {
	var err string
	if len(r.Form["stime"]) == 0 {
		err = err + "you should input meeting start time\n"
		err = err + "Example: stime=2017-10-28 09:30:00\n"
	}
	if len(r.Form["etime"]) == 0 {
		err = err + "you should input meeting end time\n"
		err = err + "Example: etime=2017-10-28 10:30:00\n"
	}
	if len(r.Form["stime"]) > 1 {
		err = err + "you should input exactly one meeting start time\n"
	}
	if len(r.Form["etime"]) > 1 {
		err = err + "you should input exactly one meeting end time\n"
	}
	return err
}

//CreateMeeting 创建会议
func CreateMeeting(w http.ResponseWriter, r *http.Request) {
	//用户需要先登录
	if !user.IsLogin() {
		fmt.Fprintf(w, "Please login first!")
		return
	}
	r.ParseForm()
	//需要输入会议名称
	if len(r.Form["title"]) == 0 {
		fmt.Fprintf(w, "you should input meeting title")
		return
	}
	//会议名称唯一
	if len(r.Form["title"]) > 1 {
		fmt.Fprintf(w, "only one title should be input!")
		return
	}
	//需要输入会议开始和结束时间
	err := checkTimeIntervalValid(r)
	if err != "" {
		fmt.Fprintf(w, err)
		return
	}
	//会议开始时间需要小于结束时间
	t1, _ := time.Parse("2006-01-02 15:04:05", r.Form["stime"][0])
	t2, _ := time.Parse("2006-01-02 15:04:05", r.Form["etime"][0])
	if !meeting.CheckStarttimelessthanEndtime(t1, t2) {
		fmt.Fprintf(w, "meeting start time should be less than end time")
		return
	}
	//会议参加成员必须已经注册
	isUsersUnregistered := 0
	var unregistered string
	for i := 0; i < len(r.Form["p"]); i++ {
		if !user.IsRegisteredUser(r.Form["p"][i]) {
			isUsersUnregistered = 1
			unregistered += r.Form["p"][i] + " is not registered\n"
		}
	}
	if isUsersUnregistered == 1 {
		fmt.Fprintf(w, unregistered)
		return
	}
	if err2 := meeting.CreateMeeting(r.Form["title"][0], r.Form["p"], t1, t2); err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}
	fmt.Fprintf(w, "meeting: "+r.Form["title"][0]+" is created")
}

//QueryMeeting 查询一个时间段的会议情况
func QueryMeeting(w http.ResponseWriter, r *http.Request) {
	//用户需要先登录
	if !user.IsLogin() {
		fmt.Fprintf(w, "Please login first!")
		return
	}
	r.ParseForm()
	//需要输入会议开始和结束时间
	err := checkTimeIntervalValid(r)
	if err != "" {
		fmt.Fprintf(w, err)
		return
	}
	//会议开始时间需要小于结束时间
	t1, _ := time.Parse("2006-01-02 15:04:05", r.Form["stime"][0])
	t2, _ := time.Parse("2006-01-02 15:04:05", r.Form["etime"][0])
	if !meeting.CheckStarttimelessthanEndtime(t1, t2) {
		fmt.Fprintf(w, "meeting start time should be less than end time")
		return
	}
	queryResult, err2 := meeting.QueryMeetingWebVersion(t1, t2)
	if err2 != nil {
		fmt.Fprintf(w, err2.Error())
		return
	}
	fmt.Fprintf(w, queryResult)
}

//CancelMeeting 取消会议
func CancelMeeting(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !user.IsLogin() {
		fmt.Fprintf(w, "Please login first!")
		return
	}
	if len(r.Form["title"]) == 0 {
		fmt.Fprintf(w, "please input the title!")
		return
	}
	if len(r.Form["title"]) > 1 {
		fmt.Fprintf(w, "only one title should be input!")
		return
	}
	if err := meeting.CancelMeeting(r.Form["title"][0]); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "meeting %s is cancelled successfully", r.Form["title"][0])
}
