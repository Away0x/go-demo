package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.AutoMigrate(&Link{})
		return err
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.DropTable(&Link{})
		return err
	}

	migrate.Add("2022_01_24_142034_add_links_table", up, down)
}
