package web

import (
	"ddd_demo/interfaces/web/handler"
	"ddd_demo/interfaces/web/middleware"
)

func WithRouter(s *WebServer) {
	// 新建 handler
	userHandler := handler.NewUserHandler(s.Apps.UserApp)
	authMiddleware := middleware.NewAuthMiddleware(s.Apps.UserApp)

	// 鉴权
	auth := s.Engin.Group("/auth")
	auth.POST("/login", userHandler.Login)
	auth.POST("/register", userHandler.Register)

	// api
	api := s.Engin.Group("/api")

	// 中间件
	api.Use(authMiddleware.Auth)

	// 路由
	api.GET("/user_info", userHandler.UserInfo)
}
