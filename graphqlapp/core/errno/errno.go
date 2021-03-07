package errno

import (
  "fmt"
  "graphqlapp/core/constants"
)

// Errno 项目统一错误类型
type Errno struct {
  HTTPCode int
  Message  string
  Code     constants.LogicCode
  Err      error
  Data     interface{}
}

func (e Errno) Error() string {
  return e.Message
}

func (e Errno) WithErr(err error) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  err.Error(),
    Err:      err,
  }
}

func (e Errno) WitData(d interface{}) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Data:     d,
  }
}

func (e Errno) WithMessage(msg string) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  msg,
    Err:      e.Err,
  }
}

func (e Errno) WithMessagef(format string, args ...interface{}) error {
  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  fmt.Sprintf(format, args...),
    Err:      e.Err,
  }
}

func (e Errno) WithErrMessage(err error, msg string) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  msg,
    Err:      err,
  }
}

func (e Errno) WithErrMessagef(err error, format string, args ...interface{}) error {
  if err == nil {
    return nil
  }

  return &Errno{
    HTTPCode: e.HTTPCode,
    Code:     e.Code,
    Message:  fmt.Sprintf(format, args...),
    Err:      err,
  }
}
