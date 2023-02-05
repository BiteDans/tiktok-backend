// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	douyin_core_user "BiteDans.com/tiktok-backend/biz/router/douyin/core/user"
	douyin_core_video "BiteDans.com/tiktok-backend/biz/router/douyin/core/video"
	douyin_extra_follow "BiteDans.com/tiktok-backend/biz/router/douyin/extra/follow"
	hello_example "BiteDans.com/tiktok-backend/biz/router/hello/example"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	douyin_core_video.Register(r)
	douyin_extra_follow.Register(r)
	douyin_core_user.Register(r)

	hello_example.Register(r)
}
