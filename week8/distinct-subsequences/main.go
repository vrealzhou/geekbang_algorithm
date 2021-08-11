/*
给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

题目数据保证答案符合 32 位带符号整数范围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distinct-subsequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)
	dp := make([]int, 2*(m+1))
	index := func(i, j int) int {
		return (i%2)*(m+1) + j
	}
	for i := 0; i < 2; i++ {
		dp[index(i, 0)] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[index(i, j)] = dp[index(i-1, j)]
			if s[i-1] == t[j-1] {
				dp[index(i, j)] += dp[index(i-1, j-1)]
			}
		}
	}
	return dp[index(n, m)]
}
