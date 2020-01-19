package tree

type BinaryTree struct {
	Data        interface{}
	Left, Right *BinaryTree
}

func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

//递归-前序遍历
//规则是若二叉树为空，则空操作返回，否则先访问根结点，然后前序遍历左子树，再前序遍历右子树
func (binaryTree *BinaryTree) PreOrderRecursion() []interface{} {
	result, tmpLeft, tmpRight := make([]interface{}, 0), make([]interface{}, 0), make([]interface{}, 0)
	if binaryTree == nil {
		return result
	}
	result = append(result, binaryTree.Data)
	if binaryTree.Left != nil {
		tmpLeft = binaryTree.Left.PreOrderRecursion()
		result = append(result, tmpLeft...)
	}
	if binaryTree.Right != nil {
		tmpRight = binaryTree.Right.PreOrderRecursion()
		result = append(result, tmpRight...)
	}
	return result
}

//非递归-前序遍历(根->左->右)
func (binaryTree *BinaryTree) PreOrder() []interface{} {
	btSlice := make([]*BinaryTree, 0)
	result := make([]interface{}, 0)
	if binaryTree == nil {
		return result
	}
	btSlice = append(btSlice, binaryTree)
	for len(btSlice) > 0 {
		curTreeNode := btSlice[len(btSlice)-1]
		btSlice = btSlice[:len(btSlice)-1]
		result = append(result, curTreeNode.Data)
		if curTreeNode.Right != nil {
			btSlice = append(btSlice, curTreeNode.Right)
		}
		if curTreeNode.Left != nil {
			btSlice = append(btSlice, curTreeNode.Left)
		}
	}
	return result
}

//递归-后序遍历:
//规则是若树为空，则空操作返回，否则从左到右先叶子后结点的方式遍历访问左右子树，最后是访问根结点
func (binaryTree *BinaryTree) AfterOrderRecursion() []interface{} {
	result, tmpLeft, tmpRight := make([]interface{}, 0), make([]interface{}, 0), make([]interface{}, 0)
	if binaryTree == nil {
		return result
	}
	if binaryTree.Left != nil {
		tmpLeft = binaryTree.Left.AfterOrderRecursion()
		result = append(result, tmpLeft...)
	}
	if binaryTree.Right != nil {
		tmpRight = binaryTree.Right.AfterOrderRecursion()
		result = append(result, tmpRight...)
	}
	result = append(result, binaryTree.Data)
	return result
}

//非递归-后序遍历
func (binaryTree *BinaryTree) AfterOrder() []interface{} {
	btSlice := make([]*BinaryTree, 0)
	tmpSlice := make([]*BinaryTree, 0)
	result := make([]interface{}, 0)

	if binaryTree == nil {
		return result
	}
	btSlice = append(btSlice, binaryTree)
	for len(btSlice) > 0 {
		curTreeNode := btSlice[len(btSlice)-1]
		btSlice = btSlice[:len(btSlice)-1]
		tmpSlice = append(tmpSlice, curTreeNode)
		if curTreeNode.Left != nil {
			btSlice = append(btSlice, curTreeNode.Left)
		}
		if curTreeNode.Right != nil {
			btSlice = append(btSlice, curTreeNode.Right)
		}
	}
	for len(tmpSlice) > 0 {
		curTmpSlice := tmpSlice[len(tmpSlice)-1]
		tmpSlice = tmpSlice[:len(tmpSlice)-1]
		result = append(result, curTmpSlice.Data)
	}
	return result
}

//递归-中序遍历
//规则是若树为空，则空操作返回，否则从根结点开始(注意并不是先访问根结点) ，中序遍历根结点的左子树，然后是访问根结点，最后中序遍历右子树
func (binaryTree *BinaryTree) MiddleOrderRecursion() []interface{} {
	result, tmpLeft, tmpRight := make([]interface{}, 0), make([]interface{}, 0), make([]interface{}, 0)
	if binaryTree == nil {
		return result
	}
	if binaryTree.Left != nil {
		tmpLeft = binaryTree.Left.MiddleOrderRecursion()
		result = append(result, tmpLeft...)
	}
	result = append(result, binaryTree.Data)
	if binaryTree.Right != nil {
		tmpRight = binaryTree.Right.MiddleOrderRecursion()
		result = append(result, tmpRight...)
	}
	return result
}

//非递归-中序遍历(左-根-右)
func (binaryTree *BinaryTree) MiddleOrder() []interface{} {
	btSlice := make([]*BinaryTree, 0)

	result := make([]interface{}, 0)
	if binaryTree == nil {
		return result
	}
	curTreeNode := binaryTree
	for len(btSlice) > 0 || curTreeNode != nil {
		if curTreeNode != nil {
			btSlice = append(btSlice, curTreeNode)
			curTreeNode = curTreeNode.Left
		} else {
			tmp := btSlice[len(btSlice)-1]
			result = append(result, tmp.Data)
			btSlice = btSlice[:len(btSlice)-1]
			curTreeNode = tmp.Right
		}
	}
	return result
}

//层序遍历
//规则是若树为空，则返回空，否则从树的第一层，也就是根节点开始访问，从上而下逐层遍历，在同一层中，按从左到右的顺序对节点逐个访问
func (binaryTree *BinaryTree) LayersOrder() [][]interface{} {
	result := make([][]interface{}, 0)
	btSlice := make([]*BinaryTree, 0)
	if binaryTree == nil {
		return result
	}
	btSlice = append(btSlice, binaryTree)
	for len(btSlice) > 0 {
		currLevel := make([]interface{}, 0)
		allTreeNode := btSlice[:]
		btSlice = []*BinaryTree{}
		for _, node := range allTreeNode {
			currLevel = append(currLevel, node.Data)
			if node.Left != nil {
				btSlice = append(btSlice, node.Left)
			}
			if node.Right != nil {
				btSlice = append(btSlice, node.Right)
			}
		}
		result = append(result, currLevel)
	}
	return result;
}

//递归-层数-对任意一个子树的根节点来说，它的深度=左右子树深度的最大值+1
func (binaryTree *BinaryTree) LayersRecursion() int {
	if binaryTree == nil {
		return 0;
	}
	leftLayers := binaryTree.Left.LayersRecursion()
	rightLayers := binaryTree.Right.LayersRecursion()

	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

//非递归-层数-借助队列，在进行按层遍历时，记录遍历的层数即可
func (binaryTree *BinaryTree) Layers() int {
	if binaryTree == nil {
		return 0
	}
	layers := 0
	nodes := []*BinaryTree{binaryTree}

	for len(nodes) > 0 {
		layers++
		size := len(nodes)
		count := 0
		for count < size {
			count++
			curNode := nodes[0]
			nodes = nodes[1:]
			if curNode.Left != nil {
				nodes = append(nodes, curNode.Left)
			}
			if curNode.Right != nil {
				nodes = append(nodes, curNode.Right)
			}
		}
	}
	return layers
}
