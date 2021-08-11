/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。

两个字符串完全匹配才算匹配成功。

说明:

    s 可能为空，且只包含从 a-z 的小写字母。
    p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([]bool, 2*(m+1))
	index := func(i, j int) int {
		return (i%2)*(m+1) + j
	}
	dp[index(0, 0)] = true
	for i := 1; i <= m && p[i-1] == '*'; i++ {
		dp[index(0, i)] = true
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '?' {
				// fmt.Println(i, j, string(s[i-1]), dp[index(i-1, j-1)])
				dp[index(i, j)] = dp[index(i-1, j-1)]
			} else if p[j-1] == '*' {
				dp[index(i, j)] = dp[index(i, j-1)] || // match empty
					dp[index(i-1, j-1)] || // match 1
					dp[index(i-1, j)] // match more
				// fmt.Println(i, j, string(s[i-1]), dp[index(i, j)])
			} else {
				dp[index(i, j)] = dp[index(i-1, j-1)] && s[i-1] == p[j-1]
			}
			fmt.Println(i, j, string(s[i-1]), dp[index(i-1, 0)], dp[index(i-1, j-1)], dp[index(i, j)])
		}
		if i == 1 { // clear initial value on 0,0
			dp[index(0, 0)] = false
		}
	}
	return dp[index(n, m)]
}

func main() {
	fmt.Println(isMatch("", "*****"))
}
