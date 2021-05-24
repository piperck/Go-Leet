//给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，
//为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。
//示例 1：
//
//输入：intervals = [[0,30],[5,10],[15,20]]
//输出：2
//示例 2：
//
//输入：intervals = [[7,10],[2,4]]
//输出：1

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Init MyHeap
type IntHeap []int

// implement Heap interface
func (p IntHeap) Len() int           { return len(p) }
func (p IntHeap) Less(i, j int) bool { return p[i] < p[j] }
func (p IntHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *IntHeap) Push(x interface{}) {
	*p = append(*p, x.(int))
}
func (p *IntHeap) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[0:p.Len()-1]
	return v
}


func minMeetingRooms(intervals [][]int) int {
	//return 0
	//heap.Interface()
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	myHeap := IntHeap{}
	maxRoom := 0

	for _, v := range intervals {
		if myHeap.Len() == 0 {
			heap.Push(&myHeap, v[1])
			maxRoom++
			continue
		}

		if v[0] < myHeap[0] {
			heap.Push(&myHeap, v[1])
		} else {
			heap.Push(&myHeap, v[1])
			heap.Pop(&myHeap)
		}
		if maxRoom < myHeap.Len() {
			maxRoom = myHeap.Len()
		}
	}

	return maxRoom
}

func main() {
	intervals := [][]int{
		{7,10},{2,4},
	}
	x := minMeetingRooms(intervals)
	fmt.Println(x)
}
