# 任务清单-后端

**此项目使用`Gin`+`Gorm` ，基于`RESTful API`实现的一个任务清单**。



## 项目功能实现

* 登录鉴权（jwt-go）

* 清单 crud

* 分页



## 项目架构

```
TodoList/
├── api
├── conf
├── middleware
├── model
├── pkg
│  ├── util
├── routes
├── serializer
└── service
```

* api : 用于定义接口函数

* conf : 用于存储配置文件

* middleware : 应用中间件

* model : 应用数据库模型

* pkg / util : 工具函数r

* routers : 路由逻辑处理

* serializer : 将数据序列化为 json 的函数

* service : 接口函数的实现



## 项目接口

```bash
[GIN-debug] POST   /api/v1/user/register     --> TodoList/api.Register (3 handlers)
[GIN-debug] POST   /api/v1/user/login        --> TodoList/api.Login (3 handlers)
[GIN-debug] POST   /api/v1/task              --> TodoList/api.CreateTask (4 handlers)
[GIN-debug] GET    /api/v1/task/:id          --> TodoList/api.ShowTask (4 handlers)
[GIN-debug] GET    /api/v1/tasks             --> TodoList/api.ListTask (4 handlers)
[GIN-debug] PUT    /api/v1/task/:id          --> TodoList/api.UpdateTask (4 handlers)
[GIN-debug] DELETE /api/v1/task/:id          --> TodoList/api.DeleteTask (4 handlers)

```



## 项目运行

```bash
go mod tidy
```

```bash
go run main.go
```
















































