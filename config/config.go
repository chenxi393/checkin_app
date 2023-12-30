package config

import (
	"strings"
	"log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() error {
	// 初始化配置文件
	if err := initConfig(); err != nil {
		return err
	}
	// 监控配置文件并热加载配置文件
	watchConfig()
	// 初始化日志包
	initLog()
	return nil
}

// 初始化配置文件
func initConfig() error {
	//设置配置文件名称(无扩展名)
	viper.AddConfigPath(".")
	// 设置配置文件的名称，此处不包括配置文件的拓展名
	viper.SetConfigName("config")
	// 设置配置文件格式为YAML格式
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为APISERVER，以下配置可以使程序读取环境变量
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 前边已经配置好一些参数了，现在找到并且读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func watchConfig() {
	// 监控配置文件变化并热加载程序
	viper.WatchConfig()
	//配置文件发生变更之后会调用的回调函数,通过该函数的viper设置，可以使viper监控配置文件变更，如有变更则热更新程序
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("配置文件已更改: %s", in.Name)
	})
}

func initLog() {
	// 先打印到控制台吧
	var logger *zap.Logger
	if viper.GetString("mode") == "debug" {
		logger, _ = zap.NewDevelopment()
	} else {
		logger, _ = zap.NewProduction()
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger) //返回值似乎是一个取消函数
	logger.Info("zap初始化: 成功")
}
