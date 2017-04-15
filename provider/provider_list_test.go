package provider_test

import (
	"testing"

	//command "github.com/CoachApplication/command"
	command_base "github.com/CoachApplication/command/base"
	command_provider "github.com/CoachApplication/command/provider"
)

func TestListProvider_Get(t *testing.T) {
	prov := command_provider.NewListProvider()

	prov.Add(command_base.NewTestCommand("one"))
	
	if getCmd, err := prov.Get("one"); err != nil {
		t.Error("ListPRovider returned an error when retrieving the Added Command")
	} else if getCmd.Id() != "one" {
		t.Error("ListProvider Get() Command gave the wrong Id()")
	}
}

func TestListProvider_Order(t *testing.T) {
	prov := command_provider.NewListProvider()

	prov.Add(command_base.NewMessageCommand("one"))
	prov.Add(command_base.NewTestCommand("two"))
	prov.Add(command_base.NewMessageCommand("three"))

	order := prov.Order()
	if len(order) == 0 {
		t.Error("ListProvider returned en empty list when Commands had been added")
	} else if len(order) != 3 {
		t.Error("ListProvider returned the wrong number of Command ids: ", order)
	} else if !(order[0]=="one" && order[1]=="two"&& order[2]=="three") {
		t.Error("ListProvider returned the Command ids in the wrong order: ", order)
	}
}
