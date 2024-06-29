package repo

import (
	"context"

	"github.com/opoccomaxao/installable-bot-concept/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetAdminPassword(
	ctx context.Context,
) (string, error) {
	var res models.Param

	err := r.db.WithContext(ctx).
		Model(&res).
		Where("id = ?", models.ParamAdminPassword).
		Take(&res).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}

		return "", errors.WithStack(err)
	}

	return res.Value, nil
}

func (r *Repo) SetAdminPassword(
	ctx context.Context,
	value string,
) error {
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).
		Create(&models.Param{
			ID:    models.ParamAdminPassword,
			Value: value,
		}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
