package provider_test

import (
	"errors"

	command "github.com/CoachApplication/command"
	command_provider "github.com/CoachApplication/command/provider"
)

type TestProvider struct {
	coms []command.Command
}

func (tp *TestProvider) Provider() command_provider.Provider {
	return command_provider.Provider(tp)
}

func (tp *TestProvider) Add(com command.Command) {
	tp.coms = append(tp.coms, com)
}

func (tp *TestProvider) Get(id string) (command.Command, error) {
	for _, com := range tp.coms {
		if com.Id() == id {
			return com, nil
		}
	}
	return nil, errors.New("No matching command found")
}

func (tp *TestProvider) Order() []string {
	ids := []string{}
	for _, com := range tp.coms {
		ids = append(ids, com.Id())
	}
	return ids
}
