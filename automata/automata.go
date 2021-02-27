package automata

// FA represents Finite Automaton
type FA struct {
	initial    int
	states     []int
	final      []int
	symbols    []rune
	transition map[int]map[rune]int
}

// NewFA creates finite automaton instance
func NewFA() FA {
	return FA{
		initial:    0,
		states:     []int{0},
		transition: make(map[int]map[rune]int),
	}
}

// WithStateNum returns new FA instance with some state
func (a FA) WithStateNum(states ...int) FA {
	a.states = states
	return a
}

// WithFinal returns FA with finalStates
func (a FA) WithFinal(finalStates ...int) FA {
	var final []int
	finalSet := make(map[int]struct{})
	for _, s := range finalStates {
		if _, ok := finalSet[s]; ok {
			continue
		}
		finalSet[s] = struct{}{}
		final = append(final, s)
	}
	a.final = final
	return a
}

// WithSymbols returns FA with input symbols ss
func (a FA) WithSymbols(ss string) FA {
	var symbols []rune
	symbolSet := make(map[rune]struct{})
	for _, s := range []rune(ss) {
		if _, ok := symbolSet[s]; ok {
			continue
		}
		symbols = append(symbols, s)
		symbolSet[s] = struct{}{}
	}
	a.symbols = symbols
	return a
}

// WithTransition returns FA with transition function(mapping) ts
func (a FA) WithTransition(ts map[int]map[rune]int) FA {
	a.transition = ts
	return a
}

// Recognize receives string and return if the string is accepted by the finite automaton.
func (a FA) Recognize(s string) bool {
	currentState := a.initial
	for _, c := range []rune(s) {
		nextState, ok := a.transition[currentState][c]
		if !ok {
			return false
		}
		currentState = nextState
	}
	for _, fs := range a.final {
		if fs == currentState {
			return true
		}
	}
	return false
}
