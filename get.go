package command

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
)

const (
	OPERATION_ID_COMMAND_GET = "command.get"
)

// GetOperation Operation to retrieve a single Command
type GetOperation struct {}

// Id Return string Id for the operation
func (gop *GetOperation) Id() string {
	return OPERATION_ID_COMMAND_GET
}

// UI Return a UI interaction definition for the Operation
func (gop *GetOperation) Ui() api.Ui {
	return base.NewUi(
		gop.Id(),
		"Get a Command",
		"Retrieve a single Command",
		"",
	)
}

// Usage Define how the Operation is intended to be executed.
func (gop *GetOperation) Usage() api.Usage {
	return base.ExternalOperationUsage{}.Usage()
}

func (gop *GetOperation) Properties() api.Properties {
	props := base.NewProperties()

	props.Add((&IdProperty{}).Property())

	return props.Properties()
}

func (gop *GetOperation) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	if idProp, good := props.Get(PROPERTY_ID_COMMAND_ID); good {
		res.MarkFailed()
		res.AddError(error(base_errors.RequiredPropertyWasEmptyError{Key: PROPERTY_ID_COMMAND_ID}))
	} else if val, ok := idProp.Get().(string); !ok {
		res.MarkFailed()
		res.AddError(error(base_errors.PropertyWrongValueTypeError{Id: PROPERTY_ID_COMMAND_ID, ExpectedType: "string"}))
	} else if val == "" {
		res.MarkFailed()
		res.AddError(error(base_errors.RequiredPropertyWasEmptyError{Key: PROPERTY_ID_COMMAND_ID}))
	} else {
		res.MarkSucceeded()
	}
	res.MarkFinished()

	return res.Result()
}
