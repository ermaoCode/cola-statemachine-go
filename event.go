package cola_statemachine_go

// Event 使用的时候使用 iota 枚举值即可
//type Event int

type Event interface {
	EventName() string
}
