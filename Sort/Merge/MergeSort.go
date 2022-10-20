package main

import "fmt"

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

	arrax := []int{5, 234, 24, 666, 231, 67, 985}
	mergeSortRecrusive(arrax, 0, len(arrax))
	fmt.Println(arrax)
}
