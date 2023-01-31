package mw

import (
	"context"

	"BiteDans.com/tiktok-backend/pkg/consts"
	"github.com/cloudwego/hertz/pkg/app"
)

var Key = consts.JWT_KEY

func JwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		c.Next(ctx)
	}
}
