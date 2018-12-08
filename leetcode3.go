//Time : 2018/12/7 15:47 
//Author : yehang
//File : leetcode3
//Software: GoLand
package main

import "fmt"

func main() {
 fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	runeStr := []rune(s)
	strLen := len(runeStr)
	exitMap := make(map[rune]int)
	startFalg := 0
	endFlag := 0
	max :=1
	for i:=0 ;i<strLen ;i++ {
		temp := runeStr[i]
		pre , isOk := exitMap[temp]
		if !isOk {
			exitMap[temp] = i
			endFlag++
		} else  {
			exitMap[temp] = i
			if pre < startFalg {
				endFlag++
			} else {
				startFalg = pre+1
				endFlag++
			}
		}
		if max < (endFlag- startFalg)  {
			max  = endFlag- startFalg
		}


	}
	return max
}