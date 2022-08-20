package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestDisplay(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(test)
}

func TestAppend(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	test = append(test, 10)
	fmt.Println(test)
}

func vanillaInsert(nums []int, idx, val int) []int {
	if idx < 0 || idx > len(nums) {
		return nums
	}

	// Save current length count
	len := len(nums)

	// Add an extra slot to our array so we can safely shift elements
	nums = append(nums, 0)

	// Shift elements up until we reach the desired index
	for i := len; i > idx; i-- {
		nums[i] = nums[i-1]
	}

	// Overwrite previous value with desired insert value
	nums[idx] = val

	return nums
}

func TestVanillaInsert(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	test = vanillaInsert(test, 5, 10)
	fmt.Println(test)

	test = vanillaInsert(test, 0, -1)
	fmt.Println(test)

	test = vanillaInsert(test, 4, 99)
	fmt.Println(test)
}

func goLazyInsert(slice []int, idx, val int) []int {
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice
}

func TestGoLazyInsert(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	test = goLazyInsert(test, 5, 10)
	fmt.Println(test)

	test = goLazyInsert(test, 0, -1)
	fmt.Println(test)

	test = goLazyInsert(test, 4, 99)
	fmt.Println(test)
}

func vanillaDelete(slice []int, idx int) []int {
	if idx < 0 || idx > len(slice) {
		return slice
	}

	// Overwrite the delted index with the next element.
	// Continue shifting all elements to right by overwriting with the next
	// element until reaching n-1
	for i := idx; i < len(slice)-1; i++ {
		// Copy the next element to current position
		slice[i] = slice[i+1]
	}

	// We want to return an updated slice that contains the udpated length
	return slice[:len(slice)-1]
}

func TestVanillaDelete(t *testing.T) {
	test := []int{8, 3, 7, 15, 6, 9, 10}
	test = vanillaDelete(test, 3)
	fmt.Printf("len=%d for %v\n", len(test), test)
}

func goDelete(slice []int, idx int) []int {
	// Shift elements up to perform actual delete
	copy(slice[idx:], slice[idx+1:])

	// Set last copied, initialized, item to zero value for gc
	slice[len(slice)-1] = 0

	// Return new slice, excluding zeroed value and updated length
	return slice[:len(slice)-1]
}

func TestGoDelete(t *testing.T) {
	test := []int{8, 3, 7, 15, 6, 9, 10}
	test = goDelete(test, 3)
	fmt.Printf("len=%d for %v\n", len(test), test)
}

func linearSearch(slice []int, key int) int {
	for i := 0; i < len(slice); i++ {
		if key == slice[i] {
			return i
		}
	}

	return -1
}

func TestLinearSearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(linearSearch(test, 4))  // 2
	fmt.Println(linearSearch(test, 6))  // 4
	fmt.Println(linearSearch(test, 15)) // -1
}

func swap(x *int, y *int) {
	// Store the current value at x
	temp := *x

	// Modify the value at x to now be value from y
	*x = *y

	// Modify the value at y to now be value from x
	*y = temp
}

func TestSwap(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(test)
	swap(&test[1], &test[4])
	fmt.Println(test) // 2 6 4 5 3
}

func linearSearchTransposition(slice []int, key int) int {
	for i := 0; i < len(slice); i++ {
		if key == slice[i] {
			if i == 0 {
				return 0
			}

			swap(&slice[i], &slice[i-1])
			return i - 1
		}
	}

	return -1
}

func TestLinearSearchTransposition(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 4), test)  // 1
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 6), test)  // 3
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 15), test) // -1
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 2), test)  // 0
}

func linearSearchHead(slice []int, key int) int {
	for i := 0; i < len(slice); i++ {
		if key == slice[i] {
			swap(&slice[i], &slice[0])
			return 0
		}
	}

	return -1
}

func TestLinearSearchHead(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 4), test)  // 0
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 6), test)  // 0
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 15), test) // -1
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 2), test)  // 0
}

func binarySearch(slice []int, key int) int {
	low := 0
	high := len(slice) - 1

	for low <= high {
		// Calculate current mid point
		mid := (low + high) / 2

		if key == slice[mid] {
			return mid
		} else if key < slice[mid] {
			// We need to search in the lower half, reduce high
			high = mid - 1
		} else {
			// We need to search in the higher half, increase low
			low = mid + 1
		}
	}

	// Key does not exist in array
	return -1
}

func TestBinarySearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(binarySearch(test, 5))  // 3
	fmt.Println(binarySearch(test, 2))  // 0
	fmt.Println(binarySearch(test, 10)) // -1
}

func rBinarySearch(slice []int, low, high, key int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2

	if key == slice[mid] {
		return mid
	} else if key < slice[mid] {
		return rBinarySearch(slice, low, mid-1, key)
	} else {
		return rBinarySearch(slice, mid+1, high, key)
	}
}

func TestRBinarySearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 5))  // 3
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 2))  // 0
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 10)) // -1
}

func reverse(A []int) {
	for L, R := 0, len(A)-1; L < R; L, R = L+1, R-1 {
		// Swap
		A[L], A[R] = A[R], A[L]
		// OR vanilla swap
		//temp := A[L]
		//A[L] = A[R]
		//A[R] = temp
	}
}

func TestReverse(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	reverse(test)
	fmt.Println(test) // 6 5 4 3 2
}

func getRotatedIndex(nums []int, idx, k int, isLeft bool) int {
	n := len(nums)
	offset := k % n

	if isLeft {
		idx -= offset
		return (idx + n) % n
	}

	return (idx + offset) % n
}

func rotate(nums []int, k int, isLeft bool) {
	start := 0
	count := 0

	for count < len(nums) {
		currentIdx := start
		previousVal := nums[start]

		for {
			nextIdx := getRotatedIndex(nums, currentIdx, k, isLeft)

			// swap
			temp := nums[nextIdx]       // save the current value before overwritten
			nums[nextIdx] = previousVal // overwrite with rotated value
			previousVal = temp

			// Update index position
			currentIdx = nextIdx

			count++

			if start == currentIdx {
				break
			}
		}

		start++
	}
}

func TestRotate(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	rotate(test, 13, true)
	fmt.Println(test) // 5 6 2 3 4

	test = []int{2, 3, 4, 5, 6}
	rotate(test, 13, false)
	fmt.Println(test) // 4 5 6 2 3
}

func rotate2(nums []int, k int, isLeft bool) {
	count := 0
	currentIdx := 0
	previousVal := nums[0]

	for count < len(nums) {
		nextIdx := getRotatedIndex(nums, currentIdx, k, isLeft)

		// Swap
		temp := nums[nextIdx]
		nums[nextIdx] = previousVal
		previousVal = temp

		// Update index position for next swap
		currentIdx = nextIdx

		count++
	}
}

func TestRotate2(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	rotate2(test, 13, true)
	fmt.Println(test) // 5 6 2 3 4

	test = []int{2, 3, 4, 5, 6}
	rotate2(test, 13, false)
	fmt.Println(test) // 4 5 6 2 3
}

func insertSorted(arr []int, x int) []int {
	// Get the last index position
	i := len(arr) - 1

	// Append to Go slice so we know that an element will fit
	arr = append(arr, 0)

	// While valid index and value is greater than insert value
	for i >= 0 && arr[i] > x {
		// Copy element over to next value
		arr[i+1] = arr[i]
		i--
	}

	// Inser the value in the correct position
	arr[i+1] = x

	return arr
}
func TestInsertSorted(t *testing.T) {
	test := []int{2, 3, 5, 12, 15}
	test1 := insertSorted(test, 1)
	fmt.Println(test1) // 1 2 3 5 12 15

	test2 := insertSorted(test, 20)
	fmt.Println(test2) // 2 3 5 12 15 20

	test3 := insertSorted(test, 12)
	fmt.Println(test3) // 2 3 5 12 12 15
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func TestIsSorted(t *testing.T) {
	fmt.Println(isSorted([]int{2, 3, 5, 12, 15}))
	fmt.Println(isSorted([]int{2, 3, 55, 12, 15}))
}

// make necessary swaps so negative values on one side and positive on other
func rearrange(arr []int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		// Iterate from front until encountering a positive value
		for arr[i] < 0 {
			i++
		}

		// Iterate from back until encountering a negative value
		for arr[j] >= 0 {
			j--
		}

		// If i position has not passed j, swap
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
}

func TestRearrange(t *testing.T) {
	test := []int{2, -3, 25, 10, 29, -15, -7}
	rearrange(test)
	fmt.Println(test) // -7 -3 -15 10 29 25 2
}

// arrays must be sorted
func merge(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	res := make([]int, arr1Len+arr2Len)

	arr1Idx := 0
	arr2Idx := 0
	idx := 0
	for arr1Idx < arr1Len && arr2Idx < arr2Len {
		if arr1[arr1Idx] < arr2[arr2Idx] {
			res[idx] = arr1[arr1Idx]
			arr1Idx++
		} else {
			res[idx] = arr2[arr2Idx]
			arr2Idx++
		}

		// Either way we're advancing the result index
		idx++
	}

	// Copy any trailing elements
	for arr1Idx < arr1Len {
		res[idx] = arr1[arr1Idx]
		arr1Idx++
		idx++
	}

	for arr2Idx < arr2Len {
		res[idx] = arr2[arr2Idx]
		arr2Idx++
		idx++
	}

	return res
}

func TestMerge(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 4, 7, 18, 20}
	fmt.Println(merge(test1, test2)) // 2 3 4 6 7 10 15 18 20 25
}

