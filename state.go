package cola_statemachine_go

type State struct {
	StateId         string
	EventTransition map[Event][]*Transition
}
