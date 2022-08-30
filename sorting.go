package main

// Could be used to find "K" largest elements
func BubbleSort(arr []int) {
	n := len(arr)

	// For each pass
	for i := 0; i < n-1; i++ {
		// Assume array is sorted at start
		sorted := true

		// For each element during pass
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// Swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}

func InsertionSort(arr []int) {
	n := len(arr)

	// We must start from the second index position since we compare the element
	// behind current.
	for i := 1; i < n; i++ {
		j := i - 1
		x := arr[i]

		for j > -1 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}

		// We stop when we've found the element less than what is being inserted
		// therefore we insert the elemnt at j+1
		arr[j+1] = x
	}
}

func SelectionSort(arr []int) {
	n := len(arr)

	// n-1 because the last element is sorted by default on 2nd to last sort
	for i := 0; i < n-1; i++ {
		j := i
		k := i

		for j < n {
			if arr[j] < arr[k] {
				k = j
			}

			// Advance j to check next element as lowest
			j++
		}

		// Perform the swap that sorts the ith element
		arr[i], arr[k] = arr[k], arr[i]
	}
}

// https://www.geeksforgeeks.org/quick-sort/
func Partition(arr []int, low, high int) int {
	// The element to be placed at right position
	pivot := arr[high]

	// Index of smaller element and indicates the right position of pivot found
	// so far
	i := low - 1

	for j := low; j <= high-1; j++ {
		// If current element is smaller than pivot
		if arr[j] < pivot {
			// Increment index of smaller element
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		// pi is partitioning index, arr[pi] is now at right place
		pi := Partition(arr, low, high)

		// Separately sort elements before partition and after partition
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func merge2(arr []int, low, mid, high int) {
	n := len(arr)
	res := make([]int, n)

	i := low     // start of first list
	j := mid + 1 // start of second list
	k := low     // index for smaller value in result
	for i <= mid && j <= high {
		if arr[i] < arr[j] {
			res[k] = arr[i]
			i++
		} else {
			res[k] = arr[j]
			j++
		}

		k++
	}

	for i <= mid {
		res[k] = arr[i]
		i++
		k++
	}

	for j <= high {
		res[k] = arr[j]
		j++
		k++
	}

	// Copy sorted elements back into source
	for i := low; i <= high; i++ {
		arr[i] = res[i]
	}
}

func MergeSort(arr []int) {
	n := len(arr)
	var pass int

	// O(n log n) passes
	for pass = 2; pass <= n; pass *= 2 {
		// Number of iterations needed to merege per pass
		for i := 0; i+pass-1 < n; i += pass {
			low := i
			high := i + pass - 1
			mid := (low + high) / 2
			merge2(arr, low, mid, high)
		}
	}

	if pass/2 < n {
		merge2(arr, 0, pass/2-1, n-1)
	}
}

func MergeSortR(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2

		// Sort the left half
		MergeSortR(arr, low, mid)

		// Sort the right half
		MergeSortR(arr, mid+1, high)

		// Sort
		merge2(arr, low, mid, high)
	}
}

func findMax(arr []int) int {
	max := arr[0]
	for _, n := range arr[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

func CountSort(arr []int) {
	n := len(arr)
	max := findMax(arr)
	temp := make([]int, max+1)

	for i := 0; i < n; i++ {
		temp[arr[i]]++
	}

	i := 0
	j := 0
	for i < max+1 {
		if temp[i] > 0 {
			arr[j] = i
			j++
			temp[i]--
		} else {
			i++
		}
	}
}
