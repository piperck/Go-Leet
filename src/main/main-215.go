package main

import (
	"fmt"
	"math"
)

type maxHeap []int

func (h *maxHeap) peek() int {
	return (*h)[0]
}

func (h *maxHeap) lenIndex() int {
	return len(*h) - 1
}

func (h *maxHeap) isFull() bool {
	return h.lenIndex() + 1 == cap(*h)
}

func (h *maxHeap) insert(v int) {
	*h = append(*h, v)
	h.shiftUp(h.lenIndex())
}

func (h *maxHeap) remove() {
	if h.lenIndex() < 1 {
		return
	}
	(*h)[0], (*h)[h.lenIndex()] = (*h)[h.lenIndex()], (*h)[0]
	*h = (*h)[:h.lenIndex()]
	h.shiftDown(0)
}

func (h *maxHeap) shiftUp(last int) {
	parent := int(math.Floor(float64((last - 1) / 2)))
	if parent < 0 {
		return
	}

	if (*h)[parent] < (*h)[last] {
		(*h)[parent], (*h)[last] = (*h)[last], (*h)[parent]
		last = parent
	} else {
		return
	}

	h.shiftUp(last)
	return
}

func (h *maxHeap) shiftDown(parent int) {
	left := 2 * parent + 1
	right := 2 * parent + 2

	oLeft := left > h.lenIndex()
	oRight := right > h.lenIndex()

	if oLeft {
		return
	}

	if !oLeft && oRight {
		if (*h)[left] > (*h)[parent] {
			(*h)[left], (*h)[parent] = (*h)[parent], (*h)[left]
		}
		return
	}

	if !oLeft && !oRight {
		if (*h)[parent] < (*h)[left] {
			if (*h)[left] > (*h)[right] {
				(*h)[parent], (*h)[left] = (*h)[left], (*h)[parent]
				parent = left
			} else {
				(*h)[parent], (*h)[right] = (*h)[right], (*h)[parent]
				parent = right
			}
		} else if (*h)[parent] < (*h)[right] {
			(*h)[parent], (*h)[right] = (*h)[right], (*h)[parent]
			parent = right
		} else {
			return
		}
	}

	h.shiftDown(parent)
	return
}

// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
func findKthLargest(nums []int, k int) int {
	maxHeap := make(maxHeap, 0, len(nums))

	for _, v := range nums {
		maxHeap.insert(v)
	}

	for i := 0; i < k-1; i++ {
		maxHeap.remove()
	}

	return maxHeap.peek()
}

func main() {
	nums := []int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(nums, 2))
}
