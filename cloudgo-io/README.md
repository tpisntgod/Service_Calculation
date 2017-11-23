cloudgo-io  
在上次作业 Hw3_Week9/Web-todolist 的基础上修改。  

增加功能：  
1.支持静态文件服务  
2.支持简单 js 访问  
3.提交表单，并输出一个表格  
4.对 /unknown 给出开发中的提示，返回码 5xx  

项目结构：
![](../Printscreens/fileStruct.png)
assets保存静态文件

具体实现：  
1.支持静态文件服务  
将 path 以 “/” 为前缀的 URL 都定位到 webRoot + "/assets/" 为虚拟根目录的文件系统。
访问用户登录、注册，todoitem增加、删除、查找
![](../Printscreens/staticfile.png)
