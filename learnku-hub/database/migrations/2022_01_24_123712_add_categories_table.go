package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type Category struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.AutoMigrate(&Category{})
		return err
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.DropTable(&Category{})
		return err
	}

	migrate.Add("2022_01_24_123712_add_categories_table", up, down)
}
