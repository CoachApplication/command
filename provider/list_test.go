package provider_test

import (
	"time"
	"context"
	"testing"

	"github.com/CoachApplication/api"
	command_provider "github.com/CoachApplication/command/provider"
	command_base "github.com/CoachApplication/command/base"
)

func makeListOperation() api.Operation{
	prov := command_provider.NewListProvider()
	prov.Add(command_base.NewMessageCommand("one"))
	prov.Add(command_base.NewTestCommand("two"))
	prov.Add(command_base.NewMessageCommand("three"))

	return command_provider.NewListOperation(prov).Operation()
}

func TestListOperation_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")

	listOp := command_provider.NewListOperation(nil).Operation()

	ctx, _ := context.WithTimeout(context.Background(), dur)
	res := listOp.Validate(listOp.Properties())
	select {
	case <-res.Finished():
		if res.Success() {
			t.Error("ListOperation thinks it is valid when it has no Provider")
		}
	case <-ctx.Done():
		t.Error("ListOperation Validate() timed out: ", ctx.Err().Error())
	}

	listOp = makeListOperation()

	ctx, _ = context.WithTimeout(context.Background(), dur)
	res = listOp.Validate(listOp.Properties())
	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("ListOperation thinks it is invalid when it has a Provider")
		}
	case <-ctx.Done():
		t.Error("ListOperation Validate() timed out: ", ctx.Err().Error())
	}
}

func TestListOperation_Exec(t *testing.T) {

}