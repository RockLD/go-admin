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

func (Nodes) GetList() (n []Nodes,err error) {

	//nds := make(map[int]*Nodes)
	services.InitMysql()
	err = services.Engine.Find(&n)

	if err != nil {
		return n,err
	}
	return n,nil
}