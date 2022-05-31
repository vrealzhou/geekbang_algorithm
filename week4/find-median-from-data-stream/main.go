/*
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-median-from-data-stream
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"github.com/vrealzhou/geekbang_algorithm/internal/data"
)

type MedianFinder struct {
	firstHalf *data.MaxHeap[int]
	endHalf   *data.MaxHeap[int]
	count     int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		firstHalf: data.NewMaxHeap(0, func(source, target int) int {
			return source - target
		}, 0),
		endHalf: data.NewMaxHeap(0, func(source, target int) int {
			return target - source
		}, 0),
		count: 0,
	}
}

func (f *MedianFinder) AddNum(num int) {
	f.count++
	f.endHalf.Add(num)
	if f.count%2 != 0 {
		r, _ := f.endHalf.Pop()
		f.firstHalf.Add(r)
	} else {
		t1, _ := f.endHalf.Top()
		t2, _ := f.firstHalf.Top()
		if t1 < t2 {
			f.endHalf.Pop()
			f.firstHalf.Pop()
			f.firstHalf.Add(t1)
			f.endHalf.Add(t2)
		}
	}
	// fmt.Printf("count %d, first %v, second %v\n", f.count, f.firstHalf.Top(), f.endHalf.Top())
}

func (f *MedianFinder) FindMedian() float64 {
	first, more := f.firstHalf.Top()
	if !more {
		return 0
	}
	if f.count%2 != 0 {
		return float64(first)
	}
	second, _ := f.endHalf.Top()
	return float64(first+second) / 2
}
