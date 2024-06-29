package auth

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/opoccomaxao/installable-bot-concept/pkg/models"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth/repo"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/hx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const SessionCookie = "session"
const SessionMaxAge = 3600
const Key = "12345678"
const Audience = "admin"
const Issuer = "bot"

type Service struct {
	repo          *repo.Repo
	isInitialized bool
	jwt           *jwt.Parser
}

func New(
	ctx context.Context,
	repo *repo.Repo,
) (*Service, error) {
	res := &Service{
		repo: repo,
		jwt:  jwt.NewParser(),
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

func (s *Service) jwtKey(*jwt.Token) (any, error) {
	return []byte(Key), nil
}

type MiddlewareParams struct {
	SetAdminPasswordPath string
	AuthPath             string
}

func (s *Service) Middleware(
	params MiddlewareParams,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !s.isInitialized && ctx.Request.URL.Path != params.SetAdminPasswordPath {
			hx.Redirect(ctx, params.SetAdminPasswordPath)

			return
		}

		token := ginutils.GetCookieOrEmpty(ctx, SessionCookie)
		if token == "" && ctx.Request.URL.Path != params.AuthPath {
			hx.Redirect(ctx, params.AuthPath)

			return
		}

		err := s.ValidateAuth(ctx.Request.Context(), token)
		if err != nil {
			s.ClearAuth(ctx)
			hx.Redirect(ctx, params.AuthPath)

			return
		}

		s.SetAuth(ctx)
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

func (s *Service) verifyPassword(
	hashedPassword string,
	password string,
) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
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

func (s *Service) SetAuth(ctx *gin.Context) error {
	token, err := s.Sign(ctx.Request.Context())
	if err != nil {
		return err
	}

	ctx.SetCookie(SessionCookie, token, SessionMaxAge, "", "", false, true)

	return nil
}

func (s *Service) ClearAuth(ctx *gin.Context) {
	ctx.SetCookie(SessionCookie, "", -1, "", "", false, true)
}

func (s *Service) ValidateAuth(
	ctx context.Context,
	token string,
) error {
	var claims jwt.RegisteredClaims

	_, err := s.jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		for _, aud := range claims.Audience {
			if aud != Audience {
				return nil, models.ErrInvalidAuth
			}
		}

		return []byte(Key), nil
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Service) Sign(
	ctx context.Context,
) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    Issuer,
		Audience:  []string{Audience},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(SessionMaxAge * time.Second)),
	}

	res, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, &claims).
		SignedString([]byte(Key))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return res, nil
}

func (s *Service) ValidateAdminPassword(
	ctx context.Context,
	value string,
) error {
	hashedPassword, err := s.repo.GetAdminPassword(ctx)
	if err != nil {
		return err
	}

	err = s.verifyPassword(hashedPassword, value)
	if err != nil {
		return err
	}

	return nil
}
