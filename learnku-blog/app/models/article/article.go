package article

import (
	"goblog/app/models"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User

	CategoryID uint64 `gorm:"not null;index" valid:"category"`
	Category   category.Category
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
