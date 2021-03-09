package graphql_test

import (
  "github.com/stretchr/testify/require"
  "testing"
)

func TestUser(t *testing.T) {
  t.Run("get user", func(t *testing.T) {
    var resp struct {
      User struct {
        ID int
      }
    }

    graphqlClient.MustPost(`query { user(id: 22) { id } }`, &resp)
    require.Equal(t, 22, resp.User.ID)
  })
}
