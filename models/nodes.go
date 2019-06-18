package models

import (
	"go-admin/services"
)

type Nodes struct {
	NodeID int `xorm:"node_id"`
	NodeName string `xorm:"node_name"`
	ParentNodeID int `xorm:"parent_node_id"`
	AuthRule string `xorm:"auth_rule"`
	IsMenu int `xorm:"is_menu"`
	Style string `xorm:"style"`
}

func (nodes *Nodes)TableName()string{
	return "nodes"
}
//取完整菜单列表
func (Nodes) GetList() (n []Nodes,err error) {
	services.InitMysql()
	err = services.Engine.Find(&n)

	if err != nil {
		return n,err
	}
	return n,nil
}

//根据角色获取菜单列表
func (ns Nodes) GetListByRole(RoleID int) (n []Nodes,err error){
	if RoleID == 1 {
		 return ns.GetList()
	}



	rule := Roles{}.GetRule(2)
	if rule == "" {
		return n,err
	}

	services.InitMysql()
	err = services.Engine.Find(&n)

	if err != nil {
		return n,err
	}
	return n,nil
}