package models

import "go-admin/services"

type Nodes struct {
	NodeID int `json:"node_id" xorm:"node_id"`
	NodeName string `json:"node_name" xorm:"node_name"`
	ParentNodeID int `json:"parent_node_id" xorm:"parent_node_id"`
	AuthRule string `json:"auth_rule" xorm:"auth_rule"`
	IsMenu int `json:"is_menu" xorm:"is_menu"`
	Style string `json:"style" xorm:"style"`
}

func (nodes *Nodes)TableName()string{
	return "nodes"
}

func (Nodes) GetList() (n map[int]*Nodes,err error) {
	nds := make(map[int]*Nodes)
	services.InitMysql()
	err = services.Engine.Find(&nds)

	if err != nil {
		return n,err
	}

	return nds,nil
}