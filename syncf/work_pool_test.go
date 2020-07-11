package syncf

import (
	"log"
	"testing"
)

func work(jobs <-chan int){

	for job := range jobs {
		log.Print("Job ",job)
	}

}


func TestWorker(t *testing.T) {

	jobs := make(chan int)

	for i := 0; i < 3; i++ {
		go work(jobs)
	}


	for i := 0; i < 3; i++ {
		jobs <- i
	}



}
