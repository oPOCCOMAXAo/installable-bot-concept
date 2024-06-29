package repo

import (
	"context"
	"strconv"

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

func (r *Repo) IncrementParam(
	ctx context.Context,
	id models.ParamName,
	delta int64,
) error {
	err := r.db.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			var param models.Param

			err := tx.
				Where("id = ?", id).
				Take(&param).
				Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.WithStack(err)
			}

			if param.ID == "" {
				param = models.Param{
					ID:    id,
					Value: "0",
				}
			}

			value, err := strconv.ParseInt(param.Value, 10, 64)
			if err != nil {
				return errors.WithStack(err)
			}

			value += delta

			param.Value = strconv.FormatInt(value, 10)

			err = tx.Save(&param).Error
			if err != nil {
				return errors.WithStack(err)
			}

			return nil
		})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateUpdateParam(
	ctx context.Context,
	id models.ParamName,
	value string,
) error {
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{"value"}),
		}).
		Create(&models.Param{
			ID:    id,
			Value: value,
		}).
		Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Repo) GetByIDs(
	ctx context.Context,
	ids []models.ParamName,
) ([]*models.Param, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var res []*models.Param

	err := r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Order("id ASC").
		Find(&res).
		Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
