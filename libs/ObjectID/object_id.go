package ObjectID

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type ID = primitive.ObjectID

func New() ID {
	return primitive.NewObjectIDFromTimestamp(time.Now())
}
func MarshalID(id primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		b, _ := id.MarshalJSON()
		_, _ = w.Write(b)
	})
}

func UnmarshalID(v interface{}) (ID, error) {
	switch v := v.(type) {
	case string:
		_id, _ := primitive.ObjectIDFromHex(v)
		return _id, nil
	case primitive.ObjectID:
		return v, nil
	default:
		return primitive.NilObjectID, fmt.Errorf("%T is not a primitive.ObjectID", v)
	}
}
