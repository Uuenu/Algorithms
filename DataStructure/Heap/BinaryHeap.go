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

func (a *BinaryHeap) insert(x int) { // O(log n)
	a.heap = append(a.heap, x)
	a.n++
	a.siftUp(len(a.heap) - 1)
}

func (a *BinaryHeap) removeMin() int { // O(log n)
	result := a.min()
	a.heap[0], a.heap[a.n] = a.heap[a.n], a.heap[0] // ???
	a.n--
	a.heap = a.heap[0:a.n]
	a.siftDown(0)
	return result

}

func HeapSort(a []int) { //  O(nlog n) + O(n) - memory
	bh := BinaryHeap{
		make([]int, 0),
		0,
	}
	//bh.heap = a
	for i := 0; i < len(a); i++ {
		bh.insert(a[i])
	}
	for i := 0; i < len(a); i++ {
		a[i] = bh.removeMin()
	}

}

func main() {

	sliceX := []int{5, 3, 54353, 234, 54235432, 32452345, 324532, 435236242634523}
	HeapSort(sliceX)
	fmt.Println(sliceX)

}
