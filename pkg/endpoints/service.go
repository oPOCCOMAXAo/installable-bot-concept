package endpoints

import "github.com/opoccomaxao/installable-bot-concept/pkg/services/auth"

type Service struct {
	auth *auth.Service
}

func New(
	auth *auth.Service,
) *Service {
	return &Service{
		auth: auth,
	}
}
