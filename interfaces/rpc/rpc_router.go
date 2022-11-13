package rpc

import (
	"ddd_demo/interfaces/rpc/protos/in/user"
	"ddd_demo/interfaces/rpc/services"
)

func WithRouter(s *RpcServer) {
	// 新建 server
	userServer := services.NewUserServer(s.Apps.UserApp)

	// 注册路由
	user.RegisterUserServer(s.srv, userServer)
}
