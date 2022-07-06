package framework

import (
	"fmt"

	"github.com/oracle/oci-go-sdk/v49/common"
)

type serviceerror struct {
	StatusCode int
	Code       string
	Message    string
}

func NewServiceError(statusCode int, code string, message string) common.ServiceError {
	return serviceerror{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

func (se serviceerror) Error() string {
	return fmt.Sprintf("Service error:%s. %s. http status code: %d",
		se.Code, se.Message, se.StatusCode)
}

func (se serviceerror) GetHTTPStatusCode() int {
	return se.StatusCode
}

func (se serviceerror) GetMessage() string {
	return se.Message
}

func (se serviceerror) GetCode() string {
	return se.Code
}

func (se serviceerror) GetOpcRequestID() string {
	return "Not implemented"
}
