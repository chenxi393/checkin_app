package main

import (
	"checkin/config"
	"checkin/router"
	"checkin/router/middleware"
	"checkin/storage"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置文件并监控配置文件变化进行热加载程序
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 数据库初始化并建立连接
	storage.Init()

	gin.SetMode(viper.GetString("mode"))
	g := gin.Default()
	router.Load(
		g,
		middleware.Cors(),
		// 加入多个中间件...
	)

	// ping 服务器以确保路由正常工作(健康检查)
	go func() {
		if err := pingServer(); err != nil {
			zap.L().Sugar().Fatal("路由没有响应，或者启动时间过长.", err)
		}
		zap.L().Sugar().Info("路由启动成功.")
		zap.L().Sugar().Infof("开始监听 http 地址上的传入请求: %s", viper.GetString("addr"))
	}()
	zap.L().Sugar().Info(g.Run(viper.GetString("addr")).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// 通过向 /health 发送 GET 请求来 Ping 服务器。
		resp, err := http.Get(viper.GetString("url") + "/api/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// 休眠一秒钟以继续下一次 ping。
		zap.L().Sugar().Info("等待路由，1秒后重试.")
		time.Sleep(time.Second)
	}
	return errors.New("无法连接到路由")
}
