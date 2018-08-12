package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
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

func TestFileLink(t *testing.T) {
	sandbox := NewSandbox(t)
	defer sandbox.Close()

	// Set up a file to link
	message := []byte("Testing 123")
	sourcePath := sandbox.Create("source.txt", message)
	destPath := sandbox.Path("destination.txt")

	// Link file!
	err := fileLink(sourcePath, destPath, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if !os.SameFile(sandbox.Stat(sourcePath), sandbox.Stat(destPath)) {
		fmt.Println("File not linked correctly")
		t.Fail()
	}

	// Link file. Fail on existing file.
	err = fileLink(sourcePath, destPath, false)
	if err == nil {
		fmt.Println("Existing file overridden...")
		t.Fail()
	}

	// Link file. Override result.
	message = []byte("New message")
	sourcePath = sandbox.Create("source2.txt", message)
	err = fileLink(sourcePath, destPath, true)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if !os.SameFile(sandbox.Stat(sourcePath), sandbox.Stat(destPath)) {
		fmt.Println("File2 not linked correctly")
		t.Fail()
	}
}
