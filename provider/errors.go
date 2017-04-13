package provider

import "fmt"

type CommandNotFoundError struct {
	Id string
}

func (cnfe CommandNotFoundError) Error() string {
	return fmt.Sprintf("Requested Command [%s] was not found", cnfe.Id)
}
