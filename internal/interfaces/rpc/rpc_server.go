package rpc

import (
	"ddd_demo/config"
	"ddd_demo/internal/application"
	"ddd_demo/internal/infrastructure/tool/logs"
	"ddd_demo/internal/interfaces"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var _ interfaces.ServerInterface = &RpcServer{}

type RpcServer struct {
	srv  *grpc.Server
	port string
	Apps *application.Apps
}

func (s *RpcServer) SyncStart() {
	logs.Infof("[服务启动] [RPC] 服务地址: %s", s.port)
	go func() {
		l, err := net.Listen("tcp", ":"+s.port)
		if err != nil {
			logs.Fatalf("[服务启动] [RPC] 服务异常: %s", zap.Error(err))
		}

		if err := s.srv.Serve(l); err != nil {
			logs.Fatalf("[服务启动] [RPC] 服务异常: %s", zap.Error(err))
		}
	}()
}

func (s *RpcServer) Stop() {
	logs.Infof("[服务关闭] [rpc] 关闭服务")
	s.srv.GracefulStop()
}

func NewRpcServer(cfg *config.SugaredConfig, apps *application.Apps) interfaces.ServerInterface {
	s := &RpcServer{
		port: cfg.RPC.Port,
		Apps: apps,
	}

	s.srv = grpc.NewServer()

	// 注册路由
	WithRouter(s)

	return s
}
