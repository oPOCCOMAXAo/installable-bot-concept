package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
)

func (s *Service) Dashboard(ctx *gin.Context) {
	params, err := s.demo.GetAll(ctx.Request.Context())
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusOK, templates.DashboardPage(templates.DashboardData{
			Errors: []string{err.Error()},
		}))

		return
	}

	ginutils.RenderTempl(ctx, http.StatusOK, templates.DashboardPage(templates.DashboardData{
		Params: params,
	}))
}
