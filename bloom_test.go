package main

import (
	"testing"
)

func TestBloomFilterAdd(t *testing.T) {
	bf, _ := NewBloomFilter(100, 0.1)
	items := []int{1, 2, 3, 10, 20, 30}
	for _, item := range items {
		err := bf.Add(item)
		if err != nil {
			t.Errorf("Add item returns error :%v", err)
		}
	}
	if bf.n != len(items) {
		t.Errorf("Expected item count :%v, but got %v", len(items), bf.n)
	}
}

func TestBloomFilter(t *testing.T) {
	bf, _ := NewBloomFilter(100, 0.1)
	items := []string{"apple", "cherry"}

	testCases := []struct {
		name     string
		value    string
		expected bool
	}{
		{name: "Test1", value: "apple", expected: true},
		{name: "Test2", value: "banana", expected: false},
		{name: "Test3", value: "cherry", expected: true},
		{name: "Test4", value: "date", expected: false},
	}

	for _, item := range items {
		err := bf.Add(item)
		if err != nil {
			t.Errorf("Add item returns error :%v", err)
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := bf.Exist(tc.value)
			if err != nil {
				t.Errorf("Exist() return error: %v", err)
			}
			if actual != tc.expected {
				t.Errorf("Expected Exist() to return %v, but got %v", tc.expected, actual)
			}
		})
	}
}
