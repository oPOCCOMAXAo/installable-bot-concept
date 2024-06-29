package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/hx"
)

type KeyRequest struct {
	Key string `form:"key" binding:"required"`
}

func (s *Service) Key(ctx *gin.Context) {
	var req KeyRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusOK, templates.DashboardPage(templates.DashboardData{
			Errors: []string{err.Error()},
		}))

		return
	}

	err = s.demo.SetKey(ctx.Request.Context(), req.Key)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusOK, templates.DashboardPage(templates.DashboardData{
			Errors: []string{err.Error()},
		}))

		return
	}

	hx.Redirect(ctx, "/dashboard")
}
