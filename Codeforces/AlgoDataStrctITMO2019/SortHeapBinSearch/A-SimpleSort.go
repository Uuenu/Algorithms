package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BinaryHeap struct {
	heap []int
	n    int // heap size
}

func readInt(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func readLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func readArrInt(in *bufio.Reader) []int {
	numbs := readLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
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
	a.n--
	result := a.min()
	a.heap[0], a.heap[a.n] = a.heap[a.n], a.heap[0]
	a.heap = a.heap[0:a.n]
	a.siftDown(0)
	return result

}

func merge(a []int, left, mid, right int) {
	it1, it2 := 0, 0
	result := make([]int, right-left)

	for left+it1 < mid && mid+it2 < right {
		if a[left+it1] < a[mid+it2] {
			result[it1+it2] = a[left+it1]
			it1++
		} else {
			result[it1+it2] = a[mid+it2]
			it2++
		}
	}

	for left+it1 < mid {
		result[it1+it2] = a[left+it1]
		it1++
	}
	for mid+it2 < right {
		result[it1+it2] = a[mid+it2]
		it2++
	}

	for i := 0; i < it1+it2; i++ {
		a[left+i] = result[i]
	}

}

func mergeSortRecrusive(a []int, left, right int) {
	if left+1 >= right {
		return
	}
	mid := (left + right) / 2
	mergeSortRecrusive(a, left, mid)
	mergeSortRecrusive(a, mid, right)
	merge(a, left, mid, right)

}

func main() {

	in := bufio.NewReader(os.Stdin)
	n := readInt(in)
	arr := readArrInt(in)
	mergeSortRecrusive(arr, 0, len(arr))
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

}
