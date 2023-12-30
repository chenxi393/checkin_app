package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

// UpdateUserInfo 修改用户信息 controller
func UpdateUserInfo(ctx *gin.Context) {
	var updateUserInfo viewmodel.UpdateUserInfoRequest
	if err := ctx.Bind(&updateUserInfo); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdateUser(updateUserInfo); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
