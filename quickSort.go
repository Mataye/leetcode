package main

import "fmt"

func main() {
	nums := []int{4, 7, 3, 1, 5, 8, 9, 0}
	quickSort(nums, 0,len(nums)-1)
	fmt.Println(nums)

}

func quickSort(nums []int,_left , _right int)  {
	left := _left
	right := _right
	temp := 0
	if left<right  {
		temp = nums[left]
		for left != right  {
			//从后向前找出第一个小于temp的数字
			for nums[right] >= temp && right >left {
				right--
			}
			//找到第一个小于temp的数字之后,放到temp出
			nums[left] = nums[right]

			//从前到后找打第一个大于temp的
			for nums[left] <= temp && left < right {
				left++
			}
			//找到之后放到后面
			nums[right] = nums[left]
		}
		nums[left] = temp
		quickSort(nums,left,_left)
		quickSort(nums,left+1,_right)
	}
	return
}


