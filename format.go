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
type FormatterBase interface {
	Parse(string)
	Build() string
}

type FormatterGetter interface {
	GetName() string
	GetVersion() int
}

type FormatterSetter interface {
	SetName(string)
	SetVersion(int)
}

type FormatterAll interface {
	FormatterBase
	FormatterGetter
	FormatterSetter
}

// Simple formatter.
// name_v02.ext
// name_latest.ext
type SimpleFormat struct {
	name    string
	version int
}

var SIMPLEPARSER = regexp.MustCompile(`^(.+)(_v\d+|_latest)(.\w+)?$`)

func (self *SimpleFormat) Parse(name string) {
	parts := SIMPLEPARSER.FindStringSubmatch(name)
	if len(parts) > 2 {
		self.name = parts[1] + parts[3]
		if parts[2] == "_latest" {
			self.version = LATEST
		} else if ver, err := strconv.Atoi(parts[2][2:]); err == nil && ver > 0 {
			self.version = ver
		}
	} else {
		self.name = name
	}
}

func (self *SimpleFormat) Build() string {
	version := ""
	ext := filepath.Ext(self.name)
	base := self.name[:len(self.name)-len(ext)]
	switch self.version {
	case LATEST:
		version = "_latest"
	case UNVERSIONED:
		version = ""
	default:
		version = fmt.Sprintf("_v%02d", self.version)
	}
	return base + version + ext
}

func (self *SimpleFormat) GetName() string {
	return self.name
}

func (self *SimpleFormat) SetName(value string) {
	self.name = value
}

func (self *SimpleFormat) GetVersion() int {
	return self.version
}

func (self *SimpleFormat) SetVersion(value int) {
	self.version = value
}
