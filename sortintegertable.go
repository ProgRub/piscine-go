package piscine

func SortIntegerTable(table []int) {
	mergeSort(table, len(table))
}

func mergeSort(v []int, m int) {
	if m < 2 {
		return
	}
	mid := m / 2
	left := make([]int, mid)
	right := make([]int, m-mid)
	for i := 0; i < mid; i++ {
		left[i] = v[i]
	}
	for j := mid; j < m; j++ {
		right[j-mid] = v[j]
	}
	mergeSort(left, mid)
	mergeSort(right, m-mid)
	merge(left, right, v, mid, m-mid, m)
}

func merge(left []int, right []int, v []int, n_left int, n_right int, m int) {
	i := 0
	j := 0
	k := 0
	for i < n_left && j < n_right {
		if left[i] < right[j] {
			v[k] = left[i]
			i++
		} else {
			v[k] = right[j]
			j++
		}
		k++
	}
	for i < n_left {
		v[k] = left[i]
		k++
		i++
	}
	for j < n_right {
		v[k] = right[j]
		j++
		k++
	}
}
