package service

import (
	"checkin/model"
	"checkin/pkg/app"
	"checkin/pkg/util"
	"checkin/storage"
	"checkin/viewmodel"
	"strings"

	"go.uber.org/zap"
)

// CreateClass 创建班级 s
func CreateClass(createClassRequest viewmodel.CreateClassRequest) error {
	//查看要创建的班级是否存在
	if className := storage.GetByClassNameMapper(createClassRequest.ClassName); className {
		zap.L().Sugar().Errorf( "创建班级失败")
		return app.ErrClassExist
	}

	//创建班级
	if err := storage.CreateClassMapper(&model.Class{
		ClassId:       util.GetUUID(), //所以生成班级id
		ClassName:     createClassRequest.ClassName,
		ClassCapacity: createClassRequest.Capacity,
		CreateId:      "1", //都默认1
	}); err != nil {
		return err
	}
	return nil
}

// GetAllClass 获取班级列表 service
func GetAllClass(page util.PageRequest) (viewmodel.GetClassListResponse, error) {
	var infos []viewmodel.ClassInfo
	var count int64
	var err error
	if strings.Compare(page.Logo, "register") == 0 {
		//返回班级表里的所有记录数,与所有班级的ID和Name
		infos, count, err = storage.GetAllClassMapper()
	} else if strings.Compare(page.Logo, "universal") == 0 {
		//返回班级表里的记录数,分页获取指定班级列表的ID和Name
		infos, count, err = storage.GetAllClassPageMapper(page.Offset, page.Limit)
	}
	return viewmodel.GetClassListResponse{
		TotalCount: count,
		Class:      infos,
	}, err
}
