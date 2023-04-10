package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/edu_user_course"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EduClassSessionSearch struct{
    edu_user_course.EduClassSession
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
