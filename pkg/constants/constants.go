package constants

const (
	MySQLDefaultDSN     = "gorm:gorm@tcp(localhost:3307)/gorm?charset=utf8&parseTime=True&loc=Local"
	BUCKET_NAME         = "tiktok-bitedans"
	REGION              = "us-west-1"
	PROFILE_PIC_ADDR    = "https://tiktok-bitedans.s3.us-west-1.amazonaws.com/%E5%A4%B4%E5%83%8F.jpg"
	BACKGROUND_PIC_ADDR = "https://tiktok-bitedans.s3.us-west-1.amazonaws.com/%E8%83%8C%E6%99%AF.jpg"
	SIGNATURE           = "喜欢唱、跳、Rap、篮球。"
	FOLLOW_ACTION       = 1
	UNFOLLOW_ACTION     = 0
	POST_COMMENT        = 1
	DELETE_COMMENT      = 2
	LIKE_VIDEO          = 1
	UNLIKE_VIDEO        = 2
)
