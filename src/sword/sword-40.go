// 剑指 Offer 40. 最小的 K 个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
package main

import (
	"fmt"
	"math"
)

type maxHeap []int

func (h *maxHeap) insert(v int) {
	*h = append(*h, v)
	h.shiftUp(h.lenIndex())
}

func (h *maxHeap) peek() int {
	return (*h)[0]
}

func (h *maxHeap) lenIndex() int {
	return len(*h) - 1
}

func (h *maxHeap) isFull() bool {
	return h.lenIndex() + 1 == cap(*h)
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

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	mh := make(maxHeap, 0, k)
	for _, v := range arr {
		if mh.isFull() {
			if mh.peek() > v {
				mh.remove()
				mh.insert(v)
			}
			continue
		}
		mh.insert(v)
	}
	return mh
}


func main() {
	//[0,0,0,2,0,5]
	arr := []int{0,1,1,1,4,5,3,7,7,8,10,2,7,8,0,5,2,16,12,1,19,15,5,18,2,2,22,15,8,22,17,6,22,6,22,26,32,8,10,11,2,26,9,12,9,7,28,33,20,7,2,17,44,3,52,27,2,23,19,56,56,58,36,31,1,19,19,6,65,49,27,63,29,1,69,47,56,61,40,43,10,71,60,66,42,44,10,12,83,69,73,2,65,93,92,47,35,39,13,75}
	//arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	fmt.Println(getLeastNumbers(arr, 75))
}
