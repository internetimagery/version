package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

// Non-version constants
const LATEST = -1
const UNVERSIONED = 0

// Format structure
type Formater interface {
	Parse(string) *Format
	Build(*Format) string
}

type Format struct {
	Name    string
	Version int
}

// Simple formatter.
// name_v02.ext
// name_latest.ext
type SimpleFormat struct {
	Format
}

var SIMPLEPARSER = regexp.MustCompile(`^(.+)(_v\d+|_latest)(.\w+)?$`)

func (self *SimpleFormat) Parse(name string) {
	parts := SIMPLEPARSER.FindStringSubmatch(name)
	if len(parts) > 2 {
		self.Name = parts[1] + parts[3]
		if parts[2] == "_latest" {
			self.Version = LATEST
		} else if ver, err := strconv.Atoi(parts[2][2:]); err == nil {
			self.Version = ver
		}
	} else {
		self.Name = name
	}
}

func (self *SimpleFormat) Build() string {
	version := ""
	ext := filepath.Ext(self.Name)
	base := self.Name[:len(self.Name)-len(ext)]
	switch self.Version {
	case LATEST:
		version = "_latest"
	case UNVERSIONED:
		version = ""
	default:
		version = fmt.Sprintf("_v%02d", self.Version)
	}
	return base + version + ext
}
