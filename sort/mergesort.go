package sort

func merge(array *[]int, left int, median int, right int) {

	dimLeft := median - left + 1

	dimRight := right - median

	arrayLeft := make([]int, dimLeft)
	arrayRight := make([]int, dimRight)

	// fill temp arrays
	for i := 0; i < dimLeft; i++ {
		arrayLeft[i] = (*array)[left+i]
	}
	for i := 0; i < dimRight; i++ {
		arrayRight[i] = (*array)[median+1+i]
	}

	// index position of the left temp array
	leftIndex := 0

	// index position of the right temp array
	rightIndex := 0

	// index position of the final array

	arrayIndex := left

	for leftIndex < dimLeft && rightIndex < dimRight {

		if arrayLeft[leftIndex] < arrayRight[rightIndex] {
			(*array)[arrayIndex] = arrayLeft[leftIndex]
			leftIndex++
		} else {
			(*array)[arrayIndex] = arrayRight[rightIndex]
			rightIndex++
		}
		arrayIndex++
	}
	// exaust remaining values on the left and right arrays
	for leftIndex < dimLeft {
		(*array)[arrayIndex] = arrayLeft[leftIndex]
		leftIndex++
		arrayIndex++
	}
	for rightIndex < dimRight {
		(*array)[arrayIndex] = arrayRight[rightIndex]
		rightIndex++
		arrayIndex++
	}

}

// MergeSort sort an array in O(n log n)
func MergeSort(array *[]int, left int, right int) {

	if left >= right {
		return
	}

	var median int = left + (right-left)/2

	MergeSort(array, left, median)
	MergeSort(array, median+1, right)
	merge(array, left, median, right)
}
