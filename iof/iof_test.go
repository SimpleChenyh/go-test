package iof

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

func TestTemp(t *testing.T) {

	file, err := ioutil.TempFile("", "t_fffff")

	if err != nil {
		panic(err)
	}

	log.Print("File name:", file.Name())

	//os.Remove(file.Name())
	dir, err := ioutil.TempDir("", "tmpgo")
	if err != nil {
		panic(err)
	}
	join := filepath.Join(dir, "join_file")

	t.Log("join file name:", join)

	err = ioutil.WriteFile(join, []byte("Go lang is a beautiful program language"), 0666)

	if err != nil {
		t.Log(err)
	}

}
