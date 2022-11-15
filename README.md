# ddd-demo

基于 ddd 的设计思想, 核心领域需要由纯内存对象+基础设施的抽象的接口组成

* 独立于外部框架: 比如 web 框架可以是 gin, 也可以是 beego
* 独立于客户端: 比如客户端可以是 web, 可以是移动端, 也可以是其他服务 rpc 调用
* 独立于基础组件: 比如数据库可以是 MySQL, 可以是 MongoDB, 也可以是其他存储
* 独立于第三方库: 比如加密库可以是 bcrypt, 也可以是其他加密库, 不会应为第三方库的变更而大幅影响到核心领域
* 可测性: 核心领域的 domain 是纯内存对象, 依赖的基础设施的接口是抽象的, 可以 mock 掉, 便于测试

软件架构的通用方法

> 加一层, 如果实在不行, 再加一层

## 目标

创建一个 web server, 提供如下接口

* 注册 POST /auth/register
* 登录 POST /auth/login
* 获取用户信息 GET /user
* 转账 POST /transfer

### 注册

* 目前通过账号和密码注册, 以后可能增加根据手机号邮箱等注册
* 目前保存数据使用的 MySQL, 以后可能使用其他数据库

### 登录

同理注册

* 目前使用账号密码登录, 以后可能增加根据手机号邮箱等登录

### 鉴权

除了注册登录, 其他接口都需要鉴权

* 目前使用 redis 鉴权, 以后可能会更换为 jwt 鉴权, 需要有切换的能力

### 转账

一个用户转账给另一个用户

* 需要支持跨币种转账
* 目前转账的汇率从第三方 (微软的 api) 获取, 以后可能会考虑变更或者做缓存
* 目前转账是不收取手续费的, 以后可能根据用户 vip 等级收取不同的手续费
* 需要保存账单, 以便审计和对账用
* 目前账单是保存在 MySQL 中, 以后可能会考虑保存到其他数据库或者消息队列消费

### 接口

* 目前提供 web 接口, 以后可能会提供 rpc 或者其他接口

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

## 转账设计

UserA 转账给 UserB 1000 CNY

(1) 领域核心设计如下(省略错误处理):

```go
func (UserAppInterface) UserApp.Transfer(formUserID *UserID, toUserID *UserID, amount *Amount, AmountStr string) {
    // 读数据
    fromUser := userRepo.Get(formUserID)
    toUser := userRepo.Get(toUserID)

    // 获取汇率
    toAmount := NewAmount(AmountStr)
    rate := RateService.GetRate(fromUser.Amount, toAmount)

    // 转账
    transferService.Transfer(fromUser, toUser, amount, rate)

    // 保存数据
    userRepo.Save(fromUser)
    userRepo.Save(toUser)

    // 保存账单
    bill := NewBill(fromUser, toUser, amount, rate)
    billRepo.Save(bill)
}
```

transferService.Transfer(fromUser, toUser, amount, rate)

```go
func (*TransferService) Transfer(fromUser *User, toUser *User, amount *Amount, rate *Rate) {
    // 通过汇率转换金额
    fromAmount := rate.Exchange(amount)

    // 根据用户不同的 vip 等级, 计算手续费
    fee := fromUser.CalcFee(fromAmount)

    // 转账
    fromUser.Amount.Sub(fromAmount.Add(fee))
    toUser.Amount.Add(amount)
}
```

(2) 设计 repo 接口

```go
type RateService interface {
    GetRate(from *Amount, to *Amount) *Rate
}

type TransferService interface {
    Transfer(fromUser *User, toUser *User, amount *Amount, rate *Rate)
}

type BillRepo interface {
    Save(bill *Bill)
}
```

(3) 设计 domain 类

```go
type Rate struct {
    rate decimal.Decimal
}

func (r *Rate) Exchange(amount *Amount) *Amount {
    return NewAmount(amount.Value().Mul(r.rate))
}

type Bill struct {
    ID        *BillID
    FromUser  *User
    ToUser    *User
    Amount    *Amount
    Rate      *Rate
    CreatedAt time.Time
}
```
