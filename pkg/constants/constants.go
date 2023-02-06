package constants

const (
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3307)/gorm?charset=utf8&parseTime=True&loc=Local"
	BUCKET_NAME     = "tiktok-bitedans"
	REGION          = "us-west-1"
	FOLLOW_ACTION   = 1
	UNFOLLOW_ACTION = 0
	POST_COMMENT    = 1
	DELETE_COMMENT  = 2
)
