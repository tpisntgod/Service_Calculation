纯粹使用url和服务器交互的会议管理系统。是有点简陋啦，因为比较完善的项目在隔壁的Web-todolist啦！
简单说一下使用方式，就是在url的query-string中输入对应参数，具体功能和对应格式如下：  
to register an account: http://localhost:8080/register?name=qwe1&password=qwe  
to login an account: http://localhost:8080/login?name=qwe1&password=qwe  
to logout an account: http://localhost:8080/logout  
to create a meeting: http://localhost:8080/cm?title=webmeet1&p=aa&p=cc&stime=2017-10-28 09:30:00&etime=2017-10-28 10:30:00  
to query meetings in an interval: http://localhost:8080/qm?stime=2017-10-01 09:30:00&etime=2017-10-31 10:30:00  
to cancel a meeting that you host: http://localhost:8080/cancelmeeting?title=meet  
