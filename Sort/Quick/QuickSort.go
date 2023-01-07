package main

import "fmt"

func quicksort(a []int, l, r int) {
	if l < r {
		q := partition(a, l, r)
		quicksort(a, l, q)
		quicksort(a, q+1, r)
	}
}

func partition(a []int, l, r int) int {
	v := a[(l+r)/2]
	i, j := l, r
	for i <= j {
		for a[i] < v {
			i++
		}
		for a[j] > v {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return j
}

func main() {

	arax := []int{4, 234, 90, 2341, 90, 23434, 66, 91, 0, 10000000}
	fmt.Println(len(arax))
	quicksort(arax, 0, len(arax)-1)
	fmt.Println(arax)

}
