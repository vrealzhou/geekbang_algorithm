package data

type Heap struct {
	tail    int
	data    []interface{}
	compare func(t, o interface{}) int
}

// Variable size heap if size is 0
func NewHeap(size int, compare func(t, o interface{}) int) *Heap {
	return &Heap{
		tail:    -1,
		data:    make([]interface{}, size),
		compare: compare,
	}
}

func (h *Heap) Add(v interface{}) {
	h.tail++
	index := h.tail
	if len(h.data) == index {
		h.data = append(h.data, v)
	} else {
		h.data[index] = v
	}
	for index > 0 && h.compare(h.data[(index-1)/2], v) < 0 {
		newIndex := (index - 1) / 2
		h.swap(index, newIndex)
		index = newIndex
	}
}

func (h *Heap) Top() interface{} {
	if h.tail == -1 {
		return nil
	}
	return h.data[0]
}

func (h *Heap) Pop() interface{} {
	if h.tail == -1 {
		return nil
	}
	result := h.data[0]
	h.data[0] = h.data[h.tail]
	h.tail--
	index := 0
	left := index*2 + 1
	right := index*2 + 2
	for left <= h.tail {
		compareLeft := h.compare(h.data[index], h.data[left])
		if right > h.tail && compareLeft < 0 {
			h.swap(index, left)
			index = left
			break
		}
		if compareLeft > 0 && h.compare(h.data[index], h.data[right]) > 0 {
			break
		}
		if h.compare(h.data[left], h.data[right]) > 0 {
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

func (h *Heap) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}
