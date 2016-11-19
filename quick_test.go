package gosort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	A := []int{10, 3, 2, 5, 6, 2, 6, 7, 8, 9, 1, 2, 5}
	B := QuickSort(A)
	for i := 0; i < len(B)-1; i++ {
		if B[i] > B[i+1] {
			t.FailNow()
		}
	}
}
