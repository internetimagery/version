package main

import (
	"fmt"
	"testing"
)

func TestSimpleFormat(t *testing.T) {
	tests := map[string]*Format{
		"somefile_v02.txt":       &Format{Name: "somefile.txt", Version: 2},
		"another.txt":            &Format{Name: "another.txt", Version: UNVERSIONED},
		"noext_v06":              &Format{Name: "noext", Version: 6},
		"files_galore_latest.go": &Format{Name: "files_galore.go", Version: LATEST},
	}

	for test, expected := range tests {
		// Test parsing
		format := new(SimpleFormat)
		format.Parse(test)
		if expected.Name != format.Name {
			fmt.Printf("'%s': Parsed name does not match. Expected: '%s', Got: '%s'\n", test, expected.Name, format.Name)
			t.Fail()
		}
		if expected.Version != format.Version {
			fmt.Printf("'%s': Parsed version does not match. Expected: '%d', Got: '%d'\n", test, expected.Version, format.Version)
			t.Fail()
		}

		// Test building
		format = &SimpleFormat{*expected}
		formatString := format.Build()
		if test != formatString {
			fmt.Printf("'%s': Built string does not match. Expected: '%s', Got: '%s'\n", test, test, formatString)
			t.Fail()
		}

	}
}
