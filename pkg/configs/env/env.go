package env

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("pkg/configs/env/.env")

	if err != nil {
		hlog.Error("Failed to load .env, shutting down...")
		panic(err)
	}

	hlog.Info(".env loaded successfully")
}
