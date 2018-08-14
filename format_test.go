package main

import (
	"fmt"
	"testing"
)

func TestSimpleFormat(t *testing.T) {
	tests := map[string]*SimpleFormat{
		"somefile_v02.txt":       &SimpleFormat{name: "somefile.txt", version: 2},
		"another.txt":            &SimpleFormat{name: "another.txt", version: UNVERSIONED},
		"noext_v06":              &SimpleFormat{name: "noext", version: 6},
		"files_galore_latest.go": &SimpleFormat{name: "files_galore.go", version: LATEST},
	}

	for test, expected := range tests {
		// Test parsing
		format := new(SimpleFormat)
		format.Parse(test)
		if expected.GetName() != format.GetName() {
			fmt.Printf("'%s': Parsed name does not match. Expected: '%s', Got: '%s'\n", test, expected.GetName(), format.GetName())
			t.Fail()
		}
		if expected.GetVersion() != format.GetVersion() {
			fmt.Printf("'%s': Parsed version does not match. Expected: '%d', Got: '%d'\n", test, expected.GetVersion(), format.GetVersion())
			t.Fail()
		}

		// Test building
		format = expected
		formatString := format.Build()
		if test != formatString {
			fmt.Printf("'%s': Built string does not match. Expected: '%s', Got: '%s'\n", test, test, formatString)
			t.Fail()
		}

	}
}
