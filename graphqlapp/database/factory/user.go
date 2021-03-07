package factory

import (
  "fmt"
  "graphqlapp/app/models"
)

func usersTableSeeder() {
  for i := 0; i < 10; i++ {
    u := &models.User{
      Name: fmt.Sprintf("user-%d", i),
      Email: fmt.Sprintf("test%d@test.com", i),
    }
    models.CreateModel(&u)
  }
}
