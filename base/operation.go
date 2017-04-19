package base

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/command"
)

// OperationCommand A Command that wraps an Operation
type OperationCommand struct {
	op api.Operation
}

func NewOperationCommand(op api.Operation) *OperationCommand {
	return &OperationCommand{
		op: op,
	}
}

func (oc *OperationCommand) Command() command.Command {
	return command.Command(oc)
}

func (oc *OperationCommand) Id() string {
	return oc.op.Id()
}

func (oc *OperationCommand) Ui() api.Ui {
	return oc.op.Ui()
}

func (oc *OperationCommand) Usage() api.Usage {
	return oc.op.Usage()
}

func (oc *OperationCommand) Properties() api.Properties{
	return oc.op.Properties()
}

func (oc *OperationCommand) Validate(props api.Properties) api.Result {
	return oc.op.Validate(props)
}

func (oc *OperationCommand) Exec(props api.Properties) api.Result {
	return oc.op.Exec(props)
}
