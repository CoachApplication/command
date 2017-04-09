package command

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
)

const (
	OPERATION_ID_COMMAND_LIST = "command.list"
)

type ListOperation struct{}

// Id Return string Id for the operation
func (lop *ListOperation) Id() string {
	return OPERATION_ID_COMMAND_LIST
}

// UI Return a UI interaction definition for the Operation
func (lop *ListOperation) Ui() api.Ui {
	return base.NewUi(
		lop.Id(),
		"Get a Command",
		"Retrieve a single Command",
		"",
	)
}

// Usage Define how the Operation is intended to be executed.
func (lop *ListOperation) Usage() api.Usage {
	return base.ExternalOperationUsage{}.Usage()
}
