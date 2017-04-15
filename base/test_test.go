package base_test

import (
	"context"
	"testing"
	"time"

	base_test "github.com/CoachApplication/base/test"
	"github.com/CoachApplication/command/base"
)


func TestTestCommandId(t *testing.T) {
	com := base.NewTestCommand("test")

	id := com.Id()
	if id != "test" {
		t.Error("TestCommand Id() returned the wrong value")
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Validate(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	com := base.NewTestCommand("test")

	props := com.Properties()
	if validProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONVALID); err != nil {
		t.Error("TestCommand Properties() did not provide a ValidProperty")
	} else {
		validProp.Set(false)
		res := com.Validate(props)
		select {
		case <-res.Finished():
			if res.Success() {
				t.Error("TestCommand thinks it is valid when it shouldn't be")
			}
		case <-ctx.Done():
			t.Error("TestCommand Validate timed out: ", ctx.Err().Error())
		}

		validProp.Set(true)
		res = com.Validate(props)
		select {
		case <-res.Finished():
			if !res.Success() {
				t.Error("TestCommand thinks it is invald when it shouldn't be")
			}
		case <-ctx.Done():
			t.Error("TestCommand Validate timed out: ", ctx.Err().Error())
		}
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Exec_Success(t *testing.T) {
	dur, _ := time.ParseDuration("2s")
	ctx, _ := context.WithTimeout(context.Background(), dur)

	com := base.NewTestCommand("test")

	props := com.Properties()
	if successProp, err := props.Get(base_test.PROPERTY_ID_OPERATIONSUCCESS); err != nil {
		t.Error("TestCommand Properties() did not provide a SuccessdProperty")
	} else {
		successProp.Set(false)
		res := com.Exec(props)
		select {
		case <-res.Finished():
			if res.Success() {
				t.Error("TestCommand Exec succeeds when it shouldn't")
			}
		case <-ctx.Done():
			t.Error("TestCommand Exec timed out: ", ctx.Err().Error())
		}

		successProp.Set(true)
		res = com.Exec(props)
		select {
		case <-res.Finished():
			if !res.Success() {
				t.Error("TestCommand Exec fails when it shouldn't")
			}
		case <-ctx.Done():
			t.Error("TestCommand Exec timedout: ", ctx.Err().Error())
		}
	}
}

// test our test command struct (and demonstrate what it is expected to do
func TestTestComand_Exec_Value(t *testing.T) {
	// com := NewTestCommand()

}