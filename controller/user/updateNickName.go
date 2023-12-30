package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

// UpdateNickName 修改用户昵称 controller
func UpdateNickName(ctx *gin.Context) {
	var updateNickName viewmodel.UpdateNickNameRequest
	if err := ctx.Bind(&updateNickName); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdateNickName(updateNickName); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
