package builder

import cola "github.com/ermaoCode/cola-go"

func NewStateMachineBuilder() *StateMachineBuilder {
	return &StateMachineBuilder{
		StateMap:     make(map[string]*cola.State),
		StateMachine: &cola.StateMachine{},
	}
}
