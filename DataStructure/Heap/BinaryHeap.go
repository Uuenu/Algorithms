package main

import "fmt"

type BinaryHeap struct {
	heap []int
	n    int // heap size
}

func (a *BinaryHeap) siftUp(i int) {
	for i > 0 && a.heap[i] <= a.heap[(i-1)/2] {
		a.heap[i], a.heap[(i-1)/2] = a.heap[(i-1)/2], a.heap[i]
		i = (i - 1) / 2
	}
}

func (a *BinaryHeap) siftDown(i int) {
	for 2*i+1 < a.n {
		j := 2*i + 1
		if 2*i+2 < a.n && a.heap[2*i+2] < a.heap[j] {
			a.heap[i], a.heap[2*i+2] = a.heap[2*i+2], a.heap[i]
			i = 2*i + 2
		}
		if a.heap[i] < a.heap[j] {
			break
		} else {
			a.heap[i], a.heap[j] = a.heap[j], a.heap[i]
			i = j
		}
	}
}

func (a *BinaryHeap) min() int {
	return a.heap[0]
}

func (a *BinaryHeap) insert(x int) {
	a.heap = append(a.heap, x)
	a.n++
	a.siftUp(len(a.heap) - 1)
}

func (a *BinaryHeap) removeMin() {
	n := len(a.heap) - 1
	a.n--
	a.heap[0], a.heap[n] = a.heap[n], a.heap[0]
	a.heap = a.heap[0:n]
	a.siftDown(0)
}

func main() {
	arr := BinaryHeap{
		make([]int, 0),
		0,
	}
	arr.insert(5)
	arr.insert(3434)
	arr.insert(43234)
	arr.insert(90)
	arr.insert(71)
	arr.insert(874)
	arr.insert(1)
	fmt.Println(arr)
	arr.removeMin()
	fmt.Println(arr)
}
