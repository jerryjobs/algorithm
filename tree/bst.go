package tree

// nodeBST
// 二叉排序树（Binary Sort Tree）又称二叉查找树。
// 它或者是一棵空树；或者是具有下列性质的二叉树：
// （1）若左子树不空，则左子树上所有结点的值均小于它的根结点的值；
// （2）若右子树不空，则右子树上所有结点的值均大于它的根结点的值；
// （3）左、右子树也分别为二叉排序树；
type nodeBST struct {
	Key                 int
	Left, Right, parent *nodeBST
}

// NewBSTree init bstree
func NewBSTree(key int) *nodeBST {
	return &nodeBST{
		Key: key,
	}
}

// Max max of node bst
func (t *nodeBST) Max() *nodeBST {
	if t.Left == nil && t.Right == nil {
		return t
	}
	if t.Right != nil {
		return t.Right.Max()
	} else {
		return t.Right
	}
}

// Min of node bst
func (t *nodeBST) Min() *nodeBST {
	if t.Left == nil && t.Right == nil {
		return t
	}
	if t.Left != nil {
		return t.Left.Min()
	} else {
		return t.Right.Min()
	}
}

/**
为什么先实现搜索呢？
一般BST里面没有重复的元素，你增添或者删除元素，都必须要先查找一下，
看有没有呀，所以BST搜索要先实现，这个搜索是很简单的，慢慢看我讲解吧,
*/

// Search search by key
// if find then ok is true
func (t *nodeBST) Search(key int) (ok bool, node *nodeBST) {
	if key < t.Key {
		if t.Left == nil {
			return false, nil
		}
		return t.Left.Search(key)
	} else if key > t.Key {
		if t.Right == nil {
			return false, nil
		}
		return t.Right.Search(key)
	} else {
		return true, t
	}
}

//Insert insert new key
func (t *nodeBST) Insert(key int) *nodeBST {
	// if t.parent == nil {
	// 	t.Key = key
	// 	return t
	// }
	if key < t.Key {
		if t.Left == nil {
			t.Left = NewBSTree(key)
			t.Left.parent = t
			return t.Left
		} else {
			return t.Left.Insert(key)
		}
	} else if key > t.Key {
		if t.Right == nil {
			t.Right = NewBSTree(key)
			t.Right.parent = t
			return t.Right
		} else {
			t.Right.Insert(key)
		}
	}
	return t
}

// Remove remove by
// 删除总共有3种情况，
// 1，只有左子树；
// 2，只有右子树；
// 3，左右子树都有。
func (t *nodeBST) Remove(key int) *nodeBST {
	if ok, node := t.Search(key); ok {

		if node.Left != nil && node.Right == nil {
			max := node.Left.Max()
			node.Key = max.Key
			max.parent.Left = nil
			max.parent = nil
			return node
		} else if node.Left == nil && node.Right != nil {
			min := node.Right.Min()
			node.Key = min.Key
			min.parent.Right = nil
			min.parent = nil
			return node
		} else if node.Left == nil && node.Right == nil {
			if node.parent.Left == node {
				node.parent.Left = nil
			} else {
				node.parent.Right = nil
			}
			node.parent = nil
			return node
		} else {
			min := node.Left.Min()
			node.Key = min.Key
			if min.parent.Left == min {
				min.parent.Left = nil
			} else {
				min.parent.Right = nil
			}
			min.parent = nil
			return min
		}
	}
	return nil
}
