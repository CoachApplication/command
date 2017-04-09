package provider

import (
	command "github.com/CoachApplication/command"
)

type Provider interface {
	Get(id string) (command.Command, error)
	Order() []string
}
