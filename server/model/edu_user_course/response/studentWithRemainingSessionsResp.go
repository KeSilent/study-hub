/*
 * @Author: Yang
 * @Date: 2023-05-05 16:18:50
 * @Description: 请填写简介
 */
package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type StudentWithRemainingSessions struct {
	system.SysUser
	RemainingSessions *int `json:"remainingSessions"`
}
