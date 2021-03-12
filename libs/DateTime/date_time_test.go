package DateTime

import (
	"fmt"
	"testing"
	"time"
)

func TestMarshalID(t *testing.T) {
	tm := time.Now()

	o := MarshalDateTime(tm)
	if o == nil {
		t.Log("Cannot marshal DateTime")
		t.Fail()
	}

	fmt.Println(o, " Output")
}

func TestUnmarshalID(t *testing.T) {
	tm := "2020-04-25T05:11:40.346Z"

	o, err := UnmarshalDateTime(tm)
	if err != nil {
		t.Log("Cannot unmarshal DateTime")
		t.Fail()
	}

	fmt.Println(o, " Output")
}
