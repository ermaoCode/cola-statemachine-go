package builder

import (
	cola "github.com/ermaoCode/cola-go"
)

type StateMachineBuilder struct {
	StateMap     map[string]*cola.State
	StateMachine *cola.StateMachine
}

func (s *StateMachineBuilder) ExternalTransition() ExternalBuilder {
	t := &TransitionBuilder{
		Transition: &cola.Transition{
			Src:            &cola.State{},
			Dst:            &cola.State{},
			Event:          nil,
			Condition:      nil,
			Perform:        nil,
			TransitionType: cola.ExternalTransitionType,
		},
		StateMap: s.StateMap,
	}
	return t
}

func (s *StateMachineBuilder) InternalTransition() *TransitionBuilder {
	t := &TransitionBuilder{
		Transition: &cola.Transition{
			Src:            &cola.State{},
			Dst:            &cola.State{},
			Event:          nil,
			Condition:      nil,
			Perform:        nil,
			TransitionType: cola.InternalTransitionType,
		},
	}
	return t
}

func (s *StateMachineBuilder) Build(stateMachineId string) *cola.StateMachine {
	s.StateMachine.StateMachineId = stateMachineId
	s.StateMachine.StateMap = s.StateMap

	return s.StateMachine
}
