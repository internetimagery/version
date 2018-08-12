package main

import (
	"fmt"
	"testing"
)

func TestSimpleFormat(t *testing.T) {
	tests := map[string]*Format{
		"somefile_v02.txt": &Format{Name: "somefile.txt", Version: 2},
	}

	for test, expected := range tests {
		// Test parsing
		format := new(SimpleFormat)
		format.Parse(test)
		if expected.Name != format.Name {
			fmt.Printf("Parsed name does not match. Expected: '%s', Got: '%s'\n", expected.Name, format.Name)
			t.Fail()
		}
		if expected.Version != format.Version {
			fmt.Printf("Parsed version does not match. Expected: '%d', Got: '%d'\n", expected.Version, format.Version)
			t.Fail()
		}

		// Test building
		format = &SimpleFormat{*expected}
		formatString := format.Build()
		if test != formatString {
			fmt.Printf("Built string does not match. Expected: '%s', Got: '%s'\n", test, formatString)
			t.Fail()
		}

	}
}
