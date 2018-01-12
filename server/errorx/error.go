package errorx

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code    int
	Message string
}

func (msg *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":      map[string]interface{}{"ret": msg.Code},
		"message": msg.Message,
	})
}

func (msg Error) Error() string {
	return fmt.Sprintf("%v/%v", msg.Code, msg.Message)
}

func newError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	EmptyToken    = newError(-200, "empty_token")
	EmptyMsg      = newError(-200, "empty_alter")
	EmptyUid      = newError(-200, "empty_uid")
	UnKnownDevice = newError(-200, "UnKnownDevice")
)