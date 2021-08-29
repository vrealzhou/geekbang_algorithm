package data

type Heap struct {
	tail    int
	size    int
	index   int
	data    []item
	compare func(t, o interface{}) int
}

type item struct {
	value interface{}
	index int
}

// Variable size heap if size is 0
func NewHeap(size int, compare func(t, o interface{}) int) *Heap {
	return &Heap{
		tail:    -1,
		size:    size,
		data:    make([]item, size),
		compare: compare,
	}
}

func (h *Heap) Add(v interface{}) {
	h.tail++
	index := h.tail
	if len(h.data) == index {
		h.data = append(h.data, item{
			value: v,
			index: h.index,
		})
	} else {
		h.data[index] = item{
			value: v,
			index: h.index,
		}
	}
	h.index++
	for index > 0 && h.compare(h.data[(index-1)/2].value, v) < 0 {
		newIndex := (index - 1) / 2
		h.swap(index, newIndex)
		index = newIndex
	}
}

func (h *Heap) Top() interface{} {
	if h.tail == -1 {
		return nil
	}
	h.removeOutdated()
	return h.data[0].value
}

func (h *Heap) removeOutdated() {
	for h.data[0].index < h.index-h.size {
		h.pop()
	}
}

func (h *Heap) pop() interface{} {
	if h.tail == -1 {
		return nil
	}
	result := h.data[0].value
	h.data[0] = h.data[h.tail]
	h.tail--
	index := 0
	left := index*2 + 1
	right := index*2 + 2
	for left <= h.tail {
		compareLeft := h.compare(h.data[index].value, h.data[left].value)
		if right > h.tail && compareLeft < 0 {
			h.swap(index, left)
			index = left
			break
		}
		if compareLeft > 0 && h.compare(h.data[index].value, h.data[right].value) > 0 {
			break
		}
		if h.compare(h.data[left].value, h.data[right].value) > 0 {
			h.swap(index, left)
			index = left
		} else {
			h.swap(index, right)
			index = right
		}
		left = index*2 + 1
		right = index*2 + 2
	}
	return result
}

func (h *Heap) Pop() interface{} {
	h.removeOutdated()
	return h.pop()
}

func (h *Heap) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}
