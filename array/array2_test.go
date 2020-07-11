package array

import (
	"testing"
)

func TestArray1(t *testing.T) {
	array1 := [...]int{1, 2, 2, 31, 23, 12, 312, 3, 123, 12, 31}
	array2 := [2]int{1, 2}
	array3 := [2]int{1, 2}

	array4 := [...][12]int{{12, 3, 124, 5, 43, 52, 12, 31, 23, 12, 41, 4},
		{12, 3, 124, 5, 43, 52, 12, 31, 23, 12, 41, 4}}

	t.Log(array1)
	t.Log(array2)
	t.Log(array2 == array3)

	t.Log(array4)

}

func TestSplit(t *testing.T) {
	a4 := [...][12]int{{12, 3, 124, 5, 43, 52, 12, 31, 23, 12, 41, 4},
		{13, 3, 124, 5, 43, 52, 12, 31, 23, 12}}

	t.Log(a4[1:])

	t.Log(a4[0][3:])

	for idx, v := range a4 {
		t.Log("1 st:",idx, v)

		for idx1, v1 := range v {
			t.Log("2 nd",idx1, v1)
		}

	}

}


// Slice share memory
func TestSplit1(t *testing.T) {

	a1 := []int{1,2,3,4,5}

	a2 := a1[1:]

	a2[0] = 100

	t.Log(a1)
	t.Log(a2)

}

// The new slice share the capacity with origin slice
// The new slice start to end
func TestSplitShareCapacity(t *testing.T) {

	a1 := []int{1,2,3,4,5}

	a2 := a1[1:2]

	a2[0] = 100

	t.Log(a1)
	t.Log(a2)
	t.Log(cap(a2))

}