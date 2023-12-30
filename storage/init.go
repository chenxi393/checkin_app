package storage

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Init() {
	DB = openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func openDB(username, password, addr, name string) *gorm.DB {
	var ormLogger logger.Interface
	//根据配置文件设置不同的日志等级
	if viper.GetString("mode") == "debug" {
		ormLogger = logger.Default // 进去看 这里是Warn级别的
	} else { // default 的慢sql是200ms
		ormLogger = logger.Default.LogMode(logger.Info) // Info应该是最低的等级 都会打印
	}
	masterDNS := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		addr,
		name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       masterDNS,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //单数表
		},
	})
	if err != nil {
		zap.L().Sugar().Errorf("数据库连接失败.Database Name: %s   err:%e", name, err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(0)
	return db
}
