package context

import (
  "github.com/spf13/viper"
  "graphqlapp/core/constants"
  "graphqlapp/core/errno"
  "net/http"
)

type RespData map[string]interface{}

type CommonResponse struct {
  Code    constants.LogicCode `json:"code"`
  Message string              `json:"msg"`
  Data    interface{}         `json:"data,omitempty"`
}

// NewCommonResponse new CommonResponse
func NewCommonResponse(code constants.LogicCode, message string, data interface{}) *CommonResponse {
  return &CommonResponse{
    Code:    code,
    Message: message,
    Data:    data,
  }
}

// NewSuccessResponse new success response
func NewSuccessResponse(message string, data interface{}) *CommonResponse {
  return NewCommonResponse(constants.SuccessCode, message, data)
}

// NewErrResponse new error response
func NewErrResponse(e *errno.Errno) *CommonResponse {
  return NewCommonResponse(e.Code, e.Message, e.Data)
}

func (c *AppContext) SuccessJSON(data interface{}) error {
  return c.JSON(http.StatusOK, NewSuccessResponse("ok", data))
}

func (c *AppContext) ErrorJSON(e *errno.Errno) error {
  return c.JSON(e.HTTPCode, NewErrResponse(e))
}

func (c *AppContext) SuccessHTML(tpl string, data interface{}) error {
  name := tpl + "." + viper.GetString("APP.TEMPLATE_EXT")

  if typed, ok := data.(RespData); ok {
    return c.Render(http.StatusOK, name, map[string]interface{}(typed))
  }

  return c.Render(http.StatusOK, name, data)
}

func (c *AppContext) SuccessHTMLNoData(tpl string) error {
  return c.SuccessHTML(tpl, map[string]interface{}{})
}
