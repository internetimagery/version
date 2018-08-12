package main

import (
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

// Link one file to another.
func fileLink(src, dest string, override bool) error {
	if _, err := os.Stat(dest); override && err == nil {
		err = os.Remove(dest)
		if err != nil {
			return err
		}
	}
	return os.Link(src, dest)
}
