package services

import "graphqlapp/app/models"

type IUserServices interface {
  Detail(id int) (*models.User, error)
  List() ([]*models.User, error)
  Create(u *models.User) (*models.User, error)
}

type UserServices struct {}

func NewUserServices() *UserServices {
  return &UserServices{}
}

func (*UserServices) List() (users []*models.User, err error) {
  users = make([]*models.User, 0)
  err = models.DB().Find(&users).Error
  return
}

func (*UserServices) Detail(id int) (user *models.User, err error) {
  user = new(models.User)
  err = models.DB().First(user, id).Error
  return
}

func (*UserServices) Create(u *models.User) (*models.User, error) {
  err := models.CreateModel(u)
  return u, err
}
