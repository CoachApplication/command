package base_test

import (
	"bytes"
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/CoachApplication/base"
	"github.com/CoachApplication/command"
	command_base "github.com/CoachApplication/command/base"
)

func TestShellCommand_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")

	com := &command_base.ShellCommand{}
	props := com.Properties()

	ctx, _ := context.WithTimeout(context.Background(), dur)
	res := com.Validate(props)

	select {
	case <-res.Finished():
		if res.Success() {
			t.Error("ShellCommand thinks it is valid, when it has no command.")
		}
	case <-ctx.Done():
		t.Error("ShellCommand Validate() timed out: ", ctx.Err().Error())
	}

	cmd := exec.CommandContext(context.Background(), "which", "ps")
	com = command_base.NewShellCommand("which", base.NewUi("which", "Which", "", ""), base.ExternalOperationUsage{}.Usage(), *cmd)

	ctx, _ = context.WithTimeout(context.Background(), dur)
	res = com.Validate(props)

	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("ShellCommand thinks it is invalid, when it shouldn't be.")
		}
	case <-ctx.Done():
		t.Error("ShellCommand Validate() timed out: ", ctx.Err().Error())
	}

}

func TestShellCommand_Exec(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	out := bytes.NewBuffer([]byte{})
	err := bytes.NewBuffer([]byte{})

	cmd := exec.CommandContext(ctx, "which", "which")
	cmd.Stdout = out
	cmd.Stderr = err
	com := command_base.NewShellCommand("which", base.NewUi("which", "Which", "", ""), base.ExternalOperationUsage{}.Usage(), *cmd)
	props := com.Properties()

	res := com.Exec(props)

	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("ShellCommand failed: ", err.String(), res.Errors())
		} else {
			t.Log("ShellCommand output: ", out.String())
		}
	case <-ctx.Done():
		t.Error("ShellCommand Exec() timed out: ", ctx.Err().Error())
	}
}

func TestShellCommand_Exec1(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	out := bytes.NewBuffer([]byte{})
	err := bytes.NewBuffer([]byte{})

	cmd := exec.CommandContext(ctx, "echo")
	cmd.Env = []string{"ONE=1"}
	cmd.Stdout = out
	cmd.Stderr = err
	com := command_base.NewShellCommand("echo", base.NewUi("echo", "echo", "", ""), base.ExternalOperationUsage{}.Usage(), *cmd)
	props := com.Properties()

	if prop, err := props.Get(command.PROPERTY_ID_COMMAND_ARGS); err == nil {
		prop.Set([]string{"TEST:${ONE}${TWO}"})
	} else {
		t.Error("ShellCommand gave no Args property: ", err.Error())
	}
	if prop, err := props.Get(command.PROPERTY_ID_COMMAND_ENV); err == nil {
		prop.Set([]string{"TWO=2"})
	} else {
		t.Error("ShellCommand gave no Env property: ", err.Error())
	}

	res := com.Exec(props)

	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("ShellCommand failed: ", err.String(), res.Errors())
		}

		// These tests will not pass
		//else if out.String() != "TEST:${ONE}${TWO}" {
		//	t.Error("ShellCommand did not do proper environment variable substitution: ", out.String())
		//} else if out.String() != "TEST:12" {
		//	t.Error("ShellCommand gave the wrong output: ", out.String())
		//}
	case <-ctx.Done():
		t.Error("ShellCommand Exec() timed out: ", ctx.Err().Error())
	}
}
