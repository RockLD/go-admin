package models

import "go-admin/services"

type Admins struct {
	AdminId int `json:"admin_id" xorm:"admin_id"`
	Username string `json:"username" xorm:"username"`
	Password string `json:"password" xorm:"password"`
	Status int `json:"status" xorm:"status"`
	RoleId int `json:"role_id" xorm:"role_id"`
	Rule string `json:"rule" xorm:"rule"`
	AddTime int `json:"status" xorm:"add_time"`
}

func (admin Admins) GetInfoByUser(User string)(a Admins,err error)  {
	admin.Username = User
	services.Engine.Get(admin)

	return
}