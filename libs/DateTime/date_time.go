package DateTime

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)

type DateTime = time.Time

func New() DateTime {
	return time.Now()
}
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		b, _ := t.MarshalJSON()
		_, _ = w.Write(b)
	})
}

func UnmarshalDateTime(v interface{}) (DateTime, error) {
	switch v := v.(type) {
	case string:
		t, _ := time.Parse(time.RFC3339, v)
		return t, nil
	case time.Time:
		return v, nil
	default:
		return DateTime{}, fmt.Errorf("%T is not a DateTime", v)
	}
}
