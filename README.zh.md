## 一个能自我进化的Go微服务框架

其他语言：

### **[English](README.md)**

这是一个Go微服务框架。它是对[jfeng45/servicetmpl](https://github.com/jfeng45/servicetmpl1)的一个重大升级版。

下面是描述它的系列文章:

+ [一个能自我进化的Go微服务框架](https://blog.csdn.net/weixin_38748858/article/details/106996260)

+ [一个非侵入的Go事务管理库--怎样使用](https://blog.csdn.net/weixin_38748858/article/details/106885990)

+ [一个非侵入的Go事务管理库--工作原理](https://blog.csdn.net/weixin_38748858/article/details/106886184)

+ [清晰架构（Clean Architecture）的Go微服务--重大升级](https://blog.csdn.net/weixin_38748858/article/details/107565358)

## 运行

### 安装和设置

不需要完成本节中的所有步骤以使代码运行。 最简单的方法是从github获取代码并运行它，然后在真正需要某些部件时再返回安装。 但是，访问数据库时会遇到错误。
所以，我建议你至少安装一个数据库（MySQL更好），然后大部分代码就都可以运行了。

#### 下载程序

```
go get github.com/jfeng45/servicetmpl1
```

#### 设置MySQL

有两个数据库实现，MySQL和CouchDB，但大多数函数都是在MySQL中实现的。 你最好安装至少其中一个。

```
安装MySQL
在script文件夹中运行SQL脚本以创建数据库和表
```
#### 安装CouchDB

没有它，代码工作正常。创建CouchDB用来完成切换数据库的功能（通过更改配置）。

安装[Windows](https://docs.couchdb.org/en/2.2.0/install/windows.html)

安装[Linux](https://docs.couchdb.org/en/2.2.0/install/unix.html)

安装[Mac](https://docs.couchdb.org/en/2.2.0/install/mac.html)

CouchDB[Example](https://github.com/go-kivik/kivik/wiki/Usage-Examples)

#### 设置CouchDB

```
通过浏览器访问“Fauxton”：http://localhost:5984/_utils/#（使用：admin/admin登录）。
在“Fauxton”中创建新数据库“service_config”。
将以下文档添加到数据库（“_id”和“_rev”由数据库生成，无需更改）：
{
  "_id": "80a9134c7dfa53f67f6be214e1000fa7",
  "_rev": "4-f45fb8bdd454a71e6ae88bdeea8a0b4c",
  "uid": 10,
  "username": "Tony",
  "department": "IT",
  "created": "2018-02-17T15:04:05-03:00"
}
```
#### 安装缓存服务（另一个微服务）

没有它，调用另一个微服务部分将无法正常工作，其余部分工作正常。请按照[reservegrpc](https://github.com/jfeng45/reservegrpc)中的说明设置服务。

### 启动应用程序

#### 启动MySQL
```
cd [MySQLroot]/bin
mysqld
```

#### 启动CouchDB
```
它应该已经启动了
```
#### 启动缓存服务

请按照[reservegrpc](https://github.com/jfeng45/reservegrpc)中的说明启动服务器。

#### 运行main

##### 作为本地应用程序运行

在“main.go”的“main（）”函数中，有两个函数“testMySql（）”和“testCouchDB（）”。
“testMySql（）”从“configs/appConifgDev.yaml”读取配置并访问MySQL。 “testCouchDB（）”从“configs/appConifgProd.yaml”读取配置并访问CouchDB。
“testMySql（）”中有多个函数，你可以通过注释掉其他函数来单独测试一个函数。

```
cd [rootOfProject]/cmd
go run main.go
```
##### 作为gRPC微服务应用程序运行

启动gRPC服务器
```
cd [rootOfProject]/cmd/grpcserver
go run grpcServerMain.go
```
启动gRPC客户端
```
cd [rootOfProject]/cmd/grpcclient
go run grpcClientMain.go
```

### 授权

[MIT](LICENSE.txt) 授权


