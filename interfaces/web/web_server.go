package web

import (
	"context"
	"ddd_demo/application"
	"ddd_demo/config"
	"ddd_demo/infrastructure/tool/logs"
	"ddd_demo/interfaces"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebServer struct {
	httpServer *http.Server
	Engin      *gin.Engine
	Apps       *application.Apps
}

func (s *WebServer) SyncStart() {
	logs.Debugf("[服务启动] 服务地址: %s", s.httpServer.Addr)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatalf("[服务启动] http服务异常: %s", zap.Error(err))
		}
	}()
}

func (s *WebServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logs.Debugf("[服务关闭] 关闭服务")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		logs.Fatalf("[服务关闭] 关闭服务异常: %s", zap.Error(err))
	}
}

func NewWebServer(cfg *config.SugaredConfig, apps *application.Apps) interfaces.ServerInterface {
	e := gin.Default()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Web.Port),
		Handler: e,
	}

	s := &WebServer{
		httpServer: httpServer,
		Engin:      e,
		Apps:       apps,
	}

	// 注册路由
	WithRouter(s)

	return s
}
