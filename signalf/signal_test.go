package signalf

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignal(t *testing.T) {

	sigs := make(chan os.Signal,1)
	done := make(chan bool, 1) // Done channel


	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM,syscall.SIGSTOP)

	go func() {
		sig := <-sigs
		t.Log("Receiver signal ", sig)
		done <- true
	}()

	t.Log("Waiting for signal")
	<- done
	t.Log("ALl done")



}

