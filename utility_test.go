package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
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

func TestFileUnique(t *testing.T) {
	sandbox := NewSandbox(t)
	defer sandbox.Close()

	sourcePath := sandbox.Create("file.txt", []byte("Some data here"))

	// Get a few unique names
	name1 := sandbox.Create(filepath.Base(fileUnique(sourcePath)), []byte{})
	name2 := sandbox.Create(filepath.Base(fileUnique(sourcePath)), []byte{})
	name3 := sandbox.Create(filepath.Base(fileUnique(sourcePath)), []byte{})
	if name1 == name2 || name1 == name3 || name2 == name3 {
		fmt.Println("File names not unique")
		t.Fail()
	}
}

func TestFileCompare(t *testing.T) {
	sandbox := NewSandbox(t)
	defer sandbox.Close()

	data1 := []byte("Some data n' stuff")
	data2 := []byte("Some data is different")

	source1 := sandbox.Create("source1.txt", data1)
	source2 := sandbox.Create("source2.txt", data2)
	source3 := sandbox.Create("source3.txt", data1)

	// Two equal files
	compared, err := fileCompare(source1, source1, 4096)
	if err != nil {
		t.Fatal(err)
	}
	if !compared {
		fmt.Println("Failed to compare two of the same files")
		t.Fail()
	}

	// Two equal content files, different creation time
	compared, err = fileCompare(source1, source3, 4096)
	if err != nil {
		t.Fatal(err)
	}
	if !compared {
		fmt.Println("Failed to compare two of the same content")
		t.Fail()
	}

	// Two different content files
	compared, err = fileCompare(source2, source3, 4096)
	if err != nil {
		t.Fatal(err)
	}
	if compared {
		fmt.Println("Failed to compare two different content files")
		t.Fail()
	}

}
