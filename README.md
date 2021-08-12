# awesome-mall-server-go
> 毛教员讲过，"实践出真知"，说明了真正的知识只有从实践中获得。  
> 习老大说过，"调查研究是谋事之基、成事之道，没有调查，就没有发言权。"
>   
> 所以我们通过"从零搭建电商系统"实现对golang语言的学习，并梳理电商平台的业务流程。
> 
> 

## 介绍
基于gin编写golang版本的电子商城服务端

## 项目目录
#### go.mod 项目依赖包管理
#### main.go 项目主入口
#### db 数据查询文件相关
SQL，即Structured Query Language的缩写，是一种特定目的编程语言，用于管理关系数据库管理系统（RDBMS）。该目录存放相关 sql语句文件，可以通过Navicat等数据库操作软件直接执行。
#### dao 数据访问相关
在传统的多层应用程序中，通常是Web层调用业务层，业务层调用数据访问层。业务层负责处理各种业务逻辑，而数据访问层只负责对数据进行增删改查。

## 软件架构

软件架构说明
编写数据访问层的时候，可以使用DAO模式。DAO即Data Access Object的缩写，即数据访问层 。

#### service 服务
#### route 路由
#### model 数据模型
#### handler 处理工具类

## 安装教程

1. `git clone https://xxx.git` 此项目并使用相关IDE打开(推荐GoLand)；
2. `go mod tidy` 下载包依赖
3. `go run main.go` 执行main函数

## 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

## 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

## 补充说明
#### dll  
初始化脚本，比如创建数据的.sql文件，放至git仓库中便于版本管理。

## 参考链接
1. [Go操作MySQL](https://www.cnblogs.com/nickchen121/p/11517430.html)
2. 

## 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
