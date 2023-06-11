package main

import "testing"

// calc関数のテスト
func TestCalc(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{1, 2, 3},      // 1 + 2 = 3
		{-5, 10, 5},    // -5 + 10 = 5
		{0, 0, 0},      // 0 + 0 = 0
		{100, -100, 0}, // 100 + (-100) = 0
	}

	for _, testCase := range testCases {
		result := calc(testCase.x, testCase.y)
		if result != testCase.expected {
			t.Errorf("calc(%d, %d) = %d, expected %d", testCase.x, testCase.y, result, testCase.expected)
		}
	}
}
