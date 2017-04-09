package provider_test

import (
	"github.com/CoachApplication/api"
	command "github.com/CoachApplication/command"
	command_provider "github.com/CoachApplication/command/provider"
	"errors"
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


type TestCommand struct {

}

// Id Unique string machine name identifier for the Operation
func (tc *TestCommand) Id() string {

}

// UI Return a UI interaction definition for the Operation
func (tc *TestCommand) Ui() api.Ui {

}

// Usage Define how the Operation is intended to be executed.
func (tc *TestCommand) Usage() api.Usage {

}

// Properties provide the expected Operation with default values
func (tc *TestCommand) Properties() api.Properties {

}

// Validate Validate that the Operation can Execute if passed proper Property data
func (tc *TestCommand) Validate(props api.Properties) api.Result {

}

/**
 * Exec runs the operation from a Properties set, and return a result
 *
 * Exec is expected to handle any forking needed internally, and to pass any response Property changes via the
 * api_result.Result.Properties() method.
 *
 * Exec receives a Properties list that it should consider disposable.  This is by design so that any Operation
 * consumer can reused the Properties object for subsequent calls, which may run in parrallel.
 */
func (tc *TestCommand) Exec(props api.Properties) api.Result {

}
