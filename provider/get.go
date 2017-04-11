package provider

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
	"github.com/CoachApplication/command"
)

type GetOperation struct {
	command.GetOperation

	provider Provider
}

func (gop *GetOperation) Operation() api.Operation {
	return api.Operation(gop)
}

func (gop *GetOperation) Properties() api.Properties {
	props := base.NewProperties()

	props.Add((&command.IdProperty{}).Property())

	return props.Properties()
}

func (gop *GetOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties) {
		if idProp, err := props.Get(command.PROPERTY_ID_COMMAND_ID); err != nil {
			res.MarkFailed()
			res.AddError(err)
			res.AddError(error(base_errors.RequiredPropertyWasEmptyError{Key: command.PROPERTY_ID_COMMAND_ID}))
		} else if val, ok := idProp.Get().(string); !ok {
			res.MarkFailed()
			res.AddError(error(base_errors.PropertyWrongValueTypeError{Id: command.PROPERTY_ID_COMMAND_ID, Type: "string", Val: val}))
		} else if val == "" {
			res.MarkFailed()
			res.AddError(error(base_errors.RequiredPropertyWasEmptyError{Key: command.PROPERTY_ID_COMMAND_ID}))
		} else if com, err := gop.provider.Get(val); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			comProp := &command.CommandProperty{}
			comProp.Set(com)

			res.AddProperty(comProp.Property())
			res.MarkSucceeded()
		}

		res.MarkFinished()
	}(props)

	return res.Result()
}