// arrays must be sorted
func union(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	res := make([]int, 0, arr1Len+arr2Len)

	arr1Idx := 0
	arr2Idx := 0
	idx := 0
	for arr1Idx < arr1Len && arr2Idx < arr2Len {
		if arr1[arr1Idx] < arr2[arr2Idx] {
			res = append(res, arr1[arr1Idx])
			arr1Idx++
		} else if arr2[arr2Idx] < arr1[arr1Idx] {
			res = append(res, arr2[arr2Idx])
			arr2Idx++
		} else { // we have equal values
			res = append(res, arr1[arr1Idx])
			arr1Idx++
			arr2Idx++
		}

		// Either way we're advancing the result index
		idx++
	}

	// Copy any trailing elements
	for arr1Idx < arr1Len {
		res = append(res, arr1[arr1Idx])
		arr1Idx++
		idx++
	}

	for arr2Idx < arr2Len {
		res = append(res, arr2[arr2Idx])
		arr2Idx++
		idx++
	}

	return res
}

func TestUnion(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(union(test1, test2)) // 2 3 6 7 10 15 20 25
}

// arrays must be sorted
func intersection(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	res := make([]int, 0, arr1Len+arr2Len)

	arr1Idx := 0
	arr2Idx := 0
	idx := 0
	for arr1Idx < arr1Len && arr2Idx < arr2Len {
		if arr1[arr1Idx] < arr2[arr2Idx] {
			arr1Idx++
		} else if arr2[arr2Idx] < arr1[arr1Idx] {
			arr2Idx++
		} else { // we have equal values
			res = append(res, arr1[arr1Idx])
			arr1Idx++
			arr2Idx++
		}

		// Either way we're advancing the result index
		idx++
	}

	return res
}

func TestIntersection(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(intersection(test1, test2)) // 6 15
}

// arrays must be sorted
func difference(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	res := make([]int, 0, arr1Len+arr2Len)

	arr1Idx := 0
	arr2Idx := 0
	idx := 0
	for arr1Idx < arr1Len && arr2Idx < arr2Len {
		if arr1[arr1Idx] < arr2[arr2Idx] {
			res = append(res, arr1[arr1Idx])
			arr1Idx++
		} else if arr2[arr2Idx] < arr1[arr1Idx] {
			arr2Idx++
		} else { // we have equal values
			arr1Idx++
			arr2Idx++
		}

		// Either way we're advancing the result index
		idx++
	}

	// Copy any trailing elements
	for arr1Idx < arr1Len {
		res = append(res, arr1[arr1Idx])
		arr1Idx++
		idx++
	}

	return res
}

func TestDifference(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(difference(test1, test2)) // 2 10 25
}

func contains(nums []int, key int) bool {
	return linearSearch(nums, key) != -1
}

func unionUnsorted(arr1, arr2 []int) []int {
	// Create map of arr1 values
	unionMap := make(map[int]bool)
	for _, num := range arr1 {
		unionMap[num] = true
	}

	for _, num := range arr2 {
		unionMap[num] = true
	}

	res := make([]int, len(unionMap))
	i := 0
	for k := range unionMap {
		res[i] = k
		i++
	}

	return res
}

func TestUnionUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(unionUnsorted(test1, test2)) // 2 6 10 15 25 3 7 20 (or some other order)
}

func intersectionUnsorted(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)

	res := make([]int, 0, arr1Len+arr2Len)

	// Create map of arr2 values
	a2Map := make(map[int]bool)
	for _, num := range arr2 {
		a2Map[num] = true
	}

	for _, num := range arr1 {
		if _, ok := a2Map[num]; ok {
			res = append(res, num)
		}
	}

	return res
}

func TestIntersectionUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(intersectionUnsorted(test1, test2)) // 6 15
}

func differenceUnsorted(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	res := make([]int, 0, arr1Len)

	// Create map of arr2 values
	a2Map := make(map[int]bool)
	for _, num := range arr2 {
		a2Map[num] = true
	}

	for _, num := range arr1 {
		if _, ok := a2Map[num]; !ok {
			res = append(res, num)
		}
	}

	return res
}

func TestDifferenceUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(differenceUnsorted(test1, test2)) // 2 10 25
}

// Expectes nums to be 1...N with a single missing element
func findMissingElement1(nums []int) int {
	n := len(nums)
	sum := 0
	max := nums[0]

	// Get the sum of nums and max value
	for i := 0; i < n; i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	// Caclulated the expected consecutive sum
	expectedSum := max * (max + 1) / 2

	return expectedSum - sum
}

