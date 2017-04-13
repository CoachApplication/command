package provider

import (
	"github.com/CoachApplication/command"
)

// ListProvider a class which can be used to create a list of Command interface
type ListProvider struct {
	coms []command.Command
}

// NewListProvider Constructor for ListProvider
func NewListProvider() *ListProvider {
	return &ListProvider{}
}

func (tp *ListProvider) Provider() Provider {
	return Provider(tp)
}

func (tp *ListProvider) Add(com command.Command) {
	tp.coms = append(tp.coms, com)
}

func (tp *ListProvider) Get(id string) (command.Command, error) {
	for _, com := range tp.coms {
		if com.Id() == id {
			return com, nil
		}
	}
	return nil, error(CommandNotFoundError{Id: id})
}

func (tp *ListProvider) Order() []string {
	ids := []string{}
	for _, com := range tp.coms {
		ids = append(ids, com.Id())
	}
	return ids
}
