package test

import (
	"fmt"
	cola "github.com/ermaoCode/cola-go"
	"github.com/ermaoCode/cola-go/builder"
	"testing"
)

type event int

const (
	Event1 event = iota
	Event2
)

func (e event) EventName() string {
	return [...]string{"Event1", "Event2"}[e]
}

var fooCondition = func(ctx cola.Ctx) bool {
	return ctx["foo"] == "foo"
}
var barCondition = func(ctx cola.Ctx) bool {
	return ctx["foo"] == "bar"
}

var fooFunc = func() { fmt.Println("hello world foo") }
var barFunc = func() { fmt.Println("hello world bar") }

func TestNewStateMachineBuilder(t *testing.T) {
	//condition1 := &staticCondition{}

	stateMachine1 := builder.NewStateMachineBuilder()
	stateMachine1.ExternalTransition().
		From("Start").To("End").On(Event1).When(fooCondition).Perform(fooFunc)
	stateMachine1.ExternalTransition().
		From("Start").To("BarEnd").On(Event1).When(barCondition).Perform(barFunc)
	stateMachine1.ExternalTransition().
		From("Start").To("end3").On(Event2).Perform(barFunc)
	machine1 := stateMachine1.Build("machine1")

	fmt.Println(machine1.PrintView(cola.FmtVisitor{}))
	fmt.Println(machine1.PrintView(cola.UmlVisitor{}))

	ctx := cola.Ctx{}
	ctx["foo"] = "bar"
	dstStateId, err := machine1.FireEvent("Start", Event1, ctx)
	if err != nil {
		t.Fail()
	}
	if dstStateId == "BarEnd" {
		fmt.Println("transition success, now in " + dstStateId)
	} else {
		t.Fail()
	}

	ctx["foo"] = "foo"
	dstStateId, err = machine1.FireEvent("Start", Event1, ctx)
	if err != nil {
		t.Fail()
	}
	if dstStateId == "End" {
		fmt.Println("transition success, now in " + dstStateId)
	} else {
		fmt.Println("transition fail, now in " + dstStateId)
		t.Fail()
	}

	dstStateId, err = machine1.FireEvent("Start", Event2, ctx)
	if err != nil {
		return
	}
	if dstStateId == "end3" {
		fmt.Println("transition success, now in " + dstStateId)
	} else {
		fmt.Println("transition fail, now in " + dstStateId)
		t.Fail()
	}
}

type staticCondition struct {
}

func (c *staticCondition) IsSatisfied(ctx map[string]interface{}) bool {
	return true
}
