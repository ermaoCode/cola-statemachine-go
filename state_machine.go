package cola_statemachine_go

import "strings"

type StateMachine struct {
	StateMachineId string
	StateMap       map[string]*State
}

func (stateMachine *StateMachine) FireEvent(srcStateId string, e Event, ctx Ctx) (dstStateId string, err error) {
	srcState, ok := stateMachine.StateMap[srcStateId]
	if !ok {
		return "", ErrStateNotExist
	}

	transitionList, ok := srcState.EventTransition[e]
	if !ok {
		return srcStateId, nil
	}

	for _, transition := range transitionList {
		if transition.Condition == nil {
			transition.Perform()
			return transition.Dst.StateId, nil
		} else if transition.Condition(ctx) {
			transition.Perform()
			return transition.Dst.StateId, nil
		}
	}
	return srcStateId, nil
}

func (stateMachine *StateMachine) GetState(stateId string) *State {
	return stateMachine.StateMap[stateId]
}

func (stateMachine *StateMachine) PrintView(visitor Visitor) string {
	var builder strings.Builder
	builder.WriteString(visitor.VisitOnStateMachineEntry(stateMachine))

	for _, state := range stateMachine.StateMap {
		builder.WriteString(visitor.VisitOnStateEntry(state))
		builder.WriteString(visitor.VisitOnStateExit(state))
	}

	builder.WriteString(visitor.VisitOnStateMachineExit(stateMachine))
	return builder.String()

}
