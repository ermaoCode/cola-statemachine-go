package cola_statemachine_go

import "strings"

type Visitor interface {
	VisitOnStateEntry(state *State) string
	VisitOnStateExit(state *State) string
	VisitOnStateMachineEntry(machine *StateMachine) string
	VisitOnStateMachineExit(machine *StateMachine) string
}

type FmtVisitor struct {
}

func (v FmtVisitor) VisitOnStateMachineEntry(state *StateMachine) string {
	return "------------ begin state machine -------------\n"
}

func (v FmtVisitor) VisitOnStateMachineExit(state *StateMachine) string {
	return "------------  end state machine -------------\n"
}

func (v FmtVisitor) VisitOnStateEntry(state *State) string {
	var sb strings.Builder
	for _, transitions := range state.EventTransition {
		for _, transition := range transitions {
			sb.WriteString(transition.Src.StateId + "[" + transition.Event.EventName() + "]")
			if transition.Condition != nil {
				sb.WriteString("[c]")
			}
			sb.WriteString("--->" + transition.Dst.StateId + "\n")
		}
	}
	return sb.String()
}

func (v FmtVisitor) VisitOnStateExit(state *State) string {
	return ""
}

type UmlVisitor struct {
}

func (v UmlVisitor) VisitOnStateMachineEntry(state *StateMachine) string {
	return "@startuml\n"
}

func (v UmlVisitor) VisitOnStateMachineExit(state *StateMachine) string {
	return "@enduml\n"
}

func (v UmlVisitor) VisitOnStateEntry(state *State) string {
	var sb strings.Builder
	for _, transitions := range state.EventTransition {
		for _, transition := range transitions {
			sb.WriteString("\"" + transition.Src.StateId + "\" -->[" + strings.ToLower(transition.Event.EventName()))
			if transition.Condition != nil {
				sb.WriteString(",c")
			}
			sb.WriteString("] \"" + transition.Dst.StateId + "\"\n")
		}
	}
	return sb.String()
}

func (v UmlVisitor) VisitOnStateExit(state *State) string {
	return ""
}
