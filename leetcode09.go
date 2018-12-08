package main

import (
	"strconv"
)

func main()  {

}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	runes := []rune(str)
	strLen := len(runes)

	var dp [][]int
	for i := 0; i < strLen; i++ {
		temp := make([]int, strLen)
		dp = append(dp, temp)
	}
	for i := 0; i < strLen; i++ {
		dp[i][i] = 1
	}

	max, left, right := -1, 0, 0

	for i := strLen - 1; i >= 0; i-- {
		for j := i + 1; j < strLen; j++ {
			if runes[i] == runes[j] {
				if i == j-1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if max < j-i+1 && dp[i][j] == 1 {
				max, left, right = j-i+1, i, j
			}
		}
	}
	if left == 0 && right == strLen - 1 {
		return true
	}
	return false

}