func TestFindMissingElement1(t *testing.T) {
	fmt.Println(findMissingElement1([]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12})) // 7
	fmt.Println(findMissingElement1([]int{12, 8, 3, 9, 5, 6, 11, 1, 10, 4, 2})) // 7
}

func findMissingElement2(nums []int) (int, error) {
	n := len(nums)

	// Sort the array
	sort.Ints(nums)

	// Expected difference from val-idx
	diff := nums[0]

	for i := 0; i < n; i++ {
		if nums[i]-i != diff {
			// Found missing!
			return i + diff, nil
		}
	}

	return 0, fmt.Errorf("no missing element")
}

func TestFindMissingElement2(t *testing.T) {
	test := []int{6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17}
	miss, err := findMissingElement2(test)
	if err != nil {
		t.Fatalf("expected no error, got: %s", err)
	}

	fmt.Println(miss) // 12

	test = []int{17, 6, 16, 7, 15, 8, 14, 9, 13, 10, 11}
	miss, err = findMissingElement2(test)
	if err != nil {
		t.Fatalf("expected no error, got: %s", err)
	}

	fmt.Println(miss) // 12
}

func findMissingElements(nums []int) []int {
	n := len(nums)
	res := make([]int, 0, n)

	// Sort the array
	sort.Ints(nums)

	// Expected different
	diff := nums[0]

	for i := 0; i < n; i++ {
		currentDiff := nums[i] - i

		if currentDiff != diff {
			// Handles cases where multiple values missing between indexes
			for diff < currentDiff {
				missing := i + diff
				res = append(res, missing)

				diff++
			}
		}
	}

	return res
}

func TestFindMissingElements(t *testing.T) {
	fmt.Println(findMissingElements([]int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19})) // 10 13 14
	fmt.Println(findMissingElements([]int{19, 6, 18, 7, 17, 8, 16, 9, 15, 11, 12})) // 10 13 14
}

func findMissingElements2(nums []int) []int {
	n := len(nums)
	missing := make([]int, 0, n)
	min := nums[0]
	max := nums[0]

	// Set min and max
	for i := 0; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	// Create hash table for missing lookup
	hashTable := make([]int, max+1)

	for i := 0; i < n; i++ {
		// Use the value from nums as the hash index
		hash := nums[i]

		// Update the bit/count in the hash table to show this number exists
		hashTable[hash]++
	}

	// Iterate the hash table to find missing numbers
	for i := min; i < max; i++ {
		if hashTable[i] == 0 {
			missing = append(missing, i)
		}
	}

	return missing
}

func TestFindMissingElements2(t *testing.T) {
	fmt.Println(findMissingElements2([]int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19})) // 10 13 14
	fmt.Println(findMissingElements2([]int{19, 6, 18, 7, 17, 8, 16, 9, 15, 11, 12})) // 10 13 14
}

func duplicatesInSorted(nums []int) []int {
	lastDuplicate := 0
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] && nums[i] != lastDuplicate {
			res = append(res, nums[i])
			lastDuplicate = nums[i]
		}
	}

	return res
}

func TestDuplicatesInSorted(t *testing.T) {
	fmt.Println(duplicatesInSorted([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // 8 15
}

func duplicatesAndCountInSorted(nums []int) [][]int {
	var res [][]int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			j := i + 1

			for nums[j] == nums[i] {
				j++
			}

			// Get count by taking difference from j and i
			count := j - i

			// Add dup and count to output
			res = append(res, []int{nums[i], count})

			// Update i to last duplciated element seen
			i = j - 1
		}
	}

	return res
}

func TestDuplicatesAndCountInSorted(t *testing.T) {
	fmt.Println(duplicatesAndCountInSorted([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // [8, 2] [15, 3]
}

func duplicatseAndCountWithHash(nums []int) [][]int {
	var res [][]int
	hashTable := make(map[int]int)

	// Populate hashtable with elements and counts
	for _, num := range nums {
		hashTable[num]++
	}

	// Iterate hash and pick out the duplicate instances
	for num, count := range hashTable {
		if count > 1 {
			res = append(res, []int{num, count})
		}
	}

	return res
}

func TestDuplicatesAndCountWithHash(t *testing.T) {
	fmt.Println(duplicatseAndCountWithHash([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // [8, 2] [15, 3]
}

// Returns the index positions of the two values that sum to K
func twoSum(nums []int, k int) [][]int {
	var res [][]int
	hashTable := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := k - nums[i]

		if idx, ok := hashTable[complement]; ok {
			// We have seen a num in the list that matches the needed complement
			res = append(res, []int{i, idx})
		} else {
			// Add the current num and its index to our hash table for later ref
			hashTable[nums[i]] = i
		}
	}

	return res
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{6, 3, 8, 10, 16, 7, 5, 2, 9, 14}, 10)) // [5 1] [7 2]
}
