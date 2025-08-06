package sdk

import "fmt"

type SDKError struct {
	Code    uint32
	Meaning string
	Message string
}

func (e SDKError) Error() string {
	return fmt.Sprintf(`Code: %v; Meaning: %v; Message: %v`, e.Code, e.Meaning, e.Message)
}

var ErrorCode000 = SDKError{Code: 0, Meaning: "Failed to send request to Node. Node might be offline", Message: ""}
var ErrorCode001 = SDKError{Code: 1, Meaning: "Response returned with status code different than OK(200)", Message: ""}
var ErrorCode002 = SDKError{Code: 2, Meaning: "JSON deserialization failure", Message: ""}
var ErrorCode003 = SDKError{Code: 3, Meaning: "Transaction was not found. No confirmation can be given regarding transaction execution", Message: ""}
var ErrorCode004 = SDKError{Code: 4, Meaning: "Failed to decode value", Message: ""}
var ErrorCode005 = SDKError{Code: 5, Meaning: "Failed to recieve value from Node", Message: ""}

func newError(err error, wrapper SDKError) error {
	if err == nil {
		return nil
	}
	wrapper.Message = err.Error()
	return &wrapper
}
