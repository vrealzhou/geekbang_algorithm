/*
在本问题中, 树指的是一个连通且无环的无向图。

输入一个图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N) 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。

结果图是一个以边组成的二维数组。每一个边的元素是一对[u, v] ，满足 u < v，表示连接顶点u 和v的无向图的边。

返回一条可以删去的边，使得结果图是一个有着N个节点的树。如果有多个答案，则返回二维数组中最后出现的边。答案边 [u, v] 应满足相同的格式 u < v。

注意:

    输入的二维数组大小在 3 到 1000。
    二维数组中的整数在1到N之间，其中N是输入数组的大小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/redundant-connection
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func findRedundantConnection(edges [][]int) []int {
	ds := make([]int, len(edges)+1)
	var find func(i int) int
	find = func(i int) int {
		if ds[i] != i {
			ds[i] = find(ds[i])
		}
		return ds[i]
	}
	for _, edge := range edges {
		if ds[edge[0]] == 0 {
			ds[edge[0]] = edge[0]
		}
		if ds[edge[1]] == 0 {
			ds[edge[1]] = edge[1]
		}
		x, y := find(edge[0]), find(edge[1])
		if x == y {
			return edge
		}
		ds[y] = x // 链接两个树必须是根连根
		// fmt.Println(edge, x, y, ds)
	}
	return nil
}
