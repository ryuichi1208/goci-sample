package main

import (
	"testing"
	"time"
)

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

func TestCalc3(t *testing.T) {
	// テストケースの入力と期待される出力を定義します
	testCases := []struct {
		input    time.Time
		expected time.Time
	}{
		{
			input:    time.Date(2023, time.June, 11, 12, 0, 0, 0, time.UTC),
			expected: time.Date(2023, time.June, 11, 15, 0, 0, 0, time.UTC),
		},
		{
			input:    time.Date(2023, time.June, 11, 23, 45, 0, 0, time.UTC),
			expected: time.Date(2023, time.June, 12, 2, 45, 0, 0, time.UTC),
		},
		// 追加のテストケースを必要に応じて追加できます
	}

	// テストケースを順番に実行します
	for _, testCase := range testCases {
		// calc3 関数を実行して結果を取得します
		result := calc3(testCase.input)

		// 結果が期待される値と一致するかテストします
		if result != testCase.expected {
			t.Errorf("calc3(%v) の結果は %v であり、期待される結果 %v と一致しません", testCase.input, result, testCase.expected)
		}
	}
}

