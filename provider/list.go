package provider

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/command"
)

type ListOperation struct {
	command.ListOperation

	provider Provider
}

func (lo *ListOperation) Operation() api.Operation {
	return api.Operation(lo)
}

func (lo *ListOperation) Properties() api.Properties {
	return base.NewProperties().Properties()
}

func (lo *ListOperation) Validate(props api.Properties) api.Result {
	if lo.provider == nil {
		return base.MakeFailedResult()
	} else {
		return base.MakeSuccessfulResult()
	}
}

func (lo *ListOperation) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(provider Provider) {
		idsProp := command.IdsProperty{}
		idsProp.Set(provider.Order())
		res.AddProperty(idsProp.Property())

		res.MarkSucceeded()
		res.MarkFinished()
	}(lo.provider)

	return res.Result()
}
