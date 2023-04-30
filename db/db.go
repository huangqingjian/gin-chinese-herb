package db

import (
	"gin-first/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	err error
	sqlDsn string = "sqldsn"
)
// 初始化DB
func InitDB() {
	DB, err = gorm.Open(mysql.Open(config.Global.Mysql.Dsn), &gorm.Config{
		// DryRun: false, //直接运行生成sql而不执行
		Logger: logger.Default.LogMode(logger.Info), // 可以打印SQL
		// QueryFields: true, // 使用表的所有字段执行SQL查询
		// SkipDefaultTransaction: true, // 关闭事务，gorm默认是打开的，关闭可以提升性能
	})
	if err != nil {
		zap.L().Error("Mysql Init Error", zap.Error(err))
		panic(any(err))
	}
}
