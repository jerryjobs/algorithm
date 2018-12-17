package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/go-ffmt/ffmt"
	"jerrysd.cn/algorithm/tree"
)

func TestNodeTree(t *testing.T) {
	node := &tree.Node{}
	t.Log(node)
	json.Marshal(node)
}

func TestBST(t *testing.T) {
	t.Log("ok")

	btree := tree.NewBSTree(10)
	t.Log(btree)
	btree.Insert(19)
	btree.Insert(20)
	btree.Insert(22)
	btree.Insert(8)
	btree.Insert(9)
	btree.Insert(10)
	btree.Insert(17)
	//ffmt.Print(btree)
	ffmt.Pjson(btree)
	fmt.Println(strings.Repeat("-", 20))
	btree.Remove(19)
	ffmt.Pjson(btree)
}

func TestAVLTree(t *testing.T) {
	t.Log("ok")
	avl := tree.NewAVLTree()

	avl.Insert(2)
	avl.Insert(1)
	avl.Insert(3)
	avl.Insert(4)

	ffmt.Pjson(avl)
	t.Log(avl)
}
