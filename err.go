package apperrdef

import (
	"fmt"
	"math"
	"sync"
)

type ErrCode int32

const (
	ErrCodeUnknown ErrCode = math.MinInt32
)

type Error struct {
	code ErrCode
	msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:%d msg:%s", e.code, e.msg)
}

func (e *Error) GetErrMsg() string {
	return e.msg
}

func (e *Error) GetErrCode() ErrCode {
	return e.code
}

func NewErrorWithMsg(code ErrCode, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func GetErrCode(err error) ErrCode {
	if okErr, ok := ToError(err); ok {
		return okErr.code
	}

	return ErrCodeUnknown
}

func GetErrMsg(err error) string {
	if okErr, ok := ToError(err); ok {
		return okErr.msg
	}

	return err.Error()
}

func ToError(err error) (*Error, bool) {
	okErr, ok := err.(*Error)
	return okErr, ok
}

func NewError(code ErrCode) *Error {
	return NewErrorWithMsg(code, getDefaultErrMsgDefaultEmpty(code))
}

var defaultErrMsgMap sync.Map

func getDefaultErrMsgDefaultEmpty(code ErrCode) string {
	msg, ok := defaultErrMsgMap.Load(code)
	if !ok {
		return ""
	}

	return msg.(string)
}

func RegisterDefaultErrMsg(code ErrCode, msg string) {
	defaultErrMsgMap.Store(code, msg)
}