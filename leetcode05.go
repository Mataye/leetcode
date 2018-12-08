//Time : 2018/12/7 20:55 
//Author : yehang
//File : leetcode05
//Software: GoLand
package main

import "fmt"

func main()  {
	fmt.Println(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	strLen := len(runes)
	left, right := 0, 0
	max := -1
	var dp [][]int
	for i:=0; i<strLen; i++ {
		temp := make([]int,strLen)
		dp = append(dp,temp)
	}

	for i := strLen - 1; i >= 0; i-- {
		dp[i][i] = 1
	}
	for i := strLen - 1; i >= 0; i-- {
		for j := i + 1; j < strLen; j++ {
			if runes[i] == runes[j] {
				if i == j-1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 {
				if max < j-i+1 {
					max = j - i + 1
					left, right = i, j
				}
			}

		}
	}
	return s[left:right+1]
}