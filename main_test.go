package main

import "testing"

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestWordBreakSimple(t *testing.T) {
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}
	mustMatch := []string{"cat sand dog", "cats and dog"}

	sentences := wordBreak(s, dict)

	if len(sentences) == 0 {
		t.Error("Expected to find 2 matches.")
	}

	if testEq(sentences, mustMatch) == false {
		t.Error("Found sentences do not match", sentences, mustMatch)
	}
}

func TestWordBreakIncorrect(t *testing.T) {
	s := "-catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}

	sentences := wordBreak(s, dict)

	if len(sentences) > 0 {
		t.Error("Expected not to be able to break given string,", s)
	}
}
