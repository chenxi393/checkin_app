package user

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

// UpdateEmail 修改用户邮箱 controller
func UpdateEmail(ctx *gin.Context) {
	var updateEmail viewmodel.UpdateEmailRequest
	if err := ctx.Bind(&updateEmail); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.UpdateEmail(updateEmail); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
