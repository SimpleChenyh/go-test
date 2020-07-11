package syncf

import (
	"sync"
	"testing"
	"time"
)


func TestSync1(t *testing.T) {
	lock := &sync.Mutex{}

	counter := 0


	for n := 0; n < 20000; n++ {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter++

		}()
	}
	time.Sleep(time.Millisecond * 10)
	t.Log(counter)
}




func TestSyncWait1(t *testing.T) {
	lock := &sync.Mutex{}
	wait := &sync.WaitGroup{}

	counter := 0


	for n := 0; n < 20000; n++ {
		wait.Add(1)
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter++
			wait.Done()
		}()
	}
	wait.Wait()
	t.Log(counter)
}

