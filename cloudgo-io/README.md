cloudgo-io  
在上次作业 Hw3_Week9/Web-todolist 的基础上修改。  

增加功能：  
1.支持静态文件服务  
2.支持简单 js 访问  
3.提交表单，并输出一个表格  
4.对 /unknown 给出开发中的提示，返回码 5xx  

项目结构：  
![](Printscreens/fileStruct.png)  
assets保存静态文件

具体实现：  
1.支持静态文件服务  
将 path 以 “/” 为前缀的 URL 都定位到 webRoot + "/assets/" 为虚拟根目录的文件系统。  
访问用户登录、注册，todoitem增加、删除、查找的网页是，定位到/assets/对应的html。  
![](Printscreens/staticfile.png)  

2.支持简单 js 访问  
先贴一下所有路由对应的Handler的设置，方便以后解释说明，称下图为图H：  
![](Printscreens/allhandler.png)  
默认界面index.html有一个链接。  
![](Printscreens/indexpage.png)  
该路由对应的处理函数是view.MainPage   
处理函数返回mainpage.html  
![](Printscreens/writeMainpage.png)
mainpage.html中使用js修改class="username"的 "<p>"标签的信息  
![](Printscreens/MainPage.png)
mainpage.js代码，通过ajax异步，GET方法访问/api/mainpage，根据返回的信息修改mainpage中class="username"的"<p>"标签的信息  
![](Printscreens/mainpagejs.png)
根据图H得知，处理路由/api/mainpage的函数是mx.HandleFunc("/api/mainpage", mainPageHandler(formatter)).Methods("GET")
该函数输出了一个匿名结构 ，并使用JSON序列化输出。该json用于js文件更改数据。  
![](Printscreens/apitestfunc.png)

3.提交表单，并输出一个表格  
输出表格的html是使用template写的
TodoList主界面提供了一个查询使用方法或者开发人员信息的功能（Check TodoList information按钮对应功能）  
![](../Printscreens/browsermainpage.png)
点击Check TodoList information按钮进入该页面，有一个表单，服务端根据表单填写的信息，输出一个表格：  
![](Printscreens/CheckTodoListinformation.png)
输出系统使用方法的表格：  
![](Printscreens/systemusage.png)
输出开发者信息的表格：  
![](Printscreens/developerInfo.png)

4.对 /unknown 给出开发中的提示，返回码 5xx  
模仿http包的NotFound函数实现。
![](Printscreens/notimplement.png)
需要注意的是，mx.PathPrefix("/api").Handler(view.NotImplementedHandler())这个语句放置的位置
要放在处理已经实现的功能对应的路由的函数之后，这样才不会覆盖掉已经实现的功能对应的路由处理函数。
