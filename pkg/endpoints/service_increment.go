package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/hx"
)

func (s *Service) Increment(ctx *gin.Context) {
	err := s.demo.IncrementCounter(ctx.Request.Context())
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusOK, templates.DashboardPage(templates.DashboardData{
			Errors: []string{err.Error()},
		}))

		return
	}

	hx.Redirect(ctx, "/dashboard")
}
