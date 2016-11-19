package gosort

func mergeSort(A []int) []int {
	l := len(A)
	if l <= 1 {
		return A
	}
	var la int = l/2 + l%2
	left := mergeSort(A[0:la])
	right := mergeSort(A[la:])
	return merge(left, right)
}

func merge(A, B []int) []int {
	la, lb := len(A), len(B)
	C := make([]int, la+lb)
	i, j := 0, 0
	for i < la && j < lb {
		if A[i] < B[j] {
			C[i+j] = A[i]
			i++
		} else {
			C[i+j] = B[j]
			j++
		}
	}
	for k := i; k < la; k++ {
		C[k+lb] = A[k]
	}
	for k := j; k < lb; k++ {
		C[k+la] = B[k]
	}
	return C
}
