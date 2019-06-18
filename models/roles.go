package models

import "go-admin/services"

type Roles struct {
	RoleID int `xorm:"role_id"`
	RoleName string `xorm:"role_name"`
	Rule string `xorm:"rule"`
	Status int `xorm:"status"`
}


func (Roles)GetRule(RoleID int)(rule string){
	//超级管理员不取，返回空
	if RoleID == 1 {
		return
	}

	services.Engine.Where("role_id = ?",RoleID).Cols("rule").Get(rule)

	return
}