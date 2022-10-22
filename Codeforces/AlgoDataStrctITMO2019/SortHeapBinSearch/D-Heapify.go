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
	n    int
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

func (b *BinaryHeap) siftUp(i int) {
	for i > 0 && b.heap[i] >= b.heap[(i-1)/2] {
		b.heap[i], b.heap[(i-1)/2] = b.heap[(i-1)/2], b.heap[i]
		i = (i - 1) / 2
	}
}

func (b *BinaryHeap) siftDown(i int) {
	for 2*i+1 < b.n {
		j := 2*i + 1
		if 2*i+2 < b.n && b.heap[2*i+2] >= b.heap[j] {
			b.heap[i], b.heap[2*i+2] = b.heap[2*i+2], b.heap[i]
			i = 2*i + 2
		}
		if b.heap[i] > b.heap[j] {
			break
		} else {
			b.heap[i], b.heap[j] = b.heap[j], b.heap[i]
			i = 2*i + 1
		}
	}
}

func (b *BinaryHeap) insert(x int) {
	b.heap = append(b.heap, x)
	b.siftUp(b.n)
	b.n++
}

func (b *BinaryHeap) extract() int { // if n = 1
	res := b.heap[0]
	b.heap[0], b.heap[b.n-1] = b.heap[b.n-1], b.heap[0]
	b.n--
	b.heap = b.heap[0:b.n]
	b.siftDown(0)
	return res

}

func main() {
	in := bufio.NewReader(os.Stdin)
	n := readInt(in)
	heap := BinaryHeap{
		heap: make([]int, 0),
		n:    0,
	}

	res := make([]int, 0)
	for i := 0; i < n; i++ {
		arr := readArrInt(in)
		if arr[0] == 0 {
			heap.insert(arr[1])
			//fmt.Println("Heap ", heap.heap)

		} else {
			if heap.n != 0 {
				res = append(res, heap.extract())

			}
			//fmt.Println("Heap ", heap.heap)
		}
	}
	for _, v := range res {
		fmt.Printf("%d \n", v)
	}
}
