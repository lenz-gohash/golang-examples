package manipulate

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplit(t *testing.T) {
	// A struct representing a test case.

	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		// Here each `key: value` is a test case.
		//
		// The key is the name of the test case.
		// The value is the struct that contains the test data.
		"simple":       {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	// - `name` is the key, `tc` is the test input value
	for name, tc := range tests {
		// t.Run allows for running tests in parralel, without stopping the test suite when encountering t.Fatal
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
