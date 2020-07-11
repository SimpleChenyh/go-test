package syncf

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

var urls = []string{
	"https://www.google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

var wait = &sync.WaitGroup{}

var block = &sync.Mutex{}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url,"-->",resp.Status)
	time.Sleep(3 * time.Second)
	wait.Done()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	block.Lock()
	defer block.Unlock()
	fmt.Println("HomePage Endpoint Hit")
	for _, url := range urls {
		wait.Add(1)
		go fetch(url)
		//fetch(url)
	}
	wait.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "All Responses Received")
}

func handleRequests() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8081", nil)
}

func TestReq(t *testing.T) {
	//client := http.DefaultClient
	//
	//client.Timeout = 5 * time.Second
recover()

	handleRequests()
}

