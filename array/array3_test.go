package array

import "testing"

func TestArrayMake1(t *testing.T) {

	var a1 []int

	a1 = append(a1, 1, 2, 3, 4)

	t.Log(a1)

	t.Log(len(a1),cap(a1))


}

func TestArrayMake2(t *testing.T) {

    a1:=make([]int,1,100)
	t.Log(a1)


    a1 = append(a1,1,2,3,4)

	t.Log(len(a1),cap(a1))


	a1 = append(a1,1,2,3,47)

	t.Log(len(a1),cap(a1))


	t.Log(a1[len(a1) -1 ])

	// runtime error: index out of range [9] with length 9 [recovered]
	a1[len(a1)] = 123
	t.Log(len(a1),cap(a1))

}



