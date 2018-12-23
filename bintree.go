package main

import (
	"fmt"
)

type BSTNode struct {
	Value int
	Left *BSTNode
	Right *BSTNode
	Parent *BSTNode
}

//创建节点
func CreateNode(value int,parent ,left ,right *BSTNode) *BSTNode  {
	node := new(BSTNode)
	node.Value = value
	node.Parent = parent
	node.Left = left
	node.Right = right
	return node
}

func FindNode(root *BSTNode,key int) *BSTNode  {
	if root == nil {
		return  nil
	}
	if key == root.Value  {
		return root
	}
	if key > root.Value {
		return FindNode(root.Right,key)
	} else {
		return FindNode(root.Left,key)
	}
}

func DeleteNode(root ,delNode *BSTNode) {
	//删除节点没有子节点
	if delNode.Left == nil && delNode.Right == nil {
		if delNode.Parent == nil {
			root = nil
			return
		} else {
			parent := delNode.Parent
			if parent.Left == delNode {
				parent.Left = nil
			} else if parent.Right == delNode {
				parent.Right = nil
			}
		}
	}
	//有两个节点
	//展出右节点中最小的那个
	if delNode.Left != nil && delNode.Right != nil {
	}


		//删除节点有一个子节点
	if delNode.Left != nil || delNode.Right != nil {
		if delNode.Parent == nil {
			if delNode.Left != nil {
				root = delNode.Left
				return
			}
			if delNode.Right != nil {
				root = delNode.Right
				return
			}
		} else {
			parent := delNode.Parent
			if parent.Left == delNode {
				if delNode.Left != nil {
					root = delNode.Left
					return
				}
				if delNode.Right != nil {
					root = delNode.Right
					return
				}
			} else if parent.Right == delNode {
				if delNode.Left != nil {
					root = delNode.Left
					return
				}
				if delNode.Right != nil {
					root = delNode.Right
					return
				}
			}
		}
	}










	}


func InsertNode(root,node *BSTNode) *BSTNode {
	tempRoot := root
	var parentNode *BSTNode
	for tempRoot != nil  {
		parentNode = tempRoot
		if tempRoot.Value > node.Value {
			tempRoot = tempRoot.Left
		} else if tempRoot.Value < node.Value {
			tempRoot = tempRoot.Right
		} else {
			return root
		}
	}
	//初始化的时候
	if parentNode == nil {
		root = node
		return root
	}
	node.Parent = parentNode
	if node.Value > parentNode.Value {
		parentNode.Right = node
	} else {
		parentNode.Left = node
	}
	return root
}
func InOrderTree(root *BSTNode)  {
	if root != nil {
		InOrderTree(root.Left)
		fmt.Println(root.Value)
		InOrderTree(root.Right)
	}
}
func PreOrderTree(root *BSTNode)  {
	if root != nil {
		fmt.Println(root.Value)
		PreOrderTree(root.Left)
		PreOrderTree(root.Right)
	}
}
func BackOrderTree(root *BSTNode)  {
	if root != nil {
		BackOrderTree(root.Left)
		BackOrderTree(root.Right)
		fmt.Println(root.Value)
	}
}

func main()  {
	nums :=[]int{1,5,4,3,2,6}
	var root *BSTNode
	for i := 0; i < len(nums);i++ {
		root = InsertNode(root,CreateNode(nums[i],nil,nil,nil))
	}
	tempRoot := root
	myNode := FindNode(tempRoot,11)
	fmt.Println(myNode)
	//PreOrderTree(tempRoot)
	//fmt.Println("*********************")
	//tempRoot = root
	//InOrderTree(tempRoot)
	//fmt.Println("*********************")
	//tempRoot = root
	//BackOrderTree(tempRoot)
	//fmt.Println("*********************")


}

