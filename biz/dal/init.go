package dal

import (
	"BiteDans.com/tiktok-backend/pkg/consts"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
			Logger:                 logger.Default.LogMode(logger.Info),
		},
	)

	if err != nil {
		hlog.Error("Failed to connect to DB, shutting down...")
		panic(err)
	}

	hlog.Info("Connected to DB")
}
