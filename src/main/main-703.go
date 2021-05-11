//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
//请实现 KthLargest 类：
//
//KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
//int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

///**
// * Your KthLargest object will be instantiated and called as such:
// * obj := Constructor(k, nums);
// * param_1 := obj.Add(val);
// */
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]

package main

import (
	"fmt"
	"math"
)

type KthLargest struct {
	arr []int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	kth.arr = make([]int, 0, k)

	for _, v := range nums {
		if kth.isFull() {
			if kth.peek() > v {
				continue
			}
			kth.remove()
			kth.insert(v)
			continue
		}
		kth.insert(v)
	}
	return kth
}

func (this *KthLargest) isFull() bool {
	return len(this.arr) == cap(this.arr)
}

func (this *KthLargest) Add(val int) int {
	if this.isFull() {
		if this.peek() > val {
			return this.peek()
		}
		this.remove()
		this.insert(val)
	} else {
		this.insert(val)
	}
	return this.peek()
}

func (this *KthLargest) insert(val int) {
	this.arr = append((*this).arr, val)
	this.shiftUp(len(this.arr) - 1)
}

func (this *KthLargest) remove() {
	if len(this.arr) < 1 {
		return
	}
	this.arr[0], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[0]
	this.arr = this.arr[:len(this.arr)-1]
	this.shiftDown(0)
}

func (this *KthLargest) peek() int {
	return (*this).arr[0]
}

func (this *KthLargest) shiftUp(last int) {
	parent := int(math.Floor(float64((last - 1) / 2)))
	if parent < 0 {
		return
	}

	if this.arr[parent] > this.arr[last] {
		this.arr[parent], this.arr[last] = this.arr[last], this.arr[parent]
		last = parent
	} else {
		return
	}

	this.shiftUp(last)
	return
}

func (this *KthLargest) shiftDown(parent int) {
	left := 2 * parent + 1
	right := 2 * parent + 2

	overLeft := left > len(this.arr) - 1
	overRight := right > len(this.arr) - 1

	if overLeft {
		return
	}

	if !overLeft && overRight {
		if this.arr[parent] > this.arr[left] {
			this.arr[parent], this.arr[left] = this.arr[left], this.arr[parent]
		}
		return
	}

	if !overLeft && !overRight {
		if this.arr[parent] > this.arr[left] {
			if this.arr[left] < this.arr[right] {
				this.arr[parent], this.arr[left] = this.arr[left], this.arr[parent]
				parent = left
			} else {
				this.arr[parent], this.arr[right] = this.arr[right], this.arr[parent]
				parent = right
			}
		} else if this.arr[parent] > this.arr[right] {
			this.arr[parent], this.arr[right] = this.arr[right], this.arr[parent]
			parent = right
		} else {
			return
		}
	}

	this.shiftDown(parent)
	return
}


func main() {
	k := 2
	nums := []int{0}
	obj := Constructor(k, nums)
	param_3 := obj.Add(-1)
	//obj.Add(1)
	//obj.Add(-2)
	//obj.Add(-4)
	//param_3 := obj.Add(3)

	fmt.Println(param_3)
}
