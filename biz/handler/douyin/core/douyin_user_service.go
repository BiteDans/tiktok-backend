// Code generated by hertz generator.

package core

import (
	"BiteDans.com/tiktok-backend/biz/dal/model"
	core "BiteDans.com/tiktok-backend/biz/model/douyin/core"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// RegisterUser .
// @router /douyin/user/register/ [POST]
func RegisterUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserRegisterResponse)
	user := new(model.User)

	if err = model.FindUserByUsername(user, req.Username); err == nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Username has been used"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	user.Username = req.Username
	user.Password = req.Password

	if err = model.RegisterUser(user); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to register user"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User registered successfully"
	resp.UserId = int64(user.ID)
	resp.Token = "token"

	c.JSON(consts.StatusOK, resp)
}

// UserInfo .
// @router /douyin/user [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserResponse)

	if _error := c.Errors.Last(); _error != nil {
		resp.StatusCode = -1
		resp.StatusMsg = _error.Error()
		resp.User = nil

		c.JSON(consts.StatusUnauthorized, resp)
		return
	}

	user := new(model.User)

	if err = model.FindUserById(user, uint(req.UserId)); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "User id does not exist"
		resp.User = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User info retrieved successfully"
	resp.User = &core.User{
		ID:            int64(user.ID),
		Name:          user.Username,
		FollowCount:   123,
		FollowerCount: 456,
		IsFollow:      true,
	}

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /douyin/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserLoginResponse)
	user := new(model.User)

	user.Username = req.Username
	user.Password = req.Password

	var inputpassword string
	var getpassword string

	inputpassword = user.Password

	err, getpassword = model.LoginUser(user)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to login (Username not found)"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	if inputpassword != getpassword {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to login (Incorrect password)"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User login successfully"
	resp.UserId = int64(user.ID)
	resp.Token = "token"

	c.JSON(consts.StatusOK, resp)
}
