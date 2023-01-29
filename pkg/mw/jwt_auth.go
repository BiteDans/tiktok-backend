package mw

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

var Key = "secret_key"

func JwtMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {

		c.Next(ctx)
	}
}
