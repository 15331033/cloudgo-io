#### 处理web程序的输入与输出
##### 1.概述
设计一个web小应用,展示静态文件服务、js请求支持、磨板输出、表单处理、Filter中间件设计等方面的能力(不需要数据库支持)
##### 2.任务
编程web应用程序cloudgo-io </br>
1.支持静态文件服务、js请求支持 </br>
路由设置: </br>
![Alt text](./pic/1511101671612.png) </br>
html: </br>
![Alt text](./pic/1511101763008.png) </br>
设置api/test的路由 </br>
![Alt text](./pic/1511103301709.png) </br>
设置对应的handler </br>
![Alt text](./pic/1511103342490.png) </br>
js: </br>
通过ajax访问localhost:8080/api/test,得到返回的json数据.再给网页添加获取到的数据 </br>
![Alt text](./pic/1511101798185.png) </br>
网页显示 </br>
![Alt text](./pic/1511103577021.png) </br>
![Alt text](./pic/1511101825662.png) </br>
![Alt text](./pic/1511101884149.png) </br>
2.添加静态文件访问路径独立前缀 </br>
路由设置: </br>
![Alt text](./pic/1511101723778.png) </br>
网页显示 </br>
![Alt text](./pic/1511104498599.png) </br>
![Alt text](./pic/1511102061235.png) </br>
3.模板输出 </br>
设置模板文件,定义扩展项 </br>
![Alt text](./pic/1511102493227.png) </br>
路由设置: </br>
![Alt text](./pic/1511102251868.png) </br>
定义handler: </br>
![Alt text](./pic/1511102397107.png) </br>
定义模板: </br>
![Alt text](./pic/1511102470326.png) </br>
网页显示: </br>
![Alt text](./pic/1511102765868.png)
4.表单处理 </br>
(1)提交表单页面 </br>
路由设置: </br>
![Alt text](./pic/1511103675196.png) </br>
定义handler: </br>
![Alt text](./pic/1511103737284.png) </br>
定义模板: </br>
![Alt text](./pic/1511103766363.png) </br>
网页显示: </br>
![Alt text](./pic/1511103796693.png) </br>
(2)提交表单并输出表格 </br>
路由设置: </br>
![Alt text](./pic/1511103879768.png)
定义handler: </br>
![Alt text](./pic/1511103960155.png) </br>
定义模板: </br>
![Alt text](./pic/1511104001440.png) </br>
网页显示: </br>
![Alt text](./pic/1511104054098.png) </br>
![Alt text](./pic/1511104069173.png) </br>
5.对/unknown给出开发中的提示,返回码5xx
网页显示: </br>
![Alt text](./pic/1511104148748.png)







