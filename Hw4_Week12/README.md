"cloudgo-xorm" 项目，使用 ORM（Object Relational Mapping）库和数据库交互  

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

cloudgo-io-update 项目，在上次作业cloudgo-io的基础上修改  
上次作业cloudgo-io连接：https://github.com/tpisntgod/Service_Calculation/tree/master/cloudgo-io  
将cloudgo-io的处理数据、操作数据库的model部分改成了java 经典的entity-dao-service结构模型。  
