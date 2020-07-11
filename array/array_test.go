package array

import (
	"log"
	"testing"
)

func TestArrayCompare(t *testing.T) {

	a := [...]int{1,2,3,4,5}
	b := [...]int{1,2,3,4,5}
	//c := [...]int{1,2,3,4,5,6}


	log.Fatal( a== b)
	//log.Fatal( a == c)
	//log.Fatal( c == b)

}

const (
	Executable = 1 << iota
	Writable
	Readable
)

// Test &^ operator
func TestOperator( t *testing.T){
	a := 7

	a  = a &^ Writable

	log.Fatal(a &Executable == Executable,a &Writable == Writable,a &Readable == Readable)

}


