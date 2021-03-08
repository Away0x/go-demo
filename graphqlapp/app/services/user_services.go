package services

import (
  "graphqlapp/app/models"
)

type IUserServices interface {
  List(int, int) ([]*models.User, int64, error)
}

type UserServices struct {}

func NewUserServices() *UserServices {
  return &UserServices{}
}

func (u *UserServices) List(page, perPage int) (users []*models.User, total int64, err error) {
  offset := 0

  page = page - 1
  if page == 0 {
    offset = 0
  } else {
    offset = page * perPage
  }

  users = make([]*models.User, 0)
  err = models.DB().Offset(offset).Limit(perPage).Order("id desc").Find(&users).Error
  if err != nil {
    return
  }

  err = models.DB().Model(&models.User{}).Count(&total).Error
  return
}
