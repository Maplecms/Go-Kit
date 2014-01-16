package tests

import (
	"testing"
	"fmt"
	"github.com/refulgentsoftware/AmburWebServices/activation"
)

func TestNewUniqueID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		uuid, err := activation.NewUniqueID()
		if err != nil {
			t.Log("Could not generate regex")
			t.Fail()
		}
		if ! activation.UniqueIDRegex.MatchString(uuid.String()) {
			t.Log(uuid.String())
			t.Fail()
		}
	}
}

func TestUnmarshalText(t *testing.T) {
	uuidString := "30bbfac5-a5a6-4990-85c2-9ba81fc11ba7"
	uuid := new (activation.UniqueID)
	err := uuid.UnmarshalText([]byte(uuidString))
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	if uuid.String() != uuidString {
		t.Log(fmt.Sprintf("uuid %s does not match original string %s", uuid.String(), uuidString))
		t.Fail()
	}
}



















