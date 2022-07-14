package builder

import cola "github.com/ermaoCode/cola-go"

type ExternalBuilder interface {
	From(stateId string) From
}

type From interface {
	To(stateId string) To
}

type To interface {
	On(e cola.Event) On
}

type On interface {
	When(c cola.Condition) *TransitionBuilder
	Perform(a cola.Action) *TransitionBuilder
}
