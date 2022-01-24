package migrations

import (
	"database/sql"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	type User struct {
		City          string `gorm:"type:varchar(10);"`
		Indtroduction string `gorm:"type:varchar(255);"`
		Avatar        string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) error {
		err := migrator.AutoMigrate(&User{})
		return err
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) (err error) {
		err = migrator.DropColumn(&User{}, "City")
		err = migrator.DropColumn(&User{}, "Indtroduction")
		err = migrator.DropColumn(&User{}, "Avatar")
		return
	}

	migrate.Add("2022_01_24_164821_add_fields_to_user", up, down)
}
