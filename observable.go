package cola_statemachine_go

type Observable interface {
	PrintView(visitor Visitor) string
}
