package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
fmt.Println(threeSum([]int{-3,0,-5,2,-4,-2,1,-2,-1,3,1,3,1,3,-3,-5,1}))
}

type Node struct {
	dp [][]int
	exit []string
}
//Time Limit Error
func threeSum(nums []int) [][]int {
	var temp []int
	res := new(Node)
	if len(nums) == 0 {
		return res.dp
	}
	sort.Sort(sort.IntSlice(nums))
	findThreeSum(nums,0,temp,0,res)
	return res.dp
}

func findThreeSum(nums []int,start int,temp []int,sum int,res *Node) {
	if sum == 0 && len(temp) == 3  {
		sort.Sort(sort.IntSlice(temp))
		str := ""
		for i:=0; i<3  ;i++  {
			str += strconv.Itoa(temp[i])+"#"
		}
		for i:=0; i<len(res.exit)  ;i++ {
			if res.exit[i] == str {
				return
			}
		}
		res.exit = append(res.exit,str)
		res.dp = append(res.dp,temp)
		return
	}
	if len(temp) > 3 {
		return
	}
	for i := start ; i< len(nums);i++ {
		 findThreeSum(nums,i+1,append(temp,nums[i]),sum+nums[i],res)
		 findThreeSum(nums,i+1,temp,sum,res)
	}
	return
}

