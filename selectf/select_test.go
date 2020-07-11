package selectf

import (
	"testing"
	"time"
)

func TestSelect(t *testing.T) {

	var c1 = make(chan string)
	var c2 = make(chan string)

	go func() {
		for n := 0; n <= 5; n++  {
			select {
			case <-time.After(1 * time.Second):
				t.Log("Timeout")
			case msg := <-c1:
				t.Log("Hello c1", msg)
			case msg := <-c2:
				t.Log("Hello c2", msg)
				//default:
				//	panic("Haha panic because no msg receive")
			}
		}
	}()
	go func() {
		c1 <- "c111"
	}()
	go func() {
		c2 <- "c222"
	}()
	time.Sleep(20 * time.Second)
	t.Log("--------")
}
