package consensus

type state string

const (
	first  state = "first"
	second state = "second"
	third  state = "third"
)

var allowedState map[state][]state

func init() {
	allowedState = make(map[state][]state)
	allowedState[first] = []state{second, third}
	allowedState[second] = []state{third}
	allowedState[third] = []state{first}
}

func (s *state) CanTransition(next state) bool {

	for _, a := range allowedState[*s] {
		if a == next {
			return true
		}
	}
	return false
}

func (s *state) Transition(next state) {

	if s.CanTransition(next) {
		*s = next
	}
}
