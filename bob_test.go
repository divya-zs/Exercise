package bob

import (
	"fmt"
	"testing"
)

func TestBob(t *testing.T) {
	testcase := []struct {
		str      string
		expected string
	}{
		{"How are you?", "Sure."},
		{"YELL AT HIM", "Whoa, chill out!"},
		{"HOW ARE YOU?", "Calm down, I know what I'm doing"},
		{"", "Fine. Be that way!"},
		{"\n", "Fine. Be that way!"},
	}
	for _, tc := range testcase {
		output := MyFunction(tc.str)
		if output != tc.expected {
			fmt.Printf("Expected %v but got %v\n", tc.expected, output)
		}
	}
}
