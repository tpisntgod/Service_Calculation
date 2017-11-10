Web-todolist  

主页面mainpage使用了go自带的html/template包，用于生成对应URL返回的html文件。其他的页面就没用template了。  

稍微看了一下REST原则，概括的说就是：用URL定位资源，用HTTP动词（GET,POST,DELETE,DETC）描述操作。
具体一点的话，有一个方面是：URL不包含动词，因为"资源"表示一种实体，所以应该是名词，动词应该放在HTTP协议中。
我的url尽量改成了名词，希望能够符合REST原则。
RESTful感觉是指架构和RESTful API，使得server的API设计满足RESTful要求。  

使用了mysql数据库保存注册的用户和用户对应的todoitem。
具体就是go自带的"database/sql"包和SQL语句啦。  

使用cookie保存的用户登录状态
我看的资料说目前Go标准包没有为session提供任何支持，看了下资料感觉要加上session需要花不少时间。
如果有时间更新下一版本的话会加上session管理部分  

使用了MVP架构（大概吧，我对架构不是很清楚qwq）
