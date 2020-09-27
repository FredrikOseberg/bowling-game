package main

import (
	"testing"
)

type TestCases struct {
	Input  []int
	Result int
}

func TestGetFrames(t *testing.T) {
	cases := []TestCases{
		{
			Input:  []int{2, 3, 5, 4, 9, 1, 2, 5, 3, 2, 4, 2, 3, 3, 4, 6, 10, 3, 2},
			Result: 10,
		},
		{
			Input:  []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			Result: 10,
		},
		{
			Input:  []int{1, 3, 8, 2, 1, 4},
			Result: 0,
		},
		{
			Input:  []int{2, 3, 5, 4, 9, 1, 2, 5, 3, 2, 4, 2, 3, 3, 4, 6, 10, 3, 2, 1, 2, 3, 5, 6, 7},
			Result: 0,
		},
		{
			Input:  []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			Result: 0,
		},
	}

	for _, testCase := range cases {
		result := getFrames(testCase.Input)
		if len(result) != testCase.Result {
			t.Errorf("Wanted: %d, Received: %d", testCase.Result, len(result))
		}
	}
}
func TestCalculateResult(t *testing.T) {
	cases := []TestCases{
		{
			Input:  []int{2, 3, 5, 4, 9, 1, 2, 5, 3, 2, 4, 2, 3, 3, 4, 6, 10, 3, 2},
			Result: 90,
		},
		{
			Input:  []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			Result: 300,
		},
		{
			Input:  []int{1, 5, 8, 2, 2, 4, 10, 10, 5, 3, 9, 1, 10, 10, 10, 4, 3},
			Result: 166,
		},
	}

	for _, testCase := range cases {
		frames := getFrames(testCase.Input)
		_, totalScore := calculateScore(frames, testCase.Input)

		if totalScore != testCase.Result {
			t.Errorf("Wanted: %d, Received: %d", testCase.Result, totalScore)
		}
	}
}
