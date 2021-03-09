package services_test

import (
  "github.com/stretchr/testify/require"
  "graphqlapp/app/services"
  "testing"
)

func TestUserServices(t *testing.T) {
  us := services.NewUserServices()
  users, total, err := us.List(1, 3)
  require.Equal(t, err, nil)
  require.Equal(t, int(total), 13)
  require.NotEqual(t, users, nil)
}
