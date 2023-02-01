package env

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("pkg/configs/env/.env", "pkg/configs/env/.secret")

	if err != nil {
		hlog.Error("Failed to load env files, shutting down...")
		panic(err)
	}

	hlog.Info("Env files loaded successfully")
}
