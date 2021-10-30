/*

 Sorts integers by using N goroutines ( creating N sub-arrays ) and merging them into a single sorted array.
 For assignment purposes test and function are in the same test file, instead of being 2 different ones,
 Run the test via:

   $ go test -v
   -------------------------------
		=== RUN   Test_FourRoutineSort
		--- PASS: Test_FourRoutineSort (0.00s)
		PASS
		ok  0.294s

*/

package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"sync"
	"testing"
)

// ASSIGNMENT

//NGoRoutineSort sorts input by using N goroutines
func NGoRoutineSort(input []int, goRoutines int) []int {
	chunkSize := len(input) / goRoutines     // each routine will work on `chunkSize` elements
	outputCh := make(chan []int, goRoutines) // create the channel where we gonna store N goroutine partial sorted arrays
	wg := &sync.WaitGroup{}                  // it's used to wait for each go routine to finish sorting its chunk

	for i := 0; i < goRoutines; i++ { // spawn N goRoutines
		wg.Add(1)
		if i == goRoutines-1 {
			//feed the last go-routine with all the remaining elements of input
			go makeSortedSlice(input, i*chunkSize, len(input), wg, outputCh)
		} else {
			go makeSortedSlice(input, i*chunkSize, (i+1)*chunkSize, wg, outputCh)
		}
	}
	wg.Wait()       // all go routine sorted their sub-arrays, we have results in outputCh
	close(outputCh) // outputCh closing as no goroutine are active

	sorted := make([]int, 0)
	for msg := range outputCh {
		sorted = merge(sorted, msg) //merge a subarray with another, as sub-arrays are sorted their relative order is not important
	}
	return sorted
}

//makeSortedSlice returns a sorted subarray
func makeSortedSlice(array []int, start int, end int, wg *sync.WaitGroup, c chan []int) {
	defer wg.Done()
	out := array[start:end]
	sort.Ints(out)
	c <- out
}

//merge returns a merged and sorted array, given a and b sorted subarrays
func merge(a, b []int) []int {
	out := make([]int, len(a)+len(b))
	var k, i, j int
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			out[k] = b[j]
			j++
		} else {
			out[k] = a[i]
			i++
		}
		k++
	}
	for i < len(a) {
		out[k] = a[i]
		i++
		k++
	}
	for j < len(b) {
		out[k] = b[j]
		j++
		k++
	}

	return out
}

// TESTS

const FourRoutines = 4

func Test_FourRoutineSort(t *testing.T) {
	testcases := []struct {
		name        string
		routines    int
		sliceToSort []int
	}{
		{
			name:        "9 elements",
			routines:    FourRoutines,
			sliceToSort: []int{1, 3, 90, 11, 23, 56, 45, 67, 70},
		},
		{
			name:        "9 elements with 2 go-routines",
			routines:    2,
			sliceToSort: []int{1, 3, 90, 11, 23, 56, 45, 67, 70},
		},
		{
			name:        "9 elements with 6 go-routines",
			routines:    6,
			sliceToSort: []int{1, 3, 90, 11, 23, 56, 45, 67, 70},
		},
		{
			name:        "4 elements",
			sliceToSort: []int{1, 3, 90, 11},
			routines:    FourRoutines,
		},
		{
			name:        "1 element",
			sliceToSort: []int{1},
			routines:    FourRoutines,
		},
		{
			name:        "repeated elements",
			sliceToSort: []int{90, 7, 10, 3, 12, 5, 7, 8, 45, 3},
			routines:    5,
		},
		{
			name:        "negative numbers",
			sliceToSort: []int{90, 7, -1, 3, 12, 5, 7, -23, 45, 3},
			routines:    FourRoutines,
		},
		{
			name:        "empty",
			routines:    FourRoutines,
			sliceToSort: []int{},
		},
		{
			name:        "nil",
			routines:    FourRoutines,
			sliceToSort: nil,
		},
	}

	for _, tc := range testcases {
		expected := make([]int, len(tc.sliceToSort))
		copy(expected, tc.sliceToSort)
		sort.Ints(expected)
		assert.Equal(t, expected, NGoRoutineSort(tc.sliceToSort, tc.routines), tc.name)
	}
}
