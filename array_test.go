package main

import (
	"fmt"
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

func TestVanillaInsert(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	test = vanillaInsert(test, 5, 10)
	fmt.Println(test)

	test = vanillaInsert(test, 0, -1)
	fmt.Println(test)

	test = vanillaInsert(test, 4, 99)
	fmt.Println(test)
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

func TestLinearSearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(linearSearch(test, 4))  // 2
	fmt.Println(linearSearch(test, 6))  // 4
	fmt.Println(linearSearch(test, 15)) // -1
}

func TestSwap(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(test)
	swap(&test[1], &test[4])
	fmt.Println(test) // 2 6 4 5 3
}

func TestLinearSearchTransposition(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 4), test)  // 1
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 6), test)  // 3
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 15), test) // -1
	fmt.Printf("idx=%d of %v\n", linearSearchTransposition(test, 2), test)  // 0
}

func TestLinearSearchHead(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 4), test)  // 0
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 6), test)  // 0
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 15), test) // -1
	fmt.Printf("idx=%d of %v\n", linearSearchHead(test, 2), test)  // 0
}

func TestBinarySearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(binarySearch(test, 5))  // 3
	fmt.Println(binarySearch(test, 2))  // 0
	fmt.Println(binarySearch(test, 10)) // -1
}

func TestRBinarySearch(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 5))  // 3
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 2))  // 0
	fmt.Println(rBinarySearch(test, 0, len(test)-1, 10)) // -1
}

func TestReverse(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	reverse(test)
	fmt.Println(test) // 6 5 4 3 2
}

func TestRotate(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	rotate(test, 13, true)
	fmt.Println(test) // 5 6 2 3 4

	test = []int{2, 3, 4, 5, 6}
	rotate(test, 13, false)
	fmt.Println(test) // 4 5 6 2 3
}

func TestRotate2(t *testing.T) {
	test := []int{2, 3, 4, 5, 6}
	rotate2(test, 13, true)
	fmt.Println(test) // 5 6 2 3 4

	test = []int{2, 3, 4, 5, 6}
	rotate2(test, 13, false)
	fmt.Println(test) // 4 5 6 2 3
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

func TestRearrange(t *testing.T) {
	test := []int{2, -3, 25, 10, 29, -15, -7}
	rearrange(test)
	fmt.Println(test) // -7 -3 -15 10 29 25 2
}

func TestMerge(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 4, 7, 18, 20}
	fmt.Println(merge(test1, test2)) // 2 3 4 6 7 10 15 18 20 25
}

func TestUnion(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(union(test1, test2)) // 2 3 6 7 10 15 20 25
}

func TestIntersection(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(intersection(test1, test2)) // 6 15
}

func TestDifference(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(difference(test1, test2)) // 2 10 25
}

func TestUnionUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(unionUnsorted(test1, test2)) // 2 6 10 15 25 3 7 20 (or some other order)
}

func TestIntersectionUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(intersectionUnsorted(test1, test2)) // 6 15
}

func TestDifferenceUnsorted(t *testing.T) {
	test1 := []int{2, 6, 10, 15, 25}
	test2 := []int{3, 6, 7, 15, 20}
	fmt.Println(differenceUnsorted(test1, test2)) // 2 10 25
}

func TestFindMissingElement1(t *testing.T) {
	fmt.Println(findMissingElement1([]int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12})) // 7
	fmt.Println(findMissingElement1([]int{12, 8, 3, 9, 5, 6, 11, 1, 10, 4, 2})) // 7
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

func TestFindMissingElements(t *testing.T) {
	fmt.Println(findMissingElements([]int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19})) // 10 13 14
	fmt.Println(findMissingElements([]int{19, 6, 18, 7, 17, 8, 16, 9, 15, 11, 12})) // 10 13 14
}

func TestFindMissingElements2(t *testing.T) {
	fmt.Println(findMissingElements2([]int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19})) // 10 13 14
	fmt.Println(findMissingElements2([]int{19, 6, 18, 7, 17, 8, 16, 9, 15, 11, 12})) // 10 13 14
}

func TestDuplicatesInSorted(t *testing.T) {
	fmt.Println(duplicatesInSorted([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // 8 15
}

func TestDuplicatesAndCountInSorted(t *testing.T) {
	fmt.Println(duplicatesAndCountInSorted([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // [8, 2] [15, 3]
}

func TestDuplicatesAndCountWithHash(t *testing.T) {
	fmt.Println(duplicatseAndCountWithHash([]int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20})) // [8, 2] [15, 3]
}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{6, 3, 8, 10, 16, 7, 5, 2, 9, 14}, 10)) // [5 1] [7 2]
}

func TestTranspose(t *testing.T) {
	test := [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	display := func() {
		for row := 0; row < len(test); row++ {
			for col := 0; col < len(test[0]); col++ {
				fmt.Printf("%d ", test[row][col])
			}
			fmt.Printf("\n")
		}
	}

	fmt.Println("BEFORE:")
	display()
	test = transpose(test)
	fmt.Println("AFTER:")
	display()
}
