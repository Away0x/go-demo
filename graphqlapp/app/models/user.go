package models

type User struct {
  BaseModel
  Name  string `gorm:"type:varchar(255);not null" json:"name"`
  Email string `gorm:"type:varchar(255);unique;not null" json:"email"`
}
