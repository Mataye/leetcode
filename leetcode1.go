//Time : 2018/12/7 15:39 
//Author : yehang
//File : leetcode1
//Software: GoLand
package main
import "fmt"
func main()  {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i,k :=range nums{
		numMap[k] = i
	}
	for i,k:= range nums{
		temp := target - k
		if ff,isOk := numMap[temp];isOk {
			if i != ff {
				return []int{i,ff}
			}
		}
	}
	return []int{}
}