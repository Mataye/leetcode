package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

func maxArea(height []int) int {
	res ,left ,right := -1 ,0,len(height)-1
	for left < right  {
		res  = max(min(height[left],height[right]) * (right-left),res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(a1,a2 int) int  {
	if a1 > a2 {
		return a2
	} else {
		return  a1
	}
}

func max(a1,a2 int) int  {
	if a1 > a2 {
		return a1
	} else {
		return  a2
	}
}