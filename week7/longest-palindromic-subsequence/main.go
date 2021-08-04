/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"math"
)

func longestPalindromeSubseq(s string) int {
	// s[i]==s[j]: f(i, j) = f(i-1,j-1)+2
	// s[i]!=s[j]: f(j, j) = max(f(i+1,j), f(i, j-1))
	// 初始: f(i,j)=1
	l := []byte(s)
	n := len(l)
	index := func(i, j int) int {
		return j%2*n + i
	}
	dp := make([]int, n*2)
	for j, c := range l {
		dp[index(j, j)] = 1
		for i := j - 1; i >= 0; i-- {
			if l[i] == c {
				dp[index(i, j)] = dp[index(i+1, j-1)] + 2
			} else {
				dp[index(i, j)] = max(dp[index(i, j-1)], dp[index(i+1, j)])
			}
			fmt.Printf("s[%d]=%c, s[%d]=%c\n", i, l[i], j, l[j])
			printDP(dp, n)
		}
	}
	return dp[index(0, n-1)]
}

func printDP(dp []int, n int) {
	for i := 0; i < 2; i++ {
		fmt.Println(dp[i*n : i*n+n])
	}
	fmt.Println("***************")
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
