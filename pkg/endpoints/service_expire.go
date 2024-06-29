package endpoints

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Expire(ctx *gin.Context) {
	s.auth.ClearAuth(ctx)
}
