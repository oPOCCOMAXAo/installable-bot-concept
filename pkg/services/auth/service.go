package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth/repo"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo          *repo.Repo
	isInitialized bool
}

func New(
	ctx context.Context,
	repo *repo.Repo,
) (*Service, error) {
	res := &Service{
		repo: repo,
	}

	var err error

	res.isInitialized, err = res.hasAdminPassword(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) hasAdminPassword(ctx context.Context) (bool, error) {
	pw, err := s.repo.GetAdminPassword(ctx)
	if err != nil {
		return false, err
	}

	return pw != "", nil
}

type MiddlewareParams struct {
	SetAdminPasswordPath string
}

func (s *Service) Middleware(
	params MiddlewareParams,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !s.isInitialized && ctx.Request.URL.Path != params.SetAdminPasswordPath {
			ctx.Redirect(http.StatusFound, params.SetAdminPasswordPath)
			ctx.Abort()

			return
		}
	}
}

func (s *Service) hashPassword(
	password string,
) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(res), nil
}

func (s *Service) SetAdminPassword(
	ctx context.Context,
	value string,
) error {
	value, err := s.hashPassword(value)
	if err != nil {
		return err
	}

	err = s.repo.SetAdminPassword(ctx, value)
	if err != nil {
		return err
	}

	s.isInitialized = true

	return nil
}
