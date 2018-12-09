package main

import (
	"fmt"
)

func main() {
	//常规解法
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

//普通解法
//func longestCommonPrefix(strs []string) string {
//	arrLen := len(strs)
//	if arrLen == 0  {
//		return ""
//	}
//
//	commonStr := strs[0]
//
//	for k:=1 ;k<arrLen ;k++{
//		if strs[k] == "" {
//			return  ""
//		}
//		runes := []rune(strs[k])
//		strLen := len(runes)
//
//		cRunes := []rune(commonStr)
//		cLen := len(cRunes)
//		temp :=""
//		for i:=0 ;i<min(strLen,cLen) ;i++ {
//			if runes[i] == cRunes[i] {
//				temp +=string(runes[i])
//			} else  {
//				break
//			}
//		}
//		commonStr = temp
//	}
//	return  commonStr
//}
//func min(a1,a2 int) int  {
//	if a1 > a2 {
//		return a2
//	} else {
//		return  a1
//	}
//}


func longestCommonPrefix(strs []string) string {
	root := new(Node)
	for i:=0 ;i<len(strs) ;i++ {
		if strs[i] == "" {
			return ""
		}

		makeTree(root,strs[i])
	}

	printTree(root)
	return findCommonPorx(root,"","",len(strs))
}

type Node struct {
	isEnd bool
	value string
	count int
	son map[string]*Node
}
//字典树解法
func makeTree(root *Node ,str string)  {
	runes :=[]rune(str)
	strLen := len(runes)
	if strLen == 0 {
		return
	}
	if root.son == nil {
		root.son = make(map[string]*Node)
	}
	tempRoot := root

	for i:=0; i<strLen;i++ {
		val := string(runes[i])
		tempNode ,isOk := tempRoot.son[val]
		if isOk {
			tempNode.count++
			if i == strLen-1 {
				tempNode.isEnd = true
				return
			}
			tempRoot = tempNode
		} else {
			tempNode = new(Node)
			tempNode.value = string(runes[i])
			tempNode.count++
			tempNode.son = make(map[string]*Node)
			tempRoot.son[string(runes[i])] = tempNode
			tempRoot = tempNode
			if i == strLen-1 {
				tempNode.isEnd = true
				return
			}
		}
	}
	return
}
func findCommonPorx(root *Node,str string,pre string,total int) string {
	if len(root.son) != 1  {
		return ""
	} else  {
		for val,son :=range root.son {
			if son.count != total {
				return  ""
			}
			mid  := findCommonPorx(son, str,val,total)
			str = val+mid
		}
	}
	return str
}

func printTree(root *Node)  {
	//if root.son == nil {
	//	fmt.Println("**********************")
	//	return
	//}
	for _,val :=range root.son {
		fmt.Println(val.isEnd,val.count,val.value)
		printTree(val)
	}
	return
}
