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
	Children *Tree
}

func CreateMenuTree ( nodes []models.Nodes) (tree []Tree) {
	res := make(map[int]models.Nodes)
	for _,v := range nodes {
		res[v.NodeID] = v
	}
	return tree
}
