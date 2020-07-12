package httpFileServerf

import (
	"log"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	server := http.FileServer(http.Dir("/windowsShare"))
	log.Fatal(http.ListenAndServe(":8080", server))

}
