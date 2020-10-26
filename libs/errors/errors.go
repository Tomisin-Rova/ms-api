package rerrors

import "google.golang.org/grpc/status"

type roavaError struct {
	ErrorString string `json:"error"`
	Message     string `json:"message"`
	Detail      string `json:"detail"`
	Help        string `json:"help"`
}

func NewFromGrpc(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		if err == nil {
			return &roavaError{Message: "unknown error"}
		}
		return &roavaError{Message: err.Error()}
	}
	return &roavaError{Message: st.Message(), ErrorString: st.Code().String()}
}

func (re *roavaError) Error() string {
	return re.Message
}