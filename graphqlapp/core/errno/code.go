package errno

import (
  "graphqlapp/core/constants"
  "net/http"
)

var (
  // UnknownErr 其他错误
  UnknownErr = &Errno{HTTPCode: http.StatusOK, Code: constants.UnknownErrorCode, Message: "unknown error"}
  // ReqErr 参数错误
  ReqErr = &Errno{HTTPCode: http.StatusOK, Code: constants.RequestErrorCode, Message: "request error"}
  // ResourceErr 资源错误
  ResourceErr = &Errno{HTTPCode: http.StatusOK, Code: constants.ResourceErrorCode, Message: "resource error"}
  // DatabaseErr 数据库错误
  DatabaseErr = &Errno{HTTPCode: http.StatusOK, Code: constants.DatabaseErrorCode, Message: "database error"}
  // TokenErr jwt token 错误
  TokenErr = &Errno{HTTPCode: http.StatusOK, Code: constants.TokenErrorCode, Message: "token error"}
  // NotFoundErr route not found
  NotFoundErr = &Errno{HTTPCode: http.StatusOK, Code: constants.NotFoundErrorCode, Message: "route not found"}
  // AuthErr 权限错误
  AuthErr = &Errno{HTTPCode: http.StatusOK, Code: constants.AuthErrorCode, Message: "auth error"}
)
