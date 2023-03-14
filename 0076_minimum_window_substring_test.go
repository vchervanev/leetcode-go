package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/minimum_window_substring/
func TestMinWindow(t *testing.T) {
	str := "AxxBxCxAB"
	template := "ABC"
	want := "CxAB"
	result := MinWindow(str, template)
	if result != want {
		t.Fatalf(`Got %q, expected %q`, result, want)
	}
}
