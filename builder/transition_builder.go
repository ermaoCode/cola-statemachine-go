package builder

import cola "github.com/ermaoCode/cola-go"

type TransitionBuilder struct {
	Transition *cola.Transition
	StateMap   map[string]*cola.State
}

func (t *TransitionBuilder) From(stateId string) From {
	state := GetState(stateId, t.StateMap)
	t.Transition.Src = state
	return t
}

func (t *TransitionBuilder) To(stateId string) To {
	state := GetState(stateId, t.StateMap)
	t.Transition.Dst = state
	return t
}

func (t *TransitionBuilder) On(e cola.Event) On {
	t.Transition.Event = e
	transitionList, ok := t.Transition.Src.EventTransition[e]
	if ok {
		transitionList = append(transitionList, t.Transition)
		t.Transition.Src.EventTransition[e] = transitionList
	} else {
		t.Transition.Src.EventTransition[e] = make([]*cola.Transition, 1, 1)
		t.Transition.Src.EventTransition[e][0] = t.Transition
	}
	return t
}

func (t *TransitionBuilder) When(c cola.Condition) *TransitionBuilder {
	t.Transition.Condition = c
	return t
}

func (t *TransitionBuilder) Perform(a cola.Action) *TransitionBuilder {
	t.Transition.Perform = a
	return t
}
