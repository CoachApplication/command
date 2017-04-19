package base

import (
	"os/exec"

	"errors"
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	"github.com/CoachApplication/base/property"
	"github.com/CoachApplication/command"
	"os"
)

/**
 * ShellCommands are a bad idea:
 *
 * - in local cases they offer little advantage over just running the command without the tooling
 * - in remote cases they offer too much freedom on service hosts, no sandboxing;
 * - paths are not dependable across environments;
 * - tooling is not dependable across environments: e.g. see the difference in `man tr` for linux/osx.
 * The only advantage that I can see is that you could use environment variables from the tool to share
 * secrets into scripts/
 *
 * A containerized command makes much more sense"
 * - they can contain software applications that you don't want in your services
 * - they have standardized environments (based on predicable images)
 * - they can easily be sandboxed (in resources, file access, user access - despite security bugs in containers)
 * - they could be just as easy to define as a script (e.g. just run the script in a container)
 *
 */

/**
 * @TODO How to allow parametrization?
 *
 * See https://golang.org/pkg/os/exec/#Cmd
 *  Path string : command to run ($0)
 *  Dir string : working dir
 *  Env []string : environment variables
 *  Args []string : the command args (including $0 without path)
 *
 *  See also #CommandContext which also has a context
 */

// ShellCommand a command that runs a Shell Command locally.
type ShellCommand struct {
	id    string
	usage api.Usage
	ui    api.Ui

	cmd exec.Cmd
}

func NewShellCommand(id string, ui api.Ui, usage api.Usage, cmd exec.Cmd) *ShellCommand {
	return &ShellCommand{
		id:    id,
		ui:    ui,
		usage: usage,
		cmd:   cmd,
	}
}

func (sc *ShellCommand) Command() command.Command {
	return command.Command(sc)
}

func (sc *ShellCommand) Id() string {
	return sc.id
}

func (sc *ShellCommand) Ui() api.Ui {
	return sc.ui
}

func (sc *ShellCommand) Usage() api.Usage {
	return sc.usage
}

func (sc *ShellCommand) Properties() api.Properties {
	props := base.NewProperties()

	props.Add((&command.ArgsProperty{}).Property())
	props.Add((&command.EnvProperty{}).Property())

	return props.Properties()
}

func (sc *ShellCommand) Validate(props api.Properties) api.Result {
	res := base.NewResult()

	go func(cmd exec.Cmd, props api.Properties) {
		if cmd.Path == "" {
			res.MarkFailed()
		} else if _, err := exec.LookPath(cmd.Path); err != nil {
			res.AddError(err)
			res.MarkFailed()
		} else {
			res.MarkSucceeded()
		}

		res.MarkFinished()
	}(sc.cmd, props)

	return res.Result()
}

func (sc *ShellCommand) Exec(props api.Properties) api.Result {
	res := base.NewResult()

	go func(cmd exec.Cmd, props api.Properties) {
		if cmd.Stdout == nil {
			cmd.Stdout = os.Stdout
		}
		if cmd.Stderr == nil {
			cmd.Stderr = os.Stderr
		}
		if cmd.Stdin == nil {
			cmd.Stdin = os.Stdin
		}

		if prop, err := props.Get(command.PROPERTY_ID_COMMAND_ARGS); err == nil {
			vals := prop.Get().([]string)
			cmd.Args = append(cmd.Args, vals...)
		}
		if prop, err := props.Get(command.PROPERTY_ID_COMMAND_ENV); err == nil {
			vals := prop.Get().([]string)
			cmd.Env = append(cmd.Env, vals...)
		}

		if err := cmd.Start(); err != nil {
			res.AddError(err)
			res.AddError(errors.New("Command failed to start"))
			res.MarkFailed()
		} else if err := cmd.Wait(); err != nil {
			res.AddError(err)
			res.AddError(errors.New("Command failed to finish"))
			res.MarkFailed()
		} else {
			res.MarkFinished()
		}

		res.MarkFinished()
	}(sc.cmd, props)

	return res.Result()
}

/**
 * Properties for shell command
 */

const PROPERTY_ID_COMMAND_SHELL_EXEC = "command.executable"

type ShellExecutableProperty struct {
	property.StringProperty
}

func (sep *ShellExecutableProperty) Property() api.Property {
	return api.Property(sep)
}

func (sep *ShellExecutableProperty) Id() string {
	return PROPERTY_ID_COMMAND_SHELL_EXEC
}

func (sep *ShellExecutableProperty) Ui() api.Ui {
	return base.NewUi(
		sep.Id(),
		"Executable path",
		"Path to the executable",
		"",
	)
}

func (sep *ShellExecutableProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}
