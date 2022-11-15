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
func (u *UserApp) Transfer(fromUserID, toUserID *model.UserID, amount *model.Amount, currencyStr string) error {
	// 读数据
	fromUser := u.userRepo.Get(fromUserID)
	toUser := u.userRepo.Get(toUserID)
	toCurrency := model.NewCurrency(currencyStr)

    // 获取汇率
	rate := u.rateService.GetRate(fromUser.Currency, toCurrency)

	// 转账
	u.transferService.Transfer(fromUser, toUser, amount, rate)

	// 保存数据
	u.userRepo.Save(fromUser)
	u.userRepo.Save(toUser)

	// 保存账单
	bill := &bill_model.Bill{
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
		Amount:     amount,
		Currency:   toCurrency,
	}
	u.billApp.CreateBill(bill)

	return nil
}
```

transferService.Transfer(fromUser, toUser, amount, rate)

```go
func (*TransferService) Transfer(fromUser *User, toUser *User, amount *Amount, rate *Rate) {
	// 通过汇率转换金额
	fromAmount := rate.Exchange(amount)

	// 根据用户不同的 vip 等级, 计算手续费
	fee := fromUser.CalcFee(fromAmount)
    
    // 计算总金额
	fromTotalAmount := fromAmount.Add(fee)

	// 转账
	fromUser.Pay(fromTotalAmount)
	toUser.Receive(amount)
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
