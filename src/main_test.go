package main

import (
	"strings"
	"testing"
)

func TestScanner_ValidInput(t *testing.T) {
	input := "1\n2\n3\n"
	r := strings.NewReader(input)

	expected := []int{1, 2, 3}
	result, err := Scanner(r)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], result[i])
		}
	}
}

func TestScanner_InvalidInput(t *testing.T) {
	input := "abc\n"
	r := strings.NewReader(input)

	_, err := Scanner(r)
	if err != ErrNotInteger {
		t.Fatalf("expected %v, got %v", ErrNotInteger, err)
	}
}

func TestScanner_OutOfRange(t *testing.T) {
	input := "100001\n"
	r := strings.NewReader(input)

	_, err := Scanner(r)
	if err != ErrOutOfRange {
		t.Fatalf("expected %v, got %v", ErrOutOfRange, err)
	}
}

func TestScanner_EmptyInput(t *testing.T) {
	input := ""
	r := strings.NewReader(input)

	_, err := Scanner(r)
	if err != ErrEmptyInput {
		t.Fatalf("expected %v, got %v", ErrEmptyInput, err)
	}
}
