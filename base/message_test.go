package base_test

import (
	"context"
	"testing"
	"time"
	"bytes"

	"github.com/CoachApplication/command/base"
)

func TestMessageCommand_Id(t *testing.T) {
	com := base.NewMessageCommand("test")

	if com.Id() != "test" {
		t.Error("MessageCommand Command returned the wrong Id")
	}
}

func TestMessageCommand_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	out := bytes.NewBuffer([]byte{})

	com := base.NewMessageCommand("test")
	props := com.Properties()

	res := com.Validate(props)
	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():
		if res.Success() {
			t.Error("MessageCommand thinks it is valid when it shouldn't be.")
		}
	case <-ctx.Done():
		t.Error("MessageCommand Validate() timed out: ", ctx.Err().Error())
	}

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGE); err != nil {
		t.Error("MessageCommand did not give a Message Property")
	} else {
		messageProp.Set("this is the test message")
	}

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGEOUTPUT); err != nil {
		t.Error("MessageCommand did not give a Message Output Property")
	} else {
		messageProp.Set(out)
	}

	res = com.Validate(props)
	ctx, _ = context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("MessageCommand thinks it is invalid when it should be valid.")
		}
	case <-ctx.Done():
		t.Error("MessageCommand Validate() timed out: ", ctx.Err().Error())
	}

}

func TestMessageCommand_Exec(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	out := bytes.NewBuffer([]byte{})
	msg := []byte("this is the test message")

	com := base.NewMessageCommand("test")
	props := com.Properties()

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGE); err != nil {
		t.Error("MessageCommand did not give a Message Property")
	} else {
		messageProp.Set(msg)
	}

	if messageOutputProp, err := props.Get(base.PROPERTY_ID_MESSAGEOUTPUT); err != nil {
		t.Error("MessageCommand did not give a Message Output Property")
	} else {
		messageOutputProp.Set(out)
	}

	res := com.Exec(props)
	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("MessageCommand didn't Exec() successfully.")
		}

		if out.String() != string(msg) {
			t.Error("MessageCommand did not provide the correct message to the test output: ", out)
		}

	case <-ctx.Done():
		t.Error("MessageCommand Exec() timed out: ", ctx.Err().Error())
	}
}

func TestNewParametrizedMessageCommandProperties(t *testing.T) {
	dur, _ := time.ParseDuration("2s")

	out := bytes.NewBuffer([]byte{})
	msg := "this is the test message"
	props := base.NewParametrizedMessageCommandProperties([]byte(msg), out)

	if messageProp, err := props.Get(base.PROPERTY_ID_MESSAGE); err != nil {
		t.Error("MessageCommand did not give a Message Property")
	} else if getMsg, good := messageProp.Get().([]byte); !good {
		t.Error("Parametrized MessageCommand gave a bad MessageProperty value")
	} else if len(getMsg) == 0 {
		t.Error("Parametrized MessageCommand default Message Property was empty")
	} else if string(getMsg) != msg {
		t.Error("Parametrized MessageCommand default Message Property had the wrong value: ", getMsg)
	}

	if messageOutputProp, err := props.Get(base.PROPERTY_ID_MESSAGEOUTPUT); err != nil {
		t.Error("MessageCommand did not give a Message Output Property")
	} else {
		messageOutputProp.Set(out)
	}

	com := base.NewMessageCommand("test")

	res := com.Exec(props)
	ctx, _ := context.WithTimeout(context.Background(), dur)
	select {
	case <-res.Finished():
		if !res.Success() {
			t.Error("Parametrized MessageCommand didn't Exec() successfully.")
		}

		if out.String() != msg {
			t.Error("Parametrized MessageCommand did not provide the correct message to the test output: ", out)
		}

	case <-ctx.Done():
		t.Error("Parametrized MessageCommand Exec() timed out: ", ctx.Err().Error())
	}
}