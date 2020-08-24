package ObjectID

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestMarshalID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5ea26855ad5d91230a891295")

	o := MarshalID(id)
	if o == nil {
		t.Log("Cannot marshal ID")
		t.Fail()
	}
}

func TestUnmarshalID(t *testing.T) {
	id := "5ea26855ad5d91230a891295"

	_, err := UnmarshalID(id)
	if err != nil {
		t.Log("Cannot unmarshal ID")
		t.Fail()
	}
}
