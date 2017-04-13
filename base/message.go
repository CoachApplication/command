package base

import (
	"fmt"
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/property"
)

const (
	PROPERTY_ID_MESSAGE       = "command.message"
	PROPERTY_ID_MESSAGEOUTPUT = "command.message.output"
)

// MessageCommand is a command that just outputs a message
type MessageCommand struct {
	id string
}

func NewMessageCommand(id string) *MessageCommand {
	return &MessageCommand{id: id}
}

func (mc *MessageCommand) Id() string {
	return mc.id
}

func (mc *MessageCommand) Ui() api.Ui {
	return base.NewUi(
		mc.Id(),
		"Message",
		"Command which will output a preprogrammed message.",
		"",
	)
}

func (mc *MessageCommand) Usage() api.Usage {
	return base.ExternalOperationUsage{}.Usage()
}

func (mc *MessageCommand) Properties() api.Properties {
	props := base.NewProperties()

	props.Add((&MessageProperty{}).Property())
	props.Add((&MessageOutputProperty{}).Property())

	return props.Properties()
}

func (mc *MessageCommand) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties) {
		if messageProp, err := props.Get(PROPERTY_ID_MESSAGE); err != nil {
			res.MarkFailed()
			res.AddError(err)
		} else if message := messageProp.Get().(string); message == "" {
			res.MarkFailed()
		} else {
			res.MarkSucceeded()
		}

		res.MarkFinished()
	}(props)

	return res.Result()
}

func (mc *MessageCommand) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(props api.Properties) {
		if messageProp, err := props.Get(PROPERTY_ID_MESSAGE); err != nil {
			res.MarkFailed()
			res.AddError(err)
		} else if message := messageProp.Get().(string); message == "" {
			res.MarkFailed()
		} else {
			fmt.Print(message)
		}

		res.MarkFinished()
	}(props)

	return res.Result()
}

type MessageProperty struct {
	property.StringProperty
}

func (mp *MessageProperty) Property() api.Property {
	return api.Property(mp)
}

func (mp *MessageProperty) Id() string {
	return PROPERTY_ID_MESSAGE
}

func (mp *MessageProperty) Ui() api.Ui {
	return base.NewUi(
		mp.Id(),
		"Command Message",
		"Message which will be output when the command is executed",
		"",
	)
}

func (mp *MessageProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}

type MessageOutputProperty struct {
	property.WriterProperty
}

func (mop *MessageOutputProperty) Property() api.Property {
	return api.Property(mop)
}

func (mop *MessageOutputProperty) Id() string {
	return PROPERTY_ID_MESSAGEOUTPUT
}

func (mop *MessageOutputProperty) Ui() api.Ui {
	return base.NewUi(
		mop.Id(),
		"Command Message Output writer",
		"Where the message will be output when the command is executed",
		"",
	)
}

func (mop *MessageOutputProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}
