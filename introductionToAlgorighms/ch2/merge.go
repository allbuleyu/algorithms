package ch2


var arr []int = []int{2,4,5,6,1,2,3,6}

func merge(a []int, low, mind, high int)  {
	tmp := make([]int, high-low)
	for i := 0; i<high-low; i++ {
		tmp[i] = a[low+i]
	}

	left := tmp[:mind-low]
	right := tmp[mind-low:]

	var i, j, k int
	i = 0
	j =0

	for k = low; k < high; k++ {
		if i < (mind-low) && left[i] <= right[j] {
			a[k] = left[i]
			i++
		}else if j < high-mind {
			a[k] = right[j]
			j++
		}
	}

}

func MergeSort(a []int, low, high int)  {
	if low < high {
		mind := (high+low)/2

		MergeSort(a, low, mind)
		MergeSort(a, mind+1, high)

		merge(a, low, mind, high)
	}
}

func MergeSortFor(a []int) {


}
