package main

import (
	"ddd_demo/config"
	"ddd_demo/internal/application"
	"ddd_demo/internal/infrastructure"
	"ddd_demo/internal/infrastructure/tool/logs"
	"ddd_demo/internal/infrastructure/tool/signals"
	"ddd_demo/internal/interfaces"
	"ddd_demo/internal/interfaces/rpc"
	"ddd_demo/internal/interfaces/web"
)

func main() {
	// 初始化配置
	cfg := config.NewConfig("./config.yaml")

	// 初始化日志
	logs.Init(cfg.Log)

	// 获取 servers, 比如 WebServer, RpcServer
	servers := NewServers(cfg)

	// 启动 servers
	servers.AsyncStart()

	// 优雅退出
	signals.WaitWith(servers.Stop)
}

// NewServers 通过配置文件初始化 Repo 依赖, 然后初始化 App, 最后组装为 Server
// 比如 UserRepo -> UserApp -> WebServer
func NewServers(cfg *config.SugaredConfig) interfaces.ServerInterface {
	repos := infrastructure.NewRepos(cfg)
	apps := application.NewApps(repos)

	servers := interfaces.NewServers()
	servers.AddServer(web.NewWebServer(cfg, apps))
	servers.AddServer(rpc.NewRpcServer(cfg, apps))

	return servers
}
