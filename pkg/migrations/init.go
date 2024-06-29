package migrations

import (
	"github.com/opoccomaxao/installable-bot-concept/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	migr := db.Migrator()

	err := migr.AutoMigrate(
		&models.Param{},
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
