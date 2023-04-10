package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/edu_organization"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EduEnrollmentSearch struct{
    edu_organization.EduEnrollment
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
