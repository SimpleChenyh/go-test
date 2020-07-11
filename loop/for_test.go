package loop

import (
	"testing"
	"time"
)

func TestForLoop(t *testing.T) {

	for true {
		t.Log("biu biu")
	}

}


func TestForLoop2(t *testing.T) {

	n := 0

	for n < 10 {
		t.Log("biu biu")
		time.Sleep(1 * time.Second)
		n++
	}

}


func TestForLoop3(t *testing.T) {

	n := 10

	for i:=0;i<n;n++ {
		t.Log("biu biu")
		time.Sleep(1 * time.Second)
		n++
	}

}