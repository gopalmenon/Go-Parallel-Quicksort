package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestParallelQuicksort(t *testing.T) {

	smallArray := []int{604, 940, 664, 437, 424, 686, 65, 156, 96, 300}
	smallSortedArray := []int{65, 96, 156, 300, 424, 437, 604, 664, 686, 940}
	Quicksort(smallArray, 0, len(smallArray) - 1)
	if !testEq(smallArray, smallSortedArray) {
		t.Errorf("Sorting failed for %v:", smallArray)
	}

	for _, length := range([]int {10, 100, 1000, 10000, 100000, 1000000}) {

		toSort := randomArray(length)
		start := time.Now()
		Quicksort(toSort, 0, length - 1)
		fmt.Printf("Sorted %d elements in %s\n", length, time.Since(start))

		toParallelSort := randomArray(length)
		start = time.Now()
		ParallelQuicksort(toParallelSort, 0, length - 1, nil)
		fmt.Printf("Parallel sorted %d elements in %s\n", length, time.Since(start))

	}

}

func randomArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = int(rand.Float64() * 1000)
	}
	return arr
}

func testEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}