// Code generated by hertz generator.

package main

import (
	"os"

	"BiteDans.com/tiktok-backend/biz/dal"
	"BiteDans.com/tiktok-backend/biz/dal/model"
	"BiteDans.com/tiktok-backend/pkg/configs/env"
	"BiteDans.com/tiktok-backend/pkg/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func main() {
	h := server.New(
		server.WithMaxRequestBodySize(64 * 1024 * 1024),
	)

	h.Use(mw.AccessLog())

	//set up logger
	f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hlog.SetOutput(f)
	hlog.SetLevel(hlog.LevelTrace)

	// init database
	dal.Init()
	dal.DB.AutoMigrate(&model.User{})
	dal.DB.AutoMigrate(&model.Video{})

	// load .env
	env.LoadEnv()

	register(h)
	h.Spin()
}
