package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestFileCopy(t *testing.T) {
	// Prep a space for file copying.
	sandbox := NewSandbox(t)
	defer sandbox.Close()

	// Get our file paths
	message := []byte("Hello there!")
	sourcePath := sandbox.Create("source.txt", message)
	destPath := sandbox.Path("destination.txt")

	// Perform a copy
	err := fileCopy(sourcePath, destPath, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	messageCopy, err := ioutil.ReadFile(destPath)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if string(messageCopy) != string(message) {
		fmt.Println("File contents not preserved...")
		t.Fail()
	}

	// Perform another copy. Fail to overwrite.
	err = fileCopy(sourcePath, destPath, false)
	if err == nil {
		fmt.Println("Overwrote existing file.")
		t.Fail()
	}

	// Perform a copy. Overwrite file.
	message = []byte("Different message!")
	sourcePath = sandbox.Create("source2.txt", message)
	err = fileCopy(sourcePath, destPath, true)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
