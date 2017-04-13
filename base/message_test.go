package base_test

import (
	"context"
	"github.com/CoachApplication/command/base"
	"testing"
	"time"
	"fmt"
)

func TestMessageCommand_Id(t *testing.T) {
	com := base.NewMessageCommand("test")

	if com.Id() != "test" {
		t.Error("MessageCommand Command returned the wrong Id")
	}
}

func TestMessageCommand_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	com := base.NewMessageCommand("test")
	props := com.Properties()

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGE); err != nil {
		t.Error("MessageCommand did not give a Message Property")
	} else {
		messageProp.Set("this is the test message")
	}

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGEOUTPUT); err != nil {
		t.Error("MessageCommand did not give a Message Output Property")
	} else {
		messageProp.Set(fmt.)
	}

}

func TestMessageCommand_Exec(t *testing.T) {

}
