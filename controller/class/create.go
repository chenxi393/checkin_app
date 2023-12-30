package class

import (
	"checkin/pkg/app"
	"checkin/service"
	"checkin/viewmodel"

	"github.com/gin-gonic/gin"
)

// Create 创建班级 controller
func Create(ctx *gin.Context) {
	var createClassRequest viewmodel.CreateClassRequest
	if err := ctx.Bind(&createClassRequest); err != nil {
		app.SendResponse(ctx, app.ErrBind, nil)
		return
	}
	if err := service.CreateClass(createClassRequest); err != nil {
		app.SendResponse(ctx, err, nil)
		return
	}
	app.SendResponse(ctx, nil, nil)
}
