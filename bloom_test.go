package main

import "testing"

func TestBloomFilter(t *testing.T) {
	bf, _ := NewBloomFilter(100, 0.1)

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

	// Add elements to the filter
	bf.Add("apple")
	bf.Add("cherry")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := bf.Exist(tc.value)
			if actual != tc.expected {
				t.Errorf("Expected Test() to return %v, but got %v", tc.expected, actual)
			}
		})
	}
}
