package env

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env", ".secrets")

	if err != nil {
		hlog.Error("Failed to load env files, shutting down...")
		panic(err)
	}

	hlog.Info("Env files loaded successfully")
}
