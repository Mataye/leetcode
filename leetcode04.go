//Time : 2018/12/7 16:10 
//Author : yehang
//File : leetcode04
//Software: GoLand
package main

import (
	"fmt"
)

func main() {
	a := []int{1,3}
	b := []int{2}
	fmt.Println(findMedianSortedArrays(a,b))
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total %2 != 0 {
		return  float64(getMiddleware(nums1,len(nums1),nums2,len(nums2),total/2+1))
	} else {
		a := getMiddleware(nums1,len(nums1),nums2,len(nums2),total/2)
		b := getMiddleware(nums1,len(nums1),nums2,len(nums2),total/2+1)
		fmt.Println(a,b)
		return  float64(a+b)/2
	}
}

func getMiddleware(nums1 []int,len1 int, nums2 []int, len2,k int) int {
	if len1 > len2 {
		return getMiddleware(nums2,len2 ,nums1,len1, k)
	}
	if len1 == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums2[0],nums1[0])
	}

	temp1 :=min(k/2,len1)
	temp2 := k - temp1
	if nums1[temp1-1] > nums2[temp2-1] {
		return getMiddleware(nums1,len1, nums2[temp2:],len2-temp2, k-temp2)
	} else if nums1[temp1-1] < nums2[temp2-1] {
		return getMiddleware(nums1[temp1:],len1-temp1, nums2,len2, k-temp1)
	} else {
		return nums2[temp2-1]
	}
}


func min(a1,a2 int) int  {
	if a1 > a2 {
		return a2
	} else {
		return  a1
	}
}