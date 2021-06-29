// 一个网站域名，如"discuss.leetcode.com"，包含了多个子域名。
// 作为顶级域名，常用的有"com"，下一级则有"leetcode.com"，最低的一级为"discuss.leetcode.com"。
// 当我们访问域名"discuss.leetcode.com"时，也同时访问了其父域名"leetcode.com"以及顶级域名 "com"。
//
// 给定一个带访问次数和域名的组合，要求分别计算每个域名被访问的次数。其格式为访问次数+空格+地址，例如："9001 discuss.leetcode.com"。
// 接下来会给出一组访问次数和域名组合的列表cpdomains 。要求解析出所有域名的访问次数，输出格式和输入格式相同，不限定先后顺序。
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subdomain-visit-count
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
// 注意事项：
//  - cpdomains 的长度小于 100。
//  - 每个域名的长度小于100。
//  - 每个域名地址包含一个或两个"."符号。
//  - 输入中任意一个域名的访问次数都小于10000。
package main

import "strconv"

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		parse(cpdomain, m)
	}
	result := make([]string, len(m))
	i := 0
	for domain, count := range m {
		result[i] = strconv.Itoa(count) + " " + domain
		i++
	}
	return result
}

func parse(line string, m map[string]int) {
	count := 0
	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			count, _ = strconv.Atoi(line[:i])
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] == ' ' {
			m[line[i+1:]] += count
			break
		} else if line[i] == '.' {
			m[line[i+1:]] += count
		}
	}
}
