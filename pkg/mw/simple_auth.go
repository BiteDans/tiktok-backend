package mw

import (
	"context"

	core "BiteDans.com/tiktok-backend/biz/model/douyin/core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var err error
		var req core.DouyinUserRequest
		err = c.BindAndValidate(&req)
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}

		if req.Token != "token" {
			c.Error(errors.NewPublic("User is not logged in!"))
			return
		}
	}
}
