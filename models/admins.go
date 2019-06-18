package models

import "go-admin/services"

type Admins struct {
	AdminId  int    `xorm:"admin_id"`
	Username string `xorm:"username"`
	Password string `xorm:"password"`
	Status   int    `xorm:"status"`
	RoleId   int    `xorm:"role_id"`
	Rule     string `xorm:"rule"`
	AddTime  int    `xorm:"add_time"`
}

func (admin Admins) GetInfoByUser(UserName string) (*Admins, error) {
	a := new(Admins)
	has, err := services.Engine.Where("username = ?", UserName).Get(a)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return a, nil
}
