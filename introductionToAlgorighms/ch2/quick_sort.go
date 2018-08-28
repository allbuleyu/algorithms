package ch2

func partition(a []int, p, r int) int {
	x := a[r-1]

	i := p-1

	for j := p; j < r-1; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[r-1] = a[r-1], a[i+1]

	return i+1
}

func QuickSort(a []int, p, r int)  {
	if p < r {
		q := partition(a, p, r)

		QuickSort(a, p, q)
		QuickSort(a, q+1, r)
	}

}