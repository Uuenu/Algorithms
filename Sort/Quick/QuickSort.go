package main

func quicksort(a []int, l, r int) {
	q := partition(a, l, r)
	quicksort(a, l, q)
	quicksort(a, q+1, r)
}

func partition(a []int, l, r int) int {
	v := (l + r) / 2

	return v
}

func swap(a, b *int) {
	c := &a
	a = b
	b = *c
}

func main() {

}
