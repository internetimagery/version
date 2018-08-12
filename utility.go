package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Copy file from one location to another.
// Clean up if errors.
func fileCopy(src, dest string, override bool) error {
	// Check if destination already exists
	if _, err := os.Stat(dest); !override && !os.IsNotExist(err) {
		return fmt.Errorf("File already exists: %s", dest)
	}

	handleSrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer handleSrc.Close()

	handleDest, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		handleDest.Close()
		if err != nil {
			os.Remove(dest)
		}
	}()

	// Copy the data across
	if _, err = io.Copy(handleDest, handleSrc); err != nil {
		return err
	}
	return handleDest.Sync()
}

// Create a unique temp file name.
func fileUnique(path string) string {
	index := 0
	tmpPath := ""
	for {
		index++
		tmpPath = fmt.Sprintf("%s.%d.tmp", path, index)
		_, err := os.Stat(tmpPath)
		if os.IsNotExist(err) {
			break
		}
	}
	return tmpPath
}

// Compare two files to see if they're identical
// block: 4096
func fileCompare(source1, source2 string, block int) (bool, error) {
	// Check if sizes vary
	stat1, err := os.Stat(source1)
	if err != nil {
		return false, err
	}
	stat2, err := os.Stat(source2)
	if err != nil {
		return false, err
	}
	if stat1.Size() != stat2.Size() {
		return false, nil
	}

	// Compare contents
	handle1, err := os.Open(source1)
	if err != nil {
		return false, err
	}
	defer handle1.Close()
	handle2, err := os.Open(source2)
	if err != nil {
		return false, err
	}
	defer handle2.Close()

	data1 := make([]byte, block)
	data2 := make([]byte, block)

	for {
		_, err1 := handle1.Read(data1)
		_, err2 := handle2.Read(data2)
		if err1 == io.EOF || err2 == io.EOF {
			if err1 != err2 {
				return false, nil
			}
			break
		}
		if err1 != nil {
			return false, err1
		}
		if err2 != nil {
			return false, err2
		}
		if !bytes.Equal(data1, data2) {
			return false, nil
		}
	}

	// We've made it this far! Must be the same file!
	return true, nil
}
