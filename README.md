# ddd-demo

基于 ddd 的设计思想, 核心领域需要由纯内存对象+基础设施的抽象的接口组成

* 独立于外部框架: 比如 web 框架可以是 gin, 也可以是 beego
* 独立于客户端: 比如客户端可以是 web, 可以是移动端, 也可以是其他服务 rpc 调用
* 独立于基础组件: 比如数据库可以是 MySQL, 可以是 MongoDB, 甚至是本地文件
* 独立于第三方库: 比如加密库可以是 bcrypt, 也可以是其他加密库, 不会应为第三方库的变更而大幅影响到核心领域
* 可测性: 核心领域的 domain 是纯内存对象, 依赖的基础设施的接口是抽象的, 可以 mock 掉, 便于测试

怎么实现呢? 下面我根据一个案例来一步步展示通过 DDD 重构三层架构的过程

这里是该项目对应的文章: <https://www.yuque.com/dengjiawen8955/dsne7d/un2h14o05e8nsbur>?# 《DDD系列 实战一 应用设计案例 (golang)》

## 快速运行案例

依赖的环境

* golang
* docker
* protobuf

```bash
# 下载项目
git clone git@github.com:dengjiawen8955/ddd_demo.git  && cd ddd_demo
# 准备环境 (启动mysql, redis)
docker-compose up -d
# 准备数据库 (创建数据库, 创建表)
make exec.sql
# 启动项目
make
```

## 核心目录结构

```go
├── internal
│   ├── bill    // 账单业务
│   │   ├── app.go  // 账单 application 层
│   │   ├── model
│   │   │   ├── bill_entity.go // 账单 domain 实体
│   │   │   └── bill_po.go     // 账单持久化对象
│   │   └── repo.go // 账单 repository 层
│   ├── common  // 公共模块
│   │   ├── logs    // 日志
│   │   │   ├── interface.go
│   │   │   └── logger.go
│   │   └── signals // 信号处理
│   │       └── signal.go
│   ├── servers // 服务
│   │   ├── apps.go  // 整合需要的 app
│   │   ├── repos.go // 整合需要的 repo
│   │   ├── rpc     // rpc 服务
│   │   │   ├── proto_file  // proto 文件
│   │   │   ├── protos    // 生成的 proto 代码
│   │   │   ├── rpc_router.go   // rpc 路由
│   │   │   └── rpc_server.go   // rpc 服务
│   │   ├── servers.go  // 整合需要的服务
│   │   └── web    // web 服务
│   │       ├── response        // web 响应封装
│   │       ├── web_router.go   // web 路由
│   │       └── web_server.go   // web 服务
│   └── user    // 用户业务
│       ├── app.go  // 用户 application 层
│       ├── auth_repo.go    // 用户鉴权 repository 层
│       ├── model
│       │   ├── auth_entity.go  // 用户鉴权 domain 实体
│       │   ├── user_dto.go    // 用户 dto (data transfer object), 比如 HTTP 请求的参数
│       │   ├── user_entity.go  // 用户 domain 实体
│       │   └── user_po.go    // 用户持久化对象
│       ├── rate_service.go // 汇率服务
│       ├── repo.go  // 用户 repository 层
│       ├── rpc_server.go   // 用户 rpc 服务
│       ├── transfer_service.go // 转账服务
│       ├── web_auth_middleware.go  // web 鉴权中间件
│       └── web_handler.go  // 用户 web 服务
```

架构模型如下图所示:

![架构图](https://docs.google.com/drawings/d/e/2PACX-1vQ5ps72uaZcEJzwnJbPhzUfEeBbN6CJ04j7hl2i3K2HHatNcsoyG2tgX2vnrN5xxDKLp5Jm5bzzmZdv/pub?w=960&h=657)
