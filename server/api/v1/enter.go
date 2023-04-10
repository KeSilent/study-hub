package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/edu_organization"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/edu_user_course"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup           system.ApiGroup
	ExampleApiGroup          example.ApiGroup
	Edu_organizationApiGroup edu_organization.ApiGroup
	Edu_user_courseApiGroup  edu_user_course.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
