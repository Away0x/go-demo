package context

import (
  "graphqlapp/core/errno"
  "graphqlapp/core/pkg/validator"
  "strconv"
)

func (c *AppContext) IntParam(key ...string) (int, error) {
  k := key[0]
  if k == "" {
    k = "id"
  }

  i, err := strconv.Atoi(c.Param(k))
  if err != nil {
    return 0, errno.ReqErr.WithErr(err)
  }

  return i, nil
}

func (c *AppContext) IntQuery(key ...string) (int, error) {
  k := key[0]
  if k == "" {
    k = "id"
  }

  i, err := strconv.Atoi(c.QueryParam(k))
  if err != nil {
    return 0, errno.ReqErr.WithErr(err)
  }

  return i, nil
}

func (c *AppContext) BindValidatorStruct(v validator.Validator) error {
  if err := c.Bind(v); err != nil {
    return err
  }

  if errs, ok := validator.ValidateStruct(v); !ok {
    return errno.ReqErr.WitData(errs)
  }

  return nil
}
