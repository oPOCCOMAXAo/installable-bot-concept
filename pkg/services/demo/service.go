package demo

import (
	"context"

	"github.com/opoccomaxao/installable-bot-concept/pkg/models"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/demo/repo"
)

type Service struct {
	repo *repo.Repo
}

func New(
	repo *repo.Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) IncrementCounter(
	ctx context.Context,
) error {
	return s.repo.IncrementParam(ctx, models.ParamCounter, 1)
}

func (s *Service) SetKey(
	ctx context.Context,
	key string,
) error {
	return s.repo.CreateUpdateParam(ctx, models.ParamKey, key)
}

func (s *Service) GetAll(
	ctx context.Context,
) (map[string]string, error) {
	params, err := s.repo.GetByIDs(ctx, []models.ParamName{
		models.ParamCounter,
		models.ParamKey,
	})
	if err != nil {
		return nil, err
	}

	res := make(map[string]string, len(params))
	for _, param := range params {
		res[string(param.ID)] = param.Value
	}

	return res, nil
}
