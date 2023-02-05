package dal

import (
	"BiteDans.com/tiktok-backend/pkg/constants"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		},
	)

	if err != nil {
		hlog.Error("Failed to connect to DB, shutting down...")
		panic(err)
	}

	hlog.Info("Connected to DB")
}
