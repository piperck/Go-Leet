// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
    occurrences := map[int]int{}
    for _, num := range nums {
        occurrences[num]++
    }
    h := &BHeap{}
    heap.Init(h)
    for key, value := range occurrences {
        heap.Push(h, [2]int{key, value})
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    ret := make([]int, k)
    for i := 0; i < k; i++ {
        ret[k - i - 1] = heap.Pop(h).([2]int)[0]
    }
    return ret
}


type BHeap [][2]int
func (h BHeap) Len() int           { return len(h) }
func (h BHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h BHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *BHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *BHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0: len(*h)-1]
	return x
}

func main() {
	list := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(topKFrequent(list, k))
}
