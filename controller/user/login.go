package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginRequest viewmodel.LoginRequest
	if err := ctx.Bind(&loginRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	login, err := service.Login(loginRequest)
	if err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, login)
}
