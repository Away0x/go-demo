package controllers_test

import (
  "github.com/gavv/httpexpect/v2"
  "graphqlapp/core/constants"
  "testing"
)

func TestApiController(t *testing.T) {
  var (
    client = apiClient(t)
    url = "/test"
    resp *httpexpect.Object
  )

  resp = getOKApiJSon(client.GET(url))
  resp.Value("code").Equal(constants.SuccessCode)
}
