package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Testing helper
// Usage:
// sandbox := NewSandbox(t)
// defer sandbox.Close()
type Sandbox struct {
	path string
	t    *testing.T
}

// Create new Sandbox instance.
func NewSandbox(t *testing.T) *Sandbox {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	return &Sandbox{path: tmpDir, t: t}
}

// Clean up files.
func (self *Sandbox) Close() {
	err := os.RemoveAll(self.path)
	if err != nil {
		self.t.Fatal(err)
	}
}

// Return path local to Sandbox.
func (self *Sandbox) Path(name string) string {
	return filepath.Join(self.path, name)
}

// Create a file at path, local to sandbox and add data.
func (self *Sandbox) Create(name string, data []byte) string {
	path := self.Path(name)
	err := ioutil.WriteFile(path, data, 700)
	if err != nil {
		self.t.Fatal(err)
	}
	return path
}

// Stat a file. Expecting the file to exist, so throwing fatal exceptions if not.
func (self *Sandbox) Stat(path string) os.FileInfo {
	stat, err := os.Stat(path)
	if err != nil {
		self.t.Fatal(err)
	}
	return stat
}
