package apperrdef

import (
	"fmt"
	"sync"
)

type ErrCode int32

const (
	ErrCodeUnknown ErrCode = -1
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
	if appErr, ok := err.(*Error); ok {
		return appErr.code
	}

	return ErrCodeUnknown
}

func GetErrMsg(err error) string {
	if appErr, ok := err.(*Error); ok {
		return appErr.msg
	}

	return err.Error()
}

func IsErr(err error) bool {
	_, ok := err.(*Error)
	return ok
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