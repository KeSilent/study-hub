package edu_user_course

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/edu_user_course"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    edu_user_courseReq "github.com/flipped-aurora/gin-vue-admin/server/model/edu_user_course/request"
)

type EduEnrollmentService struct {
}

// CreateEduEnrollment 创建EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService) CreateEduEnrollment(eduEnrollment *edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Create(eduEnrollment).Error
	return err
}

// DeleteEduEnrollment 删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService)DeleteEduEnrollment(eduEnrollment edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Delete(&eduEnrollment).Error
	return err
}

// DeleteEduEnrollmentByIds 批量删除EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService)DeleteEduEnrollmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_user_course.EduEnrollment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEduEnrollment 更新EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService)UpdateEduEnrollment(eduEnrollment edu_user_course.EduEnrollment) (err error) {
	err = global.GVA_DB.Save(&eduEnrollment).Error
	return err
}

// GetEduEnrollment 根据id获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService)GetEduEnrollment(id uint) (eduEnrollment edu_user_course.EduEnrollment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduEnrollment).Error
	return
}

// GetEduEnrollmentInfoList 分页获取EduEnrollment记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduEnrollmentService *EduEnrollmentService)GetEduEnrollmentInfoList(info edu_user_courseReq.EduEnrollmentSearch) (list []edu_user_course.EduEnrollment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&edu_user_course.EduEnrollment{})
    var eduEnrollments []edu_user_course.EduEnrollment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&eduEnrollments).Error
	return  eduEnrollments, total, err
}
