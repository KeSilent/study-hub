package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/edu_organization"
	"github.com/flipped-aurora/gin-vue-admin/server/service/edu_user_course"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup           system.ServiceGroup
	ExampleServiceGroup          example.ServiceGroup
	Edu_organizationServiceGroup edu_organization.ServiceGroup
	Edu_user_courseServiceGroup  edu_user_course.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
