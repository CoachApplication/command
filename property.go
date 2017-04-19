package command

import (
	"github.com/CoachApplication/api"
	"github.com/CoachApplication/base"
	base_errors "github.com/CoachApplication/base/errors"
	base_property "github.com/CoachApplication/base/property"
)

const (
	PROPERTY_ID_COMMAND_ID      = "command.id"
	PROPERTY_ID_COMMAND_FLAGS   = "command.flags"
	PROPERTY_ID_COMMAND_COMMAND = "command.command"
	PROPERTY_ID_COMMAND_ENV     = "command.env"
	PROPERTY_ID_COMMAND_ARGS    = "command.args"
)

// IdProperty Property for command Id string
type IdProperty struct {
	base_property.StringProperty
}

// Property explicitly convert this to a Property interface
func (ip *IdProperty) Property() api.Property {
	return api.Property(ip)
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (ip *IdProperty) Id() string {
	return PROPERTY_ID_COMMAND_ID
}

// Ui Provide UI metadata for the Property
func (ip *IdProperty) Ui() api.Ui {
	return base.NewUi(
		ip.Id(),
		"Command Id",
		"String Id for the command",
		"",
	).Ui()
}

// Usage Provide Usage information about the element
func (ip *IdProperty) Usage() api.Usage {
	return base.RequiredPropertyUsage{}.Usage()
}

// IdProperty Property for command Id string
type IdsProperty struct {
	base_property.StringSliceProperty
}

// Property explicitly convert this to a Property interface
func (ip *IdsProperty) Property() api.Property {
	return api.Property(ip)
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (ip *IdsProperty) Id() string {
	return PROPERTY_ID_COMMAND_ID
}

// Ui Provide UI metadata for the Property
func (ip *IdsProperty) Ui() api.Ui {
	return base.NewUi(
		ip.Id(),
		"Command Ids",
		"Slice of String Ids for available commands",
		"",
	).Ui()
}

// Usage Provide Usage information about the element
func (ip *IdsProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}.Usage()
}

type FlagsProperty struct {
	base_property.StringSliceProperty
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (fp *FlagsProperty) Id() string {
	return PROPERTY_ID_COMMAND_FLAGS
}

// Ui Provide UI metadata for the Property
func (fp *FlagsProperty) Ui() api.Ui {
	return base.NewUi(
		fp.Id(),
		"Command flags",
		"Ordered set of string flags for the command",
		"",
	).Ui()
}

// Usage Provide Usage information about the element
func (fp *FlagsProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

// IdProperty Property for command Id string
type CommandProperty struct {
	val Command
}

// Property explicitly convert this to a Property interface
func (cp *CommandProperty) Property() api.Property {
	return api.Property(cp)
}

// Id provides a machine name string for the property.  This should be unique in an operation.
func (cp *CommandProperty) Id() string {
	return PROPERTY_ID_COMMAND_COMMAND
}

// Ui Provide UI metadata for the Property
func (cp *CommandProperty) Ui() api.Ui {
	return base.NewUi(
		cp.Id(),
		"Command",
		"Commandd",
		"",
	).Ui()
}

// Usage Provide Usage information about the element
func (cp *CommandProperty) Usage() api.Usage {
	return base.ReadonlyPropertyUsage{}.Usage()
}

// Validate Check that the property is properly configured
func (cp *CommandProperty) Validate() bool {
	return cp.val == nil
}

// Get retrieve a value from the Property
func (cp *CommandProperty) Type() string {
	return "coach.command.Command"
}

// Get retrieve a value from the Property
func (cp *CommandProperty) Get() interface{} {
	return interface{}(cp.val)
}

// Set assign a value to the Property
func (cp *CommandProperty) Set(val interface{}) error {
	if typedVal, ok := val.(Command); !ok {
		return error(base_errors.PropertyWrongValueTypeError{Id: cp.Id(), Type: cp.Type(), Val: val})
	} else {
		cp.val = typedVal
		return nil
	}
}

type ArgsProperty struct {
	base_property.StringSliceProperty
}

func (sfp *ArgsProperty) Property() api.Property {
	return api.Property(sfp)
}

func (sfp *ArgsProperty) Id() string {
	return PROPERTY_ID_COMMAND_ARGS
}

func (sfp *ArgsProperty) Ui() api.Ui {
	return base.NewUi(
		sfp.Id(),
		"arguments",
		"Executable arguments",
		"",
	)
}

func (sfp *ArgsProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}

type EnvProperty struct {
	base_property.StringSliceProperty
}

func (sep *EnvProperty) Id() string {
	return PROPERTY_ID_COMMAND_ENV
}

func (sep *EnvProperty) Property() api.Property {
	return api.Property(sep)
}

func (sep *EnvProperty) Ui() api.Ui {
	return base.NewUi(
		sep.Id(),
		"environment variable",
		"Environment variables to use for execution",
		"",
	)
}

func (sep *EnvProperty) Usage() api.Usage {
	return base.OptionalPropertyUsage{}.Usage()
}
