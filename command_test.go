package command

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	base_test "github.com/CoachApplication/base/test"
)

func TestTestCommandId(t *testing.T) {
	com := NewTestCommand("test")

	id := com.Id()
	if id != "test" {
		t.Error("TestCommand Id() returned the wrong value")
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	com := NewTestCommand("test")

	props := com.Properties()
	if validProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONVALID); err != nil {
		t.Error("TestCommand Properties() did not provide a ValidProperty")
	} else {
		validProp.Set(false)
		res := com.Validate(props)
		select {
		case <-res.Finished():
			if res.Success() {
				t.Error("TestCommand thinks it is valid when it shouldn't be")
			}
		case <-ctx.Done():
			t.Error("TestCommand Validate timed out: ", ctx.Err().Error())
		}

		validProp.Set(true)
		res = com.Validate(props)
		select {
		case <-res.Finished():
			if !res.Success() {
				t.Error("TestCommand thinks it is invald when it shouldn't be")
			}
		case <-ctx.Done():
			t.Error("TestCommand Validate timed out: ", ctx.Err().Error())
		}
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Exec_Success(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	com := NewTestCommand("test")

	props := com.Properties()
	if successProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONSUCCESS); err != nil {
		t.Error("TestCommand Properties() did not provide a SuccessdProperty")
	} else {
		successProp.Set(false)
		res := com.Exec(props)
		select {
		case <-res.Finished():
			if res.Success() {
				t.Error("TestCommand Exec succeeds when it shouldn't")
			}
		case <-ctx.Done():
			t.Error("TestCommand Exec timed out: ", ctx.Err().Error())
		}

		successProp.Set(true)
		res = com.Exec(props)
		select {
		case <-res.Finished():
			if !res.Success() {
				t.Error("TestCommand Exec fails when it shouldn't")
			}
		case <-ctx.Done():
			t.Error("TestCommand Exec timedout: ", ctx.Err().Error())
		}
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Exec_Value(t *testing.T) {
	// com := NewTestCommand()

}

const (
	OPERATION_ID_TEST = "command.test"
	PROPERTY_ID_TEST  = "command.test"
)

// TestCommand test command struct that just inputs and outputs string
type TestCommand struct {
	id string
}

// Constructor for TestCommand
func NewTestCommand(id string) *TestCommand {
	return &TestCommand{
		id: id,
	}
}

// Command explicitly convert this into a Command interface
func (tc *TestCommand) Command() Command {
	return Command(tc)
}

// Id Unique string machine name identifier for the Operation
func (tc *TestCommand) Id() string {
	return tc.id
}

// UI Return a UI interaction definition for the Operation
func (tc *TestCommand) Ui() api.Ui {
	return base.NewUi(
		tc.Id(),
		"Test Command",
		"Test Command",
		"",
	)
}

// Usage Define how the Operation is intended to be executed.
func (tc *TestCommand) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// Properties provide the expected Operation with default values
func (tc *TestCommand) Properties() api.Properties {
	success := &base_test.SuccessfulOperationProperty{}
	success.Set(true)
	valid := &base_test.ValidOperationProperty{}
	valid.Set(true)
	val := &base_test.TestProperty{}

	props := base.NewProperties()
	props.Add(success.Property())
	props.Add(valid.Property())
	props.Add(val.Property())

	return props.Properties()
}

// Validate Validate that the Operation can Execute if passed proper Property data
func (tc *TestCommand) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties) {
		if validProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONVALID); err == nil {
			if validProp.Get().(bool) {
				res.MarkSucceeded()
			} else {
				res.MarkFailed()
			}
		} else {
			res.AddError(err)
			res.AddError(errors.New("No valid property was passed.  Leaving the operation as successful"))
		}

		res.MarkFinished()
	}(props)

	return res.Result()
}

// Exec runs the operation from a Properties set, and return a result
func (tc *TestCommand) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties) {
		if successProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONSUCCESS); err == nil {
			if successProp.Get().(bool) {
				res.MarkSucceeded()
			} else {
				res.MarkFailed()
			}
		} else {
			res.AddError(err)
			res.AddError(errors.New("No success property was passed.  Leaving the operation as successful"))
		}

		res.MarkFinished()
	}(props)

	return res.Result()
}
