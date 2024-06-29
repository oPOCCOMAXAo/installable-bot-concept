package endpoints

import (
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/demo"
)

type Service struct {
	auth *auth.Service
	demo *demo.Service
}

func New(
	auth *auth.Service,
	demo *demo.Service,
) *Service {
	return &Service{
		auth: auth,
		demo: demo,
	}
}
