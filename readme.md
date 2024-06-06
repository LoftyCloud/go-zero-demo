# **go-zero项目说明文档**

## **一**. **技术栈**

1. go-zero框架

go-zero是对grpc的一种封装，go-zero提供了缓存、熔断等方面的处理，使得开发过程中只需关心业务数据。架构如图1所示，其中rpc（远程服务调用）层在项目比较小时可以不使用。

1. 说明文档：http://go-zero.dev/

2. 例子：https://github.com/zeromicro/zero-examples

![图1 go-zero架构](https://github.com/LoftyCloud/go-zero-demo/blob/main/images/image-20240605174422190.png)



1）缓存：查数据时先查找缓存，若没有再查找数据库，若数据库中存在则会将数据写入缓存。并发量大时容易出现脏数据。**缓存击穿，缓存雪崩**问题

2）熔断，降载：当服务器压力过大时禁止访问。

2. redis

Redis（REmote DIctionary Server，远程字典服务）是一个开源的使用ANSI C语言编写、遵守BSD协议、支持网络、可基于内存、分布式、可选持久性的Key-Value存储数据库，并提供多种语言的API。Redis通常被称为数据结构服务器，因为值（value）可以是字符串(String)、哈希(Hash)、列表(list)、集合(sets)和有序集合(sorted sets)等类型。

参考：https://www.runoob.com/redis/redis-commands.html

## **2.** **实际项目开发**

**1.** **安装go-zero环境**

1）安装goctl（go contorl）：它是go-zero的一个工具，可以用于**自动生成代码。**

说明文档：https://go-zero.dev/docs/tasks/installation/goctl

go install [github.com/zeromicro/go-zero/tools/goctl@latest](mailto:github.com/zeromicro/go-zero/tools/goctl@latest)

```bash
goctl –help
```

2）安装propoc：protoc是一个用于生成代码的工具，它可以根据proto文件生成C++、Java、Python、Go、PHP等多重语言的代码，而gRPC的代码生成还依赖protoc-gen-go，protoc-gen-go-grpc插件来配合生成go语言的gRPC代码。

安装goctl后使用命令自动安装protoc，protoc-gen-go和protoc-gen-go-grpc：

```bash
goctl env check -i -f
```

**2. goctl命令**

1）goctl api：基于api文件生成代码，随后只需在生成的service中编写逻辑即可。参见api-demo生成：https://go-zero.dev/docs/tasks/cli/api-demo。需要注意的时，在对.api文档进行修改后需要重新执行go api自动化生成代码命令。

```bash
goctl api go -api 【user.api】 -dir ../ -style goZero
```

![图2 .api文档说明，基于此文档生成代码](https://github.com/LoftyCloud/go-zero-demo/blob/main/images/image-20240605174506041.png)

![图2 .api文档说明，基于此文档生成代码](https://github.com/LoftyCloud/go-zero-demo/blob/main/images/image-20240605174537524.png)

2）goctl model：与java项目单独使用mapper层处理数据库不同，go-zero将操作数据库层与model结合到了一起。goctl model支持从数据库建表sql语句（goctl model mysql ddl）以及已建立好的sql数据表生成model模型（goctl model mysql datasource）。建立好的模型中会自带Insert()、FindOne()等函数方便调用，也可以额外编写其他的数据库操作函数。

~~~bash
goctl model mysql ddl --src 【user.sql】 --dir .

goctl model mysql datasource -url="【username】:【passwd】@tcp(【host】:【port】)/【dbname】" -table="【tables】" -dir="【modeldir】" 【-cache=true】 --style=goZero
~~~

使用设置-cache字段选择创建数据表模型架构时是否使用缓存，使用sqlx和sqlc分别管理数据库和缓存。

redis缓存：(1）防止数据库被击穿，特别的，当缓存为空时将插入一个“*”作为占位符。(2）缓存中查不到数据将继续去db查。(3）update或delete操作在数据库中执行完毕后，需要将缓存中对应key值的记录删除。

参见goctl model：https://go-zero.dev/docs/tutorials/cli/model

**感悟：**如果代码运行逻辑与自己设想的不相符，那么控制台的输出信息很值得关注。调用新建用户时数据库修改user表和userdata表，但是不管怎么修改代码逻辑还是返回失败，而且就好像根本不尝试执行一样。经过长时间的检查，我发现sqlx在控制台的输出了Error on getting sql instance错误，但是没有用红色进行标注所以一直被遗漏了。通过检查发现是连接数据库的用户名写错了，经过修改后正常。

3）goctl rpc：远程服务调用，使用以下命令生成rpc代码。需了解protobuf与rpc。

```bash
goctl rpc protoc 【user.proto】 -go_out ../ -go-grpc_out ../ -zrpc_out ../ -style goZero
```

4）goctl docker：当某个服务写好后，使用以下命令将服务多阶段打包到镜像，这会在当前路径下生成一个dockerfile文件。

~~~bash
goctl docker -go 【file.go】
~~~

5）goctl kube：生成k8s部署文件。这会在当前路径下生成一个yaml文件。

中间件：在调用方法前做拦截
