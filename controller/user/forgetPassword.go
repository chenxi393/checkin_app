package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

func ForgetPassword(ctx *gin.Context) {
	var forgetPasswordRequest viewmodel.ForgetPasswordRequest
	if err := ctx.Bind(&forgetPasswordRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}

	if err := service.ForgetPassword(forgetPasswordRequest); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
