# ddd-demo

## 目标

(1) 基于 ddd 创建一个 web server, 提供如下接口

* 注册 POST /auth/register
* 登录 POST /auth/login
* 获取用户信息 GET /user
* 转账 POST /transfer 支持跨币种转账

(2) 在核心领域层不变动的情况下增加 grpc server (一般可用于后台的操作)

(3) 在核心领域层不变动的情况下根据配置跟换 redis session 或者 jwt 认证

## 核心目录结构

```go
├── internal
│   ├── application    // 应用层 (领域核心) 类似MVC中无状态的service
│   │   ├── apps.go
│   │   └── user.go
│   ├── domain         // 领域层 (领域核心) 封装数据校验和无状态的逻辑
│   │   ├── auth.go              // 验证领域对象
│   │   ├── repository           // 基础设施层抽象的接口定义
│   │   ├── user_domain.go       // 用户领域对象
│   │   ├── user_dto.go          // 用户数据传输对象
│   │   └── user_po.go           // 用户持久化对象
│   ├── infrastructure  // 基础设施层, 实现基础设施层抽象的接口
│   │   ├── auth                 // Auth 基础设施, 实现领域核心层的 AuthRepo 接口
│   │   ├── persistence          // 持久化基础设施
│   │   ├── repos.go
│   │   └── tool    
│   └── interfaces      // 外部接口层, 提供交互的接口
│       ├── rpc                 // rpc 服务
│       ├── servers.go          // 服务启动
│       └── web                 // web 服务
```

架构模型如下图所示:

![架构图](https://docs.google.com/drawings/d/e/2PACX-1vQ5ps72uaZcEJzwnJbPhzUfEeBbN6CJ04j7hl2i3K2HHatNcsoyG2tgX2vnrN5xxDKLp5Jm5bzzmZdv/pub?w=960&h=657)

## 环境准备

* golang
* docker
* protobuf

## 项目启动

```bash
# 准备环境 (启动mysql, redis)
docker-compose up -d
# 准备数据库 (创建数据库, 创建表)
make init
# 启动项目
make
```

## 转账

一个用户可以转账给另一个用户, 支持跨币种转账

