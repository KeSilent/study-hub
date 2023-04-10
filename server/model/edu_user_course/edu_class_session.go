// 自动生成模板EduClassSession
package edu_user_course

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// EduClassSession 结构体
type EduClassSession struct {
      global.GVA_MODEL
      EnrollmentId  *int `json:"enrollmentId" form:"enrollmentId" gorm:"column:enrollment_id;comment:报名ID;size:10;"`
      Action  string `json:"action" form:"action" gorm:"column:action;type:enum('add', 'subtract');comment:操作类型（增加或扣除）;"`
      Reason  string `json:"reason" form:"reason" gorm:"column:reason;comment:操作原因;size:255;"`
      NumSessions  *int `json:"numSessions" form:"numSessions" gorm:"column:num_sessions;comment:课时数量;size:10;"`
}


// TableName EduClassSession 表名
func (EduClassSession) TableName() string {
  return "edu_class_session"
}

