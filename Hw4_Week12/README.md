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

ab压力测试结果，get方法，查询所有用户信息（一共8条信息）：  
个人感觉性能不是很好。  
ab -n 1000 -c 100 http://localhost:8080/service/userinfo  
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>  
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/  
Licensed to The Apache Software Foundation, http://www.apache.org/  

Benchmarking localhost (be patient)  
Completed 100 requests  
Completed 200 requests  
Completed 300 requests  
Completed 400 requests  
Completed 500 requests  
Completed 600 requests  
Completed 700 requests  
Completed 800 requests  
Completed 900 requests  
Completed 1000 requests  
Finished 1000 requests  


Server Software:          
Server Hostname:        localhost  
Server Port:            8080  

Document Path:          /service/userinfo  
Document Length:        915 bytes  

Concurrency Level:      100  
Time taken for tests:   5.427 seconds  
Complete requests:      1000  
Failed requests:        0  
Total transferred:      1039000 bytes  
HTML transferred:       915000 bytes  
Requests per second:    184.27 [#/sec] (mean)  
Time per request:       542.680 [ms] (mean)  
Time per request:       5.427 [ms] (mean, across all concurrent requests)  
Transfer rate:          186.97 [Kbytes/sec] received  

Connection Times (ms)  
              min  mean[+/-sd] median   max  
Connect:        0    0   0.7      0       3  
Processing:     0  455 1360.2      1    5423  
Waiting:        0  454 1360.2      1    5423  
Total:          0  455 1360.8      1    5426  

Percentage of the requests served within a certain time (ms)  
  50%      1  
  66%      1  
  75%      2  
  80%      3  
  90%   1680  
  95%   4592  
  98%   4661  
  99%   5349  
 100%   5426 (longest request)  
 
cloudgo-io-update 项目，在上次作业cloudgo-io的基础上修改  
上次作业cloudgo-io连接：https://github.com/tpisntgod/Service_Calculation/tree/master/cloudgo-io  
将cloudgo-io的处理数据、操作数据库的model部分改成了java 经典的entity-dao-service结构模型。  