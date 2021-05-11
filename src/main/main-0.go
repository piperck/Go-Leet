package main

import (
	"fmt"
	"math"
)

type Heap []int

func (h *Heap) shiftUp(last, parent int) {
	heap := *h
	if parent < 0 || last <= 0 {
		return
	}

	if heap[parent] > heap[last] {
		heap[parent], heap[last] = heap[last], heap[parent]
	}
	h.shiftUp(parent, int(math.Floor(float64((parent-1)/2))))
}

func (h *Heap) shiftDown(parent int) {
	heap := *h
	lenIndex := len(heap) - 1
	left := 2 * parent + 1
	right := 2 * parent + 2

	if parent >= lenIndex {
		return
	}

	overL := left > lenIndex
	overR := right > lenIndex

	if overL {
		return
	}

	if !overL && overR {
		if heap[parent] > heap[left] {
			heap[parent], heap[left] = heap[left], heap[parent]
			parent = left
		} else {
			return
		}
	}

	if !overL && !overR {
		if heap[parent] > heap[left] && heap[left] <= heap[right] {
			heap[parent], heap[left] = heap[left], heap[parent]
			parent = left
		} else if heap[parent] > heap[right] && heap[right] <= heap[left] {
			heap[parent], heap[right] = heap[right], heap[parent]
			parent = right
		} else {
			return
		}
	}

	h.shiftDown(parent)
	return
}

func (h *Heap) peek() int {
	return (*h)[0]
}

func (h *Heap) size() int {
	return len(*h)
}

func (h *Heap) insert(value int) {
	*h = append(*h, value)
	lenIndex := len(*h) - 1
	h.shiftUp(lenIndex, int(math.Floor(float64((lenIndex - 1) / 2))))
}

func (h *Heap) remove() {
	heap := *h
	heap[0], heap[len(heap) - 1] = heap[len(heap) - 1], heap[0]
	*h = heap[:len(heap) - 1]
	h.shiftDown(0)
}

// The minimum heap
func main() {
	zz := Heap{}
	zz = []int{10, 14, 25, 33, 81}
	if len(zz) <= 0 {
		return
	}

	//zz.insert(13)
	zz.remove()
	fmt.Println(zz)
}
