package edu_organization

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/edu_organization"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    edu_organizationReq "github.com/flipped-aurora/gin-vue-admin/server/model/edu_organization/request"
)

type EduOrganizationService struct {
}

// CreateEduOrganization 创建EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService) CreateEduOrganization(eduOrganization *edu_organization.EduOrganization) (err error) {
	err = global.GVA_DB.Create(eduOrganization).Error
	return err
}

// DeleteEduOrganization 删除EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService)DeleteEduOrganization(eduOrganization edu_organization.EduOrganization) (err error) {
	err = global.GVA_DB.Delete(&eduOrganization).Error
	return err
}

// DeleteEduOrganizationByIds 批量删除EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService)DeleteEduOrganizationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]edu_organization.EduOrganization{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEduOrganization 更新EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService)UpdateEduOrganization(eduOrganization edu_organization.EduOrganization) (err error) {
	err = global.GVA_DB.Save(&eduOrganization).Error
	return err
}

// GetEduOrganization 根据id获取EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService)GetEduOrganization(id uint) (eduOrganization edu_organization.EduOrganization, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eduOrganization).Error
	return
}

// GetEduOrganizationInfoList 分页获取EduOrganization记录
// Author [piexlmax](https://github.com/piexlmax)
func (eduOrganizationService *EduOrganizationService)GetEduOrganizationInfoList(info edu_organizationReq.EduOrganizationSearch) (list []edu_organization.EduOrganization, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&edu_organization.EduOrganization{})
    var eduOrganizations []edu_organization.EduOrganization
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&eduOrganizations).Error
	return  eduOrganizations, total, err
}
