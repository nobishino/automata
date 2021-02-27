package automata_test

import (
	"testing"

	"github.com/nobishino/automata/automata"
)

func TestProcessString(t *testing.T) {
	fa := automata.NewFA().WithStateNum(2).WithFinal(1).WithSymbols("01").WithTransition(
		map[int]map[rune]int{
			0: {
				'0': 2,
				'1': 1,
			},
			1: {
				'0': 0,
				'1': 2,
			},
			2: {
				'0': 2,
				'1': 2,
			},
		})
	_ = fa
}

func TestNewFA(t *testing.T) {
	_ = automata.NewFA()
}

func TestNewFAWithState(t *testing.T) {
	_ = automata.NewFA().WithStateNum(2)
}

func TestRecognize(t *testing.T) {
	a := makeSampleFA()
	testcases := [...]struct {
		input          string
		shouldAccepted bool
	}{
		{
			input:          "1",
			shouldAccepted: true,
		},
		{
			input:          "101",
			shouldAccepted: true,
		},
		{
			input:          "10101",
			shouldAccepted: true,
		},
		{
			input:          "10",
			shouldAccepted: false,
		},
		{
			input:          "1010",
			shouldAccepted: false,
		},
		{
			input:          "1011",
			shouldAccepted: false,
		},
		{
			input:          "101011",
			shouldAccepted: false,
		},
	}
	for _, tt := range testcases {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got := a.Recognize(tt.input)
			if got != tt.shouldAccepted {
				t.Errorf("expect %v, but got %v", tt.shouldAccepted, got)
			}
		})
	}

}

// 1(01)* を受理する有限オートマトンを返す
func makeSampleFA() automata.FA {
	return automata.NewFA().WithStateNum(2).WithFinal(1).WithSymbols("01").WithTransition(
		map[int]map[rune]int{
			0: {
				'0': 2,
				'1': 1,
			},
			1: {
				'0': 0,
				'1': 2,
			},
			2: {
				'0': 2,
				'1': 2,
			},
		})
}
