package chanf

import (
	"fmt"
	"log"
	"testing"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestChan(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println("t 3",x, y, x+y)
}

func TestChanStr(t *testing.T) {

	c := make(chan string)

	send := func(s string) {
		c <- s
	}
	go send("1")
	go send("2")
	go send("3")
	go send("4")

	v1 := <-c
	v2 := <-c
	v3 := <-c
	v4 := <-c

	log.Print("t 2",v1, v2, v3, v4)

}

func TestChanT1(t *testing.T) {
	log.Print("t4")
	chs := make(chan string)

	chs <- "1"
	close(chs)
	chs <- "2"
	close(chs)
	chs <- "3"
	close(chs)



	log.Print("t4",<-chs)
	log.Print("t4",<-chs)
	log.Print("t4",<-chs)


}


func TestChanT2(t *testing.T) {

	ch := make(chan string, 2)
	ch <- "1"
	ch <- "2"
	ch <- "3"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}



func TestChanBuf(t *testing.T) {

	c := make(chan string,5)

	send := func(s string) {
		c <- s
	}
	go send("1")
	go send("2")
	go send("3")
	go send("4")

	v1 := <-c
	v2 := <-c
	v3 := <-c
	v4 := <-c

	log.Print("t 1",v1, v2, v3, v4)

}