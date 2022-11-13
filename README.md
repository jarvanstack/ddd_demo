# ddd-demo

## 接口

用户UserHandler
* 每个 handler 有完整的 

* 注册 POST /register
* 登录 POST /login
* 获取用户信息 GET /user/:id

课程CourserHandler
* 每个 handler 有完整的 

* 创建课程 POST /create_course

选课PickCourseHandler
* 每个 handler 有完整的 

* 选课 POST /pick_course

转账TransferHandler
* 每个 handler 有完整的 

* 转账 POST /transfer

## 目标

首先支持 http, 然后在领域层不变动的情况下增加 grpc 接口

首先支持 MySQL 储存, 在领域层不变动的情况下更换为 MongoDB

在领域层不变动的情况下增加 redis 缓存

## 目录结构

```bash
├── infrastructure // 外部基础设施层, 实现基础设施层抽象的接口
│   ├── record        // 数据库基础设施, 实现领域核心层的 UserRepo, CourseRepo 等接口
│   └── tool          // 其他工具, 比如 log 等
├── interfaces     // 外部接口层, 提供交互的接口
│   ├── web            // web 服务
│   └── servers.go     // 整合所有 server, 比如 gin 和 rpc
└── internal       // 领域核心层 (无状态)
    ├── application    // 应用 (类似MVC中无状态的service)
    ├── domain         // 领域 (封装数据校验和无状态的逻辑)
    └── repository     // 基础设施层抽象的接口
```

目录结构代表的部分如下图所示:

![洋葱模型](https://markdown-1304103443.cos.ap-guangzhou.myqcloud.com/2022-02-0420221111220700.png)

## 执行顺序

main()

* config.New() *config.Config 初始化配置
* interfaces.NewApps(cfg) *interfaces.Apps 获取所有 Repo->App
* servers := interfaces.NewServers(ctx) interfaces.Server (Ctx 不需要了, App 占有部分 repo 就行了)
* servers.SyncRun() 启动所有 server
* singnal.NotifyWith(servers.Stop) 优雅启停, 停止所有 interfaces

interfaces.NewServers(ctx)

* 包含多个服务, 比如 gin 和 rpc
* 每个 gin 包含多个 handler, 比如 UserHandler, CourseHandler
* 每个 handler 有完整的 *interfaces.Ctx
* UserHandler.Login() 调用 UserApp.Login(entity.Login)