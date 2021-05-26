package sort

import "testing"

func TestMergeSort(t *testing.T) {
	testArray := []int{3, 4, 5, 7, 2, 7, 4, 9, 1, 32, -6}

	MergeSort(&testArray, 0, len(testArray)-1)

	sortedArray := []int{-6, 1, 2, 3, 4, 4, 5, 7, 7, 9, 32}

	for i := 0; i < len(testArray); i++ {
		if testArray[i] != sortedArray[i] {
			t.Fatal(testArray, " should be ", sortedArray)
		}
	}
}
