package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

func UpdatePassword(ctx *gin.Context) {
	var updatePasswordRequest viewmodel.UpdatePasswordRequest
	if err := ctx.Bind(&updatePasswordRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdatePassword(updatePasswordRequest); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
