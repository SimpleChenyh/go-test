package condition

import (
	"testing"
)

func TestIf(t *testing.T) {
	t.Log("")

	if a := pass(false); a {

		t.Log("Pass")

	} else {

		t.Log("Not Pass")
	}

}

func TestIf1(t *testing.T) {
	t.Log("")

	var a bool

	if a = pass(true); a {
		t.Log("Pass")
	} else {
		t.Log("Not Pass")
	}

	t.Log(a)

}

func pass(pass bool) bool {
	return pass
}

