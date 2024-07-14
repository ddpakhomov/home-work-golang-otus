package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	text := "This is a test text. The test text will be used for unit testing."
	expected := map[string]int{
		"this":    1,
		"is":      1,
		"a":       1,
		"test":    2,
		"text":    2,
		"the":     1,
		"will":    1,
		"be":      1,
		"used":    1,
		"for":     1,
		"unit":    1,
		"testing": 1,
	}

	result := countWords(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, result)
	}
}
