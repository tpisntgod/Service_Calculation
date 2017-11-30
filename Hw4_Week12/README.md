使用 ORM（Object Relational Mapping）库  
![](PrintScreens/fileStruct.png)  
创建数据库：  
![](PrintScreens/createdatabase.png)  
创建表：  
![](PrintScreens/createtables.png)  
表的属性的描述：  
![](PrintScreens/describetable.png)  
用 curl POST 用户名和部门名称到网站。数据库有相应的变化：  
![](PrintScreens/post.png)  
用 curl POST 用户名到网站，POST时没有提供部门名称。数据库有相应的变化：  
![](PrintScreens/postnodept.png)  
程序对于POST的错误处理：  
![](PrintScreens/posterrorhandle.png)  
用 curl GET 查询所有用户信息：  
![](PrintScreens/get.png)  
用 curl GET 根据userid查询一个用户信息：  
![](PrintScreens/getsuccess.png)  
程序对于GET的错误处理：  
![](PrintScreens/geterrorhandle.png)  
