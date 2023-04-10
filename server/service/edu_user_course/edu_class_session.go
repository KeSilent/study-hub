package edu_user_course

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/edu_user_course"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    edu_user_courseReq "github.com/flipped-aurora/gin-vue-admin/server/model/edu_user_course/request"
)

type EduClassSessionService struct {
}

// CreateEduClassSession 创建EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService) CreateEduClassSession(eduClassSession *edu_user_course.EduClassSession) (err error) {
	err = global.GVA_DB.Create(eduClassSession).Error
	return err
}

// DeleteEduClassSession 删除EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService)DeleteEduClassSession(eduClassSession edu_user_course.EduClassSession) (err error) {
	err = global.GVA_DB.Delete(&eduClassSession).Error
	return err
}

// DeleteEduClassSessionByIds 批量删除EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService)DeleteEduClassSessionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_user_course.EduClassSession{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEduClassSession 更新EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService)UpdateEduClassSession(eduClassSession edu_user_course.EduClassSession) (err error) {
	err = global.GVA_DB.Save(&eduClassSession).Error
	return err
}

// GetEduClassSession 根据id获取EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService)GetEduClassSession(id uint) (eduClassSession edu_user_course.EduClassSession, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduClassSession).Error
	return
}

// GetEduClassSessionInfoList 分页获取EduClassSession记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduClassSessionService *EduClassSessionService)GetEduClassSessionInfoList(info edu_user_courseReq.EduClassSessionSearch) (list []edu_user_course.EduClassSession, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&edu_user_course.EduClassSession{})
    var eduClassSessions []edu_user_course.EduClassSession
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&eduClassSessions).Error
	return  eduClassSessions, total, err
}
