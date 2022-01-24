package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type User struct {
		models.BaseModel

		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.AutoMigrate(&User{})
		return err
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.DropTable(&User{})
		return err
	}

	migrate.Add("2022_01_24_112128_user", up, down)
}
