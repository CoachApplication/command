package command

import (
	"testing"

	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	base_test "github.com/CoachApplication/base/test"
)

func TestTestCommandId(t *testing.T) {

}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Exec(t *testing.T) {
	// com := NewTestCommand()

}

const (
	OPERATION_ID_TEST = "command.test"
	PROPERTY_ID_TEST  = "command.test"
)

// TestCommand test command struct that just inputs and outputs string
type TestCommand struct {
	id        string
	propid    string
	propvalue string
}

// Constructor for TestCommand
func NewTestCommand(id, propid, propvalue string) *TestCommand {
	return &TestCommand{
		id:        id,
		propid:    propid,
		propvalue: propvalue,
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
	return base.MakeSuccessfulResult()
}

// Exec runs the operation from a Properties set, and return a result
func (tc *TestCommand) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties, propid string) {
		res.MarkFinished()
	}(props, tc.propid)

	return res.Result()
}
