package cola_statemachine_go

const (
	ExternalTransitionType = iota
	InternalTransitionType
)

type Transition struct {
	Src, Dst       *State
	Event          Event
	Condition      Condition
	Perform        Action
	TransitionType int
}
