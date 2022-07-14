package builder

import cola "github.com/ermaoCode/cola-go"

func GetState(stateId string, stateMap map[string]*cola.State) *cola.State {
	if v, ok := stateMap[stateId]; ok {
		return v
	}
	state := &cola.State{
		StateId:         stateId,
		EventTransition: make(map[cola.Event][]*cola.Transition),
	}
	stateMap[stateId] = state
	return state
}
