package common

import (
	"crypto/md5"
	"encoding/hex"
	"go-admin/models"
)

func MyMd5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

type Tree struct {
	NodeID int
	NodeName string
	ParentNodeID int
	AuthRule string
	IsMenu int
	Style string
	Children []Tree
}

func CreateMenuTree ( nodes []models.Nodes) (tree []Tree) {
	res := make(map[int]*Tree)
	for _,v := range nodes {
		res[v.NodeID] = &Tree{NodeID:v.NodeID,
			NodeName:v.NodeName,
			ParentNodeID:v.ParentNodeID,
			AuthRule:v.AuthRule,
			IsMenu:v.IsMenu,
			Style:v.Style,
			Children:[]Tree{},
		}
	}
	for index,value := range res {
		if value.ParentNodeID != 0 {
			res[value.ParentNodeID].Children = append(res[value.ParentNodeID].Children,*res[index])
		}
	}
	for _,val := range res {
		if val.ParentNodeID == 0 {
			tree = append(tree,*val)
		}
	}
	return tree
}
