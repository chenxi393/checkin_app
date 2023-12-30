package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

// Register 用户注册 controller
func Register(ctx *gin.Context) {
	var registerRequest viewmodel.RegisterRequest
	if err := ctx.Bind(&registerRequest); err != nil { //返回错误
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.CreateUser(&registerRequest); err != nil {
		app.SendResponse(ctx, err, nil) //返回错误
		return
	}
	app.SendResponse(ctx, nil, nil) //返回成果
}
