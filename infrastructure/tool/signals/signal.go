package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Allenxuxu/gev/log"
)

// WaitSignal 等待信号, 收到信号后执行回调函数
func WaitWith(stop func()) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof("优雅退出服务...")

	// 优雅退出
	stop()
}
