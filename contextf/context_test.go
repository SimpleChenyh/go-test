package contextf

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wait := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wait.Add(1)
		go task(ctx, wait, i)
	}
	time.Sleep(10 * time.Second)
	cancel()

	wait.Wait()
}

func task(ctx context.Context, group *sync.WaitGroup, taskId int) {
	for {
		select {
		case <-ctx.Done():
			log.Print("Finish !!! task:", taskId)
			group.Done()
			return
		default:
			time.Sleep(1 * time.Second)
			log.Print("Continue task:", taskId)
		}
	}
}
