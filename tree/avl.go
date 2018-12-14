package tree

// AVL tree

/**
高度为 h 的 AVL 树，节点数 N 最多2^h − 1； 最少 （其中）。
最少节点数 n 如以斐波那契数列可以用数学归纳法证明：
Nh=F【h+ 2】 - 1 (F【h+ 2】是 Fibonacci polynomial 的第h+2个数）。
即：
N0 = 0 （表示 AVL Tree 高度为0的节点总数）
N1 = 1 （表示 AVL Tree 高度为1的节点总数）
N2 = 2 （表示 AVL Tree 高度为2的节点总数）
Nh=N【h− 1】 +N【h− 2】 + 1 （表示 AVL Tree 高度为h的节点总数）
换句话说，当节点数为 N 时，高度 h 最多为

// 节点的平衡因子是它的右子树的高度减去它的左子树的高度。
// 带有平衡因子 1、0 或 -1 的节点被认为是平衡的。带有平衡因子 -2 或 2 的节点被认为是不平衡的，
// 并需要重新平衡这个树。平衡因子可以直接存储在每个节点中，或从可能存储在节点中的子树高度计算出来。\


L
*/

// AVLTreeType avl tree
type AVLTreeType interface {
	Tree
	LL()
	LR()
	RL()
	RR()
}

// AVLTree struct of tree node
type avlTree struct {
	nodeBST
	Left, Right *avlTree
	height      int
}

func NewAVLTree() *avlTree {
	return &avlTree{
		nodeBST: nodeBST{},
	}
}
