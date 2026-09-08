package protoError

import (
	"fmt"
)

type ProtocolError struct {
	OpCode        byte
	ClientMsg     string
	InternalError error
}

func (e *ProtocolError) Error() string {
	return fmt.Sprintf("protocol error [%d]: %s; (internal: %v)", e.OpCode, e.ClientMsg, e.InternalError)
}

func (e *ProtocolError) Unwrap() error {
	return e.InternalError
}
