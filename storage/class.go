package storage

import (
	"checkin/model"
	"checkin/pkg/app"
	"checkin/viewmodel"

	"go.uber.org/zap"
)

// GetByClassNameMapper 查询要创建的班级是否在数据库中存在
func GetByClassNameMapper(className string) bool {
	isClassName := DB.Where("class_name = ?", className).First(&model.Class{})
	if isClassName.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

// CreateClassMapper 创建班级
func CreateClassMapper(class *model.Class) error {
	result := DB.Create(&class)
	if result.Error != nil {
		zap.L().Sugar().Errorf("创建班级：创建班级失败", result.Error)
		return app.InternalServerError
	}
	zap.L().Sugar().Infof("创建班级：成功创建 %v 条记录", result.RowsAffected)
	return nil
}

// GetAllClassPageMapper 分页获取班级列表
func GetAllClassPageMapper(offset, limit int) ([]viewmodel.ClassInfo, int64, error) {
	if limit == 0 {
		limit = 50
	}
	classes := make([]*model.Class, 0)

	responseClassInfo := make([]viewmodel.ClassInfo, limit)
	var count int64
	classCount := DB.Table("class").Count(&count)
	if classCount.Error != nil {
		zap.L().Sugar().Errorf("获取所有用户数量出错", classCount.Error)
		return responseClassInfo, count, app.InternalServerError
	}
	zap.L().Sugar().Infof("获取所有班级：一共有：%v个用户", count)
	result := DB.Select([]string{"class_id", "class_name"}).Offset(offset).Limit(limit).Order("created_at desc").Find(&classes)
	if result.Error != nil {
		zap.L().Sugar().Errorf("分页获取用户信息出错", result.Error)
		return responseClassInfo, count, app.InternalServerError
	}
	for k := range classes {
		responseClassInfo[k].ClassId = classes[k].ClassId
		responseClassInfo[k].ClassName = classes[k].ClassName
	}
	zap.L().Sugar().Info("分页获取班级列表成功")
	return responseClassInfo, count, nil
}

// GetAllClassMapper 获取所有班级列表
func GetAllClassMapper() ([]viewmodel.ClassInfo, int64, error) {
	classes := make([]*model.Class, 0) //注意里面是指针类型
	var count int64
	//1.获取class表中共有多少条记录
	classCount := DB.Table("class").Count(&count)
	if classCount.Error != nil {
		zap.L().Sugar().Errorf("获取所有用户数量出错", classCount.Error)
		return []viewmodel.ClassInfo{}, count, app.InternalServerError //返回内部服务器错误
	}
	zap.L().Sugar().Infof("获取所有班级：一共有：%v个用户", count)

	//2.降序的获取[]*model.Class里的class_id", "class_name"字段,然后内嵌到[]viewmodel.ClassInfo
	responseClassInfo := make([]viewmodel.ClassInfo, count)
	result := DB.Select([]string{"class_id", "class_name"}).Order("created_at desc").Find(&classes)
	if result.Error != nil {
		zap.L().Sugar().Errorf("获取用户信息出错", result.Error)
		return responseClassInfo, count, app.InternalServerError
	}
	for k := range classes {
		responseClassInfo[k].ClassId = classes[k].ClassId
		responseClassInfo[k].ClassName = classes[k].ClassName
	}
	zap.L().Sugar().Info("获取所有班级列表成功")
	return responseClassInfo, count, nil
}
