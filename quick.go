package gosort

func quickSort(A []int) []int {
	qs(A, 0, len(A)-1)
	return A
}

func qs(A []int, lo, hi int) {
	if lo < hi {
		p := partition(A, lo, hi)
		qs(A, lo, p-1)
		qs(A, p+1, hi)
	}
}

func partition(A []int, lo, hi int) int {
	pivot := A[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if A[j] <= pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i], A[hi] = A[hi], A[i]
	return i
}
