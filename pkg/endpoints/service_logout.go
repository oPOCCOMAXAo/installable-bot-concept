package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/hx"
)

func (s *Service) Logout(ctx *gin.Context) {
	s.auth.ClearAuth(ctx)
	hx.Redirect(ctx, "/auth")
}
