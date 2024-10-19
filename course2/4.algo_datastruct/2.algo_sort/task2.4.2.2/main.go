package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)

	return result
}

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]

			j--
		}

		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func GeneralSort(arr []int) {
	// выбор эффективного алгоритма сортировки в зависимости от размера входного массива

	n := len(arr)

	switch {
	case n <= 30:
		insertionSort(arr)
	case n <= 1000:
		quickSort(arr)
	default:
		sortedArr := mergeSort(arr)
		copy(arr, sortedArr)
	}
}

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original: ", data)

	sortedData := mergeSort(data)
	fmt.Println("Sorted by Merge Sort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(data)
	fmt.Println("Sorted by Insertion Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(data)
	fmt.Println("Sorted by Selection Sort: ", data)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	sortedData = quickSort(data)
	fmt.Println("Sorted by Quicksort: ", sortedData)

	data = []int{64, 34, 25, 12, 22, 11, 90}
	GeneralSort(data)
	fmt.Println("Sorted by GeneralSort: ", data)
}
