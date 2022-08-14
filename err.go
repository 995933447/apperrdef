package apperrdef

import (
	"fmt"
	"sync"
)

type ErrCode int32

const (
	ErrCodeNil ErrCode = 0
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

func (e *Error) IsErr(err error) bool {
	okErr, ok := ToError(err)
	if !ok {
		return false
	}
	return okErr.GetErrCode() == e.GetErrCode()
}

func NewErrWithMsg(code ErrCode, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func GetErrCode(err error) ErrCode {
	if okErr, ok := ToError(err); ok {
		return okErr.code
	}

	return ErrCodeNil
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

func NewErr(code ErrCode) *Error {
	return NewErrWithMsg(code, getDefaultErrMsgDefaultEmpty(code))
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